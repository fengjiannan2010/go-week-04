package logic

import (
	"context"
	"oburn/rpc/client/odrive"

	"oburn/rpc/internal/svc"
	"oburn/rpc/oburn"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetDriveInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetDriveInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDriveInfoLogic {
	return &GetDriveInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetDriveInfoLogic) GetDriveInfo(in *oburn.EmptyRequest) (*oburn.DriveInfoResponse, error) {
	// todo: add your logic here and delete this line
	response, err := l.svcCtx.Odrive.GetDriveInfo(l.ctx, &odrive.EmptyRequest{})
	responseMessage := response.GetMessage()
	driveInfo:=response.GetDriveInfo()
	return &oburn.DriveInfoResponse{
		Message: &oburn.Message{
			IsSuccess: responseMessage.GetIsSuccess(),
			Code:      responseMessage.GetCode(),
			Message:   responseMessage.GetMessage(),
		},
		DriveInfo:  &oburn.DriveInfo{
			ProductId:            driveInfo.GetProductId(),
			Vendor:               driveInfo.GetVendor(),
			SerialNo:             driveInfo.GetSerialNo(),
			FwVersion:            driveInfo.GetFwVersion(),
			CdReadTime:           driveInfo.GetCdReadTime(),
			CdWriteTime:          driveInfo.GetCdWriteTime(),
			DvdReadTime:          driveInfo.GetDvdReadTime(),
			DvdWriteTime:         driveInfo.GetDvdWriteTime(),
			BdReadTime:           driveInfo.GetBdReadTime(),
			BdWriteTime:          driveInfo.GetBdWriteTime(),
			PowerOnTime:          driveInfo.GetPowerOnTime(),
		},
	}, err
}
