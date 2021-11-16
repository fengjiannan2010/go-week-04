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

type RemoveBucketLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRemoveBucketLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveBucketLogic {
	return &RemoveBucketLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RemoveBucketLogic) RemoveBucket(in *oburn.RemoveBucketRequest) (*oburn.RemoveBucketResponse, error) {
	// todo: add your logic here and delete this line
	response, err := l.svcCtx.Odrive.RemoveBucket(l.ctx, &odrive.RemoveBucketRequest{
		FilePathHash: utils.MD5V([]byte(filepath.ToSlash(in.GetDiscPath()))),
	})
	responseMessage := response.GetMessage()
	return &oburn.RemoveBucketResponse{
		Message: &oburn.Message{
			IsSuccess: responseMessage.GetIsSuccess(),
			Code:      responseMessage.GetCode(),
			Message:   responseMessage.GetMessage(),
		},
	}, err
}
