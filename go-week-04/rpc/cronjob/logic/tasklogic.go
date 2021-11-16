package logic

import (
	"context"
	"github.com/tal-tech/go-zero/core/logx"
	"oburn/data"
	"oburn/rpc/client/odrive"
	"oburn/rpc/internal/svc"
	"sync"
	"time"
)

type TaskLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	once sync.Once
}

func NewTaskLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TaskLogic {
	return &TaskLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
		once:   sync.Once{},
	}
}

func (l *TaskLogic) ResetOnce() {
	l.once = sync.Once{}
}

func (l *TaskLogic) Start() {
	for {
		select {
		case <-l.ctx.Done():
			return
		default:
		}

		realTimeDiscInfo, err := l.GetDiscInfo()
		if err != nil {
			time.Sleep(time.Second * 1)
			continue
		}

		if realTimeDiscInfo.Existed == 0 {
			time.Sleep(time.Second * 1)
			continue
		}

		resp, err := l.svcCtx.Odrive.TestUnityReady(l.ctx, &odrive.EmptyRequest{})
		if err != nil {
			time.Sleep(time.Second * 1)
			continue
		}

		if !resp.IsReady {
			time.Sleep(time.Second * 1)
			continue
		} else {
			discInfo, err := l.GetDiscInfo()
			if err != nil {
				time.Sleep(time.Second * 1)
				continue
			} else {
				l.once.Do(func() {
					//discInfo, err := l.svcCtx.DiscInfoModel.Insert(discInfo)
					//if err != nil {
					//
					//}
					discInfo.ConvertSize()
					l.svcCtx.DiscInfo = discInfo
				})
			}
		}

		notask, _ := Read.ReadTaskJob()
		if !notask {
			time.Sleep(time.Second * 1)
			continue
		}

		notask, _ = Burn.BurnTaskJob()
		if !notask {
			time.Sleep(time.Second * 1)
			continue
		}

		if notask {
			totalCount, err := l.svcCtx.TaskInfoModel.Count(data.TaskInfo{Expired: data.Valid})
			if err != nil {
				time.Sleep(time.Second * 1)
				continue
			}
			totalSuccessCount, err := l.svcCtx.TaskInfoModel.Count(data.TaskInfo{Expired: data.Valid, Status: data.Success})
			if err != nil {
				time.Sleep(time.Second * 1)
				continue
			}

			if totalCount == totalSuccessCount {
				resp, err := l.svcCtx.Odrive.UnLoadDisc(l.ctx, &odrive.EmptyRequest{})
				if err != nil {
					time.Sleep(time.Second * 1)
					continue
				}
				if resp.GetMessage().GetIsSuccess() {
					selectTask := data.TaskInfo{Expired: data.Valid, Status: data.Success}
					updateTask := data.TaskInfo{Expired: data.InValid}
					err := l.svcCtx.TaskInfoModel.UpdateAny(selectTask, updateTask)
					if err != nil {
						time.Sleep(time.Second * 1)
						continue
					}
				}
			}
		}
	}
}

func (l *TaskLogic) GetDiscInfo() (data.DiscInfo, error) {
	discInfoResponse, err := l.svcCtx.Odrive.GetDiscInfo(l.ctx, &odrive.EmptyRequest{})
	mediaInfo := discInfoResponse.GetMediaInfo()
	discFsInfo := discInfoResponse.GetDiscFsInfo()
	return data.DiscInfo{
		Existed:           mediaInfo.GetExisted(),
		IsBlank:           mediaInfo.GetIsBlank(),
		IsCompleted:       mediaInfo.GetIsCompleted(),
		SerialNo:          mediaInfo.GetSerialNo(),
		Mid:               mediaInfo.GetMid(),
		DiscType:          data.DiscType(mediaInfo.GetDiscType()),
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
		DiscLabel:         discFsInfo.GetDiscLabel(),
		DiscPasswd:        discFsInfo.GetDiscPasswd(),
		FsType:            data.DiscFsType(discFsInfo.GetFsType()),
		MediaStatus:       discFsInfo.GetMediaStatus(),
	}, err
}
