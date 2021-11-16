package config

type MinioOSS struct {
	Endpoint        string `json:",optional"`          //TODO S3 端节点
	AccessKeyId     string `json:",optional"`     //TODO S3 AccessKey
	AccessKeySecret string `json:",optional"` //TODO S3 SecretAccessKey
	UseSSL          bool   `json:",optional"`           //TODO S3 SSL
	InstanceID      string `json:",optional"`       //TODO S3 实例名称
	IsOChunkClient  string `json:",optional"`
}
