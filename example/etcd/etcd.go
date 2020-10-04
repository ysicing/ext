// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package main

import (
	"context"
	"fmt"
	"github.com/ysicing/ext/etcd"
	"github.com/ysicing/ext/logger"
	"github.com/ysicing/ext/utils/exid"
	"github.com/ysicing/ext/utils/exrand"
	"github.com/ysicing/ext/utils/extime"
)

func init() {
	logcfg := &logger.LogConfig{Simple: true}
	logger.InitLogger(logcfg)
}

func main() {
	cfg := etcd.NewCfg(etcd.EtcdConfig{
		Endpoint: []string{"10.147.20.44:2379"},
	})
	etcdapi, err := etcd.New(cfg)
	if err != nil {
		logger.Slog.Exit(-1, err)
	}
	respdata, err := etcdapi.Clientv3.Cluster.MemberList(context.Background())
	if err != nil {
		logger.Slog.Exit(-1, err)
	}
	for _, mem := range respdata.Members {
		logger.Slog.Debug(mem.Name, mem.ID, mem.PeerURLs)
	}
	id := exid.GenUUID()
	node := &etcd.Instance{
		ID:       id,
		Name:     "node",
		Type:     "/macnode",
		Endpoint: fmt.Sprintf("%v:%v", "127.0.0.1", exrand.NumRand(30000)),
		Desc:     extime.NowUnixString(),
	}
	svc, err := etcd.NewService(node, etcdapi)
	if err != nil {
		logger.Slog.Error(err)
	}
	go svc.Register()
	svcresp, err := svc.GetService("/macnode/node")
	if err != nil {
		logger.Slog.Exit(0, err)
	}
	logger.Slog.Info(svcresp)
}
