package logic

import (
	"context"
	"oburn/rpc/client/odrive"

	"oburn/rpc/internal/svc"
	"oburn/rpc/oburn"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetDiscInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetDiscInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDiscInfoLogic {
	return &GetDiscInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetDiscInfoLogic) GetDiscInfo(in *oburn.EmptyRequest) (*oburn.GetDiscInfoResponse, error) {
	// todo: add your logic here and delete this line
	discInfoResponse, err := l.svcCtx.Odrive.GetDiscInfo(l.ctx, &odrive.EmptyRequest{})
	responseMessage := discInfoResponse.GetMessage()
	mediaInfo := discInfoResponse.GetMediaInfo()
	discFsInfo := discInfoResponse.GetDiscFsInfo()
	return &oburn.GetDiscInfoResponse{
		Message: &oburn.Message{
			IsSuccess: responseMessage.GetIsSuccess(),
			Code:      responseMessage.GetCode(),
			Message:   responseMessage.GetMessage(),
		},
		MediaInfo: &oburn.MediaInfo{
			Existed:           mediaInfo.GetExisted(),
			IsBlank:           mediaInfo.GetIsBlank(),
			IsCompleted:       mediaInfo.GetIsCompleted(),
			SerialNo:          mediaInfo.GetSerialNo(),
			Mid:               mediaInfo.GetMid(),
			DiscType:          mediaInfo.GetDiscType(),
			WriteSpeed:        mediaInfo.GetWriteSpeed(),
			TotalSize:         mediaInfo.GetTotalSize(),
			FreeSize:          mediaInfo.GetFreeSize(),
			UsedSize:          mediaInfo.GetUsedSize(),
			TrackNum:          mediaInfo.GetTrackNum(),
			TrackStatusList:   mediaInfo.GetTrackStatusList(),
			TrackSizeList:     mediaInfo.GetTrackSizeList(),
			TrackUsedSizeList: mediaInfo.GetTrackUsedSizeList(),
			TrackNwaList:      mediaInfo.GetTrackNwaList(),
			UserDefinedId:     mediaInfo.GetUserDefinedId(),
		},
		DiscFsInfo: &oburn.DiscFsInfo{
			DiscLabel:   discFsInfo.GetDiscLabel(),
			DiscPasswd:  discFsInfo.GetDiscPasswd(),
			FsType:      discFsInfo.GetFsType(),
			MediaStatus: discFsInfo.GetMediaStatus(),
		},
	}, err
}
