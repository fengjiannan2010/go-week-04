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

type StatObjectLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewStatObjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StatObjectLogic {
	return &StatObjectLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *StatObjectLogic) StatObject(in *oburn.StatObjectRequest) (*oburn.StatObjectResponse, error) {
	// todo: add your logic here and delete this line
	response, err := l.svcCtx.Odrive.StatObject(l.ctx, &odrive.StatObjectRequest{
		FilePathHash: utils.MD5V([]byte(filepath.ToSlash(in.GetFilePath()))),
		FilePath:     in.GetFilePath(),
	})
	responseMessage := response.GetMessage()
	obj := response.GetObjectInfo()
	return &oburn.StatObjectResponse{
		Message: &oburn.Message{
			IsSuccess: responseMessage.GetIsSuccess(),
			Code:      responseMessage.GetCode(),
			Message:   responseMessage.GetMessage(),
		},
		ObjectInfo: &oburn.ObjectInfo{
			Name:    obj.GetName(),
			Size:    obj.GetSize(),
			Mode:    obj.GetMode(),
			ModTime: obj.GetModTime(),
			IsDir:   obj.GetIsDir(),
		},
	}, err
}
