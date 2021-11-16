package logic

import (
	"context"
	"oburn/rpc/client/odrive"
	"oburn/rpc/internal/svc"
	"oburn/rpc/oburn"

	"github.com/tal-tech/go-zero/core/logx"
)

type SystemInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSystemInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SystemInfoLogic {
	return &SystemInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SystemInfoLogic) SystemInfo(in *oburn.EmptyRequest) (*oburn.SystemInfoResponse, error) {
	// todo: add your logic here and delete this line
	response, err := l.svcCtx.Odrive.SystemInfo(l.ctx, &odrive.EmptyRequest{})
	responseMessage := response.GetMessage()
	responseVersion := response.GetFirmwareVersion()
	var firmwareVersion []*oburn.FirmwareVersion
	for _, item := range responseVersion {
		firmwareVersion = append(firmwareVersion, &oburn.FirmwareVersion{
			AppVersion: item.GetAppVersion(),
			AppDate:    item.GetAppDate(),
			Tag:        item.GetTag(),
		})
	}
	return &oburn.SystemInfoResponse{
		Message: &oburn.Message{
			IsSuccess: responseMessage.GetIsSuccess(),
			Code:      responseMessage.GetCode(),
			Message:   responseMessage.GetMessage(),
		},
		SysInfo:         response.GetSysInfo(),
		FirmwareVersion: firmwareVersion,
	}, err
}
