package logic

import (
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/pkg/errors"
	"github.com/tal-tech/go-zero/core/logx"
	"oburn/rpc/internal/svc"
	"os"
)

type MinioLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMinioLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MinioLogic {
	return &MinioLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (m *MinioLogic) GetObjectStream(bucketName, objectName string) (*minio.Object, error) {
	opts := minio.GetObjectOptions{}
	opts.Set(m.svcCtx.Config.MinioConfig.IsOChunkClient, "")
	object, err := m.svcCtx.S3Client.GetObject(m.ctx, bucketName, objectName, opts)
	if err != nil {
		return nil, errors.Wrap(err, "get object reader stream error")
	}
	return object, nil
}

func (m *MinioLogic) GetObjectInfo(object *minio.Object) (minio.ObjectInfo, error) {
	objectInfo, err := object.Stat()
	if err != nil {
		return minio.ObjectInfo{}, errors.Wrap(err, "get object info error")
	}
	return objectInfo, nil
}

func (m *MinioLogic) RemoveObject(bucketName, objectName string, removeObjectOptions minio.RemoveObjectOptions) error {
	err := m.svcCtx.S3Client.RemoveObject(m.ctx, bucketName, objectName, removeObjectOptions)
	if err != nil {
		return errors.Wrap(err, "delete objects error")
	}
	return nil
}

func (m *MinioLogic) RemoveObjects(bucketName, objectName string) error {
	objectsCh := make(chan minio.ObjectInfo)
	opts := minio.RemoveObjectsOptions{
		GovernanceBypass: true,
	}
	for err := range m.svcCtx.S3Client.RemoveObjects(m.ctx, bucketName, objectsCh, opts) {
		if err.Err != nil {
			return errors.Wrap(err.Err, "Error detected during deletion")
		}
	}
	return nil
}

func (m *MinioLogic) PutObject(fileName, bucketName, objectName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return errors.Wrap(err, "Open file error")
	}
	defer file.Close()
	fileStat, err := file.Stat()
	if err != nil {
		return errors.Wrap(err, "minio object stat error")
	}

	opts := minio.PutObjectOptions{}
	opts.UserMetadata = map[string]string{m.svcCtx.Config.MinioConfig.IsOChunkClient: ""}
	opts.ContentType = "application/octet-stream"
	info, err := m.svcCtx.S3Client.PutObject(m.ctx, bucketName, objectName, file, fileStat.Size(), opts)
	if err != nil {
		return errors.Wrap(err, "putobject error")
	}
	logx.Info("seccessfully uploaded:", info)
	return nil
}
