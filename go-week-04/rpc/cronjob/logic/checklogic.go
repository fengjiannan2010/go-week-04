package logic

import (
	"context"
	"github.com/tal-tech/go-zero/core/logx"
	"oburn/data"
	"oburn/rpc/client/odrive"
	"oburn/rpc/internal/svc"
	"time"
)

type CheckLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckLogic {
	return &CheckLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckLogic) CheckFilePartJob(second int) {
	for {
		select {
		case <-l.ctx.Done():
			return
		default:
		}
		task, err := l.svcCtx.TaskInfoModel.FindOneByStatus(data.TaskInfo{Status: data.CheckPart, Expired: data.Valid})
		if err != nil {
			time.Sleep(time.Second * time.Duration(second))
		}
		l.CheckTask(task)
	}
}

func (l *CheckLogic) CheckTask(taskInfo *data.TaskInfo) error {
	for {
		select {
		case <-l.ctx.Done():
			return nil
		default:
		}
		time.Sleep(time.Second * 1)
		parts, err := l.svcCtx.FilePartInfoModel.FindAny(taskInfo.Id)
		if err != nil {
			continue
		}
		for _, part := range *parts {
			request := odrive.ObjectExistsRequest{
				FilePathHash: part.FileNameHash,
				FilePath:     part.FilePath,
				FileType:     int32(part.FileType),
				OffsetStart:  part.OffSetStart,
				OffsetEnd:    part.OffSetEnd,
				CheckCode:    part.Crc32,
			}
			response, err := l.svcCtx.Odrive.ObjectExists(l.ctx, &request)
			if err != nil {
				continue
			}
			if response.GetIsExists() && response.GetStatus() == data.Disc.Int() {
				l.svcCtx.FilePartInfoModel.Delete(part.Id)
			}
			// else {
			//	part.IsExists = response.GetIsExists()
			//	part.Status = response.GetStatus()
			//	l.svcCtx.FilePartInfoModel.Update(part)
			//}
		}

		parts, err = l.svcCtx.FilePartInfoModel.FindAny(taskInfo.Id)
		if err != nil {
			continue
		}
		if len(*parts) == 0 {
			taskInfo.Status = data.BurningSuccess
			if l.svcCtx.BurnConfig.IsVerify {
				if l.svcCtx.Config.VerifyMode == data.AfterVerify.Int() {
					taskInfo.Status = data.Waiting
					taskInfo.DiscMode = data.DiscVerify
					l.svcCtx.TaskInfoModel.Update(*taskInfo)
				} else if l.svcCtx.Config.VerifyMode == data.SyncVerify.Int() {
					taskInfo.Status = data.Success
					l.svcCtx.TaskInfoModel.Update(*taskInfo)
				}
			} else {
				taskInfo.Status = data.Success
				l.svcCtx.TaskInfoModel.Update(*taskInfo)
			}
			break
		} else {
			append := false
			for _, part := range *parts {
				if part.IsExists && part.Status == data.None {
					append = true
				} else {
					append = false
				}
			}
			if append {
				taskInfo.Status = data.AppendPart
				taskInfo.Retries = taskInfo.Retries + 1
				l.svcCtx.TaskInfoModel.Update(*taskInfo)
				break
			}
		}
	}
	return nil
}

func (l *CheckLogic) ObjectExists(request *odrive.ObjectExistsRequest) (*odrive.ObjectExistsResponse, error) {
	response, err := l.svcCtx.Odrive.ObjectExists(l.ctx, request)
	return response, err
}
