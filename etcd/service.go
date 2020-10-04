// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package etcd

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"github.com/pkg/errors"
	"github.com/ysicing/ext/logger"
	"sync"
)

type Instance struct {
	ID       string
	Name     string
	Type     string
	Endpoint string
	Desc     string
}

type Service struct {
	instance *Instance
	leaseID  clientv3.LeaseID
	close    chan struct{}
	wg       sync.WaitGroup
	etcdapi  *EtcdClient
}

func NewService(instance *Instance, etcdapi *EtcdClient) (*Service, error) {
	return &Service{
		instance: instance,
		close:    make(chan struct{}),
		etcdapi:  etcdapi,
	}, nil
}

func (s *Service) Register() error {
	res, err := s.etcdapi.Clientv3.Grant(context.TODO(), 10)
	if err != nil {
		logger.Slog.Error(err)
	}
	s.leaseID = res.ID
	val, err := json.Marshal(&s.instance)
	if err != nil {
		logger.Slog.Error(err)
	}
	key := fmt.Sprintf("%s/%s", s.instance.Type, s.instance.Name)
	if _, err = s.etcdapi.Clientv3.Put(context.TODO(), key, string(val), clientv3.WithLease(s.leaseID)); err != nil {
		return err
	}
	logger.Slog.Infof("注册成功, ID为: %v, 服务类型: %v, 注册地址: %v", s.instance.ID, key, s.instance.Endpoint)
	ch, err := s.etcdapi.Clientv3.KeepAlive(context.TODO(), s.leaseID)
	if err != nil {
		return err
	}
	s.wg.Add(1)
	defer s.wg.Done()
	for {
		select {
		case <-s.close:
			// revoke
			return s.revoke()
		case <-s.etcdapi.Clientv3.Ctx().Done():
			return errors.New("etcd close")
		case _, ok := <-ch:
			if !ok {
				// revoke
				return s.revoke()
			}
		}
	}
}

func (s *Service) revoke() error {
	_, err := s.etcdapi.Clientv3.Revoke(context.TODO(), s.leaseID)
	if err != nil {
		logger.Slog.Error("revoke err: ", err.Error())
	} else {
		key := fmt.Sprintf("%s/%s", s.instance.Type, s.instance.Name)
		logger.Slog.Debugf("revoke done: %v, endpoint: %v", key, s.instance.Endpoint)
	}
	return err
}

func (s *Service) Close() {
	close(s.close)
	s.wg.Wait()
	if err := s.etcdapi.Clientv3.Close(); err != nil {
		logger.Slog.Error(err)
	}
}

func (s *Service) GetService(prefix string) (interface{}, error) {
	resp, err := s.etcdapi.Clientv3.KV.Get(context.Background(), prefix, clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}
	for _, kv := range resp.Kvs {
		logger.Slog.Debug(kv.Key, kv.Value, kv.Lease)
	}

	return resp, nil
}
