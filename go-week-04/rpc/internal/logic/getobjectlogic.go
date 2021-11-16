package logic

import (
	"context"
	"oburn/data"
	"oburn/rpc/internal/svc"
	"oburn/rpc/oburn"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetObjectLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetObjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetObjectLogic {
	return &GetObjectLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetObjectLogic) GetObject(in *oburn.GetObjectRequest) (*oburn.GetObjectResponse, error) {
	// todo: add your logic here and delete this line
	taskInfo:= data.TaskInfo{

	}
	l.svcCtx.TaskInfoModel.Exists(taskInfo)
	l.svcCtx.TaskInfoModel.Insert(taskInfo)
	//response, err := l.svcCtx.Odrive.GetObject(l.ctx, &odrive.GetObjectRequest{
	//	FilePathHash:         utils.MD5V([]byte(filepath.ToSlash(in.GetDiscPath()))),
	//	Offset:               in.GetOffset(),
	//	Count:                in.GetCount(),
	//})
	//responseMessage := response.
	return &oburn.GetObjectResponse{}, nil
}
