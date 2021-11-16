package config

import (
	"github.com/tal-tech/go-zero/zrpc"
	"zerogorm"
)

type Config struct {
	zrpc.RpcServerConf
	zrpc.RpcClientConf
	MinioConfig MinioOSS
	OrmConf     zerogorm.Config
	VerifyMode  int32 `json:",default=1"`
}

func (c Config) HasMinio() bool {
	return len(c.MinioConfig.Endpoint) > 0 && len(c.MinioConfig.AccessKeyId) > 0 && len(c.MinioConfig.AccessKeySecret) > 0
}
