// MIT License
// Copyright (c) 2020 ysicing <i@ysicing.me>

package ginmid

import (
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

const (
	namespace = "gin"
)

var (
	labels = []string{"status", "endpoint", "method", "host"}

	uptime = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: namespace,
			Name:      "uptime",
			Help:      "HTTP service uptime.",
		}, nil,
	)
)

// init registers the prometheus metrics
func init() {
	prometheus.MustRegister(uptime)
	go recordUptime()
}

// recordUptime increases service uptime per second.
func recordUptime() {
	for range time.Tick(time.Second) {
		uptime.WithLabelValues().Inc()
	}
}
