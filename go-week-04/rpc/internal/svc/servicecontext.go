package svc

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/tal-tech/go-zero/zrpc"
	"gorm.io/gorm"
	"oburn/data"
	"oburn/rpc/client/odriveclient"
	"oburn/rpc/internal/config"
	"zerogorm"
)

type ServiceContext struct {
	Config   config.Config
	Odrive   odriveclient.Odrive
	S3Client *minio.Client
	data.TaskInfoModel
	data.BurnConfigModel
	data.DiscInfoModel
	data.FilePartInfoModel
	data.BurnConfig
	OrmDb            *gorm.DB
	DiscInfo         data.DiscInfo
}

func NewServiceContext(c config.Config) *ServiceContext {
	odrive := odriveclient.NewOdrive(zrpc.MustNewClient(c.RpcClientConf))
	orm, err := zerogorm.New(c.OrmConf)
	if err != nil {
		panic("init gorm error")
	}
	svcCtx := ServiceContext{
		Config:            c,
		Odrive:            odrive,
		TaskInfoModel:     data.NewTaskInfoModel(orm),
		BurnConfigModel:   data.NewBurnConfigModel(orm),
		DiscInfoModel:     data.NewDiscInfoModel(orm),
		FilePartInfoModel: data.NewFilePartInfoModel(orm),
		BurnConfig:        data.BurnConfig{},
		OrmDb:             orm,
	}
	if c.HasMinio() {
		endpoint := c.MinioConfig.Endpoint
		accessKeyID := c.MinioConfig.AccessKeyId
		secretAccessKey := c.MinioConfig.AccessKeySecret
		useSSL := c.MinioConfig.UseSSL
		client, err := minio.New(endpoint, &minio.Options{
			Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
			Secure: useSSL,
		})
		if err != nil {
			client = nil
		}
		svcCtx.S3Client = client
	}
	return &svcCtx
}
