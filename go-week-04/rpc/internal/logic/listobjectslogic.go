package logic

import (
	"context"
	"oburn/rpc/client/odrive"
	"oburn/utils"
	"path/filepath"

	"oburn/rpc/internal/svc"
	"oburn/rpc/oburn"

	"github.com/tal-tech/go-zero/core/logx"
)

type ListObjectsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListObjectsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListObjectsLogic {
	return &ListObjectsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListObjectsLogic) ListObjects(in *oburn.ListObjectsRequest) (*oburn.ListObjectsResponse, error) {
	// todo: add your logic here and delete this line
	response, err := l.svcCtx.Odrive.ListObjects(l.ctx, &odrive.ListObjectsRequest{
		FilePathHash: utils.MD5V([]byte(filepath.ToSlash(in.GetFilePath()))),
		FilePath:     in.GetFilePath(),
	})
	responseMessage := response.GetMessage()
	var objectInfos []*oburn.ObjectInfo
	for _, objectInfo := range response.GetObjectInfo() {
		object := &oburn.ObjectInfo{
			Name:    objectInfo.GetName(),
			Size:    objectInfo.GetSize(),
			Mode:    objectInfo.GetMode(),
			ModTime: objectInfo.GetModTime(),
			IsDir:   objectInfo.GetIsDir(),
		}
		objectInfos = append(objectInfos, object)
	}
	return &oburn.ListObjectsResponse{
		Message: &oburn.Message{
			IsSuccess: responseMessage.GetIsSuccess(),
			Code:      responseMessage.GetCode(),
			Message:   responseMessage.GetMessage(),
		},
		ObjectInfo: objectInfos,
	}, err
}
