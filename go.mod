module github.com/ysicing/ext

go 1.15

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/go-playground/validator/v10 v10.3.0 // indirect
	github.com/gomodule/redigo v1.8.2
	github.com/google/uuid v1.1.2
	github.com/mitchellh/go-homedir v1.1.0
	github.com/prometheus/client_golang v1.7.1
	github.com/prometheus/common v0.13.0 // indirect
	github.com/stretchr/testify v1.5.1
	github.com/ugorji/go v1.1.8 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.16.0
	golang.org/x/sys v0.0.0-20200918174421-af09f7315aff // indirect
	google.golang.org/protobuf v1.25.0 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
)

replace go.uber.org/zap v1.16.0 => github.com/BeidouCloudPlatform/zap v1.16.1-0.20200920005712-401a59f8f6bb
