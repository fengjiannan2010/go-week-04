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

type MakeBucketLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMakeBucketLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MakeBucketLogic {
	return &MakeBucketLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MakeBucketLogic) MakeBucket(in *oburn.MakeBucketRequest) (*oburn.MakeBucketResponse, error) {
	// todo: add your logic here and delete this line
	response, err := l.svcCtx.Odrive.MakeBucket(l.ctx, &odrive.MakeBucketRequest{
		FilePathHash: utils.MD5V([]byte(filepath.ToSlash(in.GetDiscPath()))),
	})
	responseMessage := response.GetMessage()
	return &oburn.MakeBucketResponse{
		Message: &oburn.Message{
			IsSuccess: responseMessage.GetIsSuccess(),
			Code:      responseMessage.GetCode(),
			Message:   responseMessage.GetMessage(),
		},
	}, err
}
