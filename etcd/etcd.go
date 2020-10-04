// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package etcd

import (
	"github.com/coreos/etcd/clientv3"
	"github.com/ysicing/ext/logger"
	"time"
)

type EtcdClient struct {
	Clientv3 *clientv3.Client
}

type EtcdConfig struct {
	Endpoint []string
	Timeout  time.Duration
}

func NewCfg(config EtcdConfig) *EtcdConfig {
	if len(config.Endpoint) == 0 {
		config.Endpoint = []string{"127.0.0.1:2379"}
	}
	if config.Timeout <= 5*time.Second {
		config.Timeout = 5 * time.Second
	}
	return &EtcdConfig{
		Endpoint: config.Endpoint,
		Timeout:  config.Timeout,
	}
}

func New(ecfg *EtcdConfig) (*EtcdClient, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   ecfg.Endpoint,
		DialTimeout: ecfg.Timeout,
	})
	if err != nil {
		logger.Slog.Errorf("init etcd client (%v) err: %v", ecfg.Endpoint, err)
		return nil, err
	}
	client := EtcdClient{Clientv3: cli}
	return &client, nil
}

func (ec EtcdClient) Close() {
	defer ec.Clientv3.Close()
}

//func (ec EtcdClient) Register(service string, value string, ttl ...int64) {
//	kv := clientv3.NewKV(ec.Clientv3)
//	lease := clientv3.NewLease(ec.Clientv3)
//	var curLeaseID clientv3.LeaseID = 0
//	var defaultttl int64 = 10
//	if len(ttl) != 0 {
//		defaultttl = ttl[0]
//	}
//	for {
//		if curLeaseID == 0 {
//			leaseResp, err := lease.Grant(context.TODO(), defaultttl)
//			if err != nil {
//				logger.Slog.Panic(err)
//			}
//			if _, err := kv.Put(context.TODO(), service, value, clientv3.WithLease(leaseResp.ID)); err != nil {
//				logger.Slog.Panic(err)
//			}
//			curLeaseID = leaseResp.ID
//		} else {
//			if _, err := lease.KeepAliveOnce(context.TODO(), curLeaseID); err == rpctypes.ErrLeaseNotFound {
//				curLeaseID = 0
//				continue
//			}
//		}
//		time.Sleep(time.Duration(1) * time.Second)
//	}
//}
