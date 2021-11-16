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

type RemoveObjectLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRemoveObjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveObjectLogic {
	return &RemoveObjectLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RemoveObjectLogic) RemoveObject(in *oburn.RemoveObjectRequest) (*oburn.RemoveObjectResponse, error) {
	// todo: add your logic here and delete this line
	response, err := l.svcCtx.Odrive.RemoveObject(l.ctx, &odrive.RemoveObjectRequest{
		RemoveMode: 0,
		FilePathHash: utils.MD5V([]byte(filepath.ToSlash(in.GetDiscPath()))),
	})
	responseMessage := response.GetMessage()
	return &oburn.RemoveObjectResponse{
		Message: &oburn.Message{
			IsSuccess: responseMessage.GetIsSuccess(),
			Code:      responseMessage.GetCode(),
			Message:   responseMessage.GetMessage(),
		},
	}, err
}
