package logic

import (
	"context"
	"oburn/data"
	"oburn/rpc/client/odrive"
	"oburn/rpc/internal/svc"
	"oburn/rpc/oburn"

	"github.com/tal-tech/go-zero/core/logx"
)

type SetBurnConfigLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetBurnConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetBurnConfigLogic {
	return &SetBurnConfigLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SetBurnConfigLogic) SetBurnConfig(in *oburn.SetBurnConfigRequest) (*oburn.SetBurnConfigResponse, error) {
	// todo: add your logic here and delete this line
	//文件系统类型 FsType 0:NotSet,1:ImageBurn,2:StreamBurn,3:UdfBurn,6:ExtImage,7:UnKnown
	response, err := l.svcCtx.Odrive.SetBurnConfig(l.ctx, &odrive.SetBurnConfigRequest{
		DiscLabel:  in.GetDiscLabel(),
		DiscPasswd: in.GetDiscPasswd(),
		FsType:     in.GetFsType(),
	})
	responseMessage := response.GetMessage()
	if responseMessage.IsSuccess {
		l.svcCtx.BurnConfig = data.BurnConfig{
			DiscLabel:  in.GetDiscLabel(),
			DiscPasswd: in.GetDiscPasswd(),
			FsType:     data.DiscFsType(in.GetFsType()),
			IsVerify:   in.GetIsVerify(),
		}
	}
	return &oburn.SetBurnConfigResponse{
		Message: &oburn.Message{
			IsSuccess: responseMessage.GetIsSuccess(),
			Code:      responseMessage.GetCode(),
			Message:   responseMessage.GetMessage(),
		},
	}, err
}
