package logic

import (
	"context"
	"github.com/pkg/errors"
	"io/fs"
	"oburn/data"
	"oburn/errorcode"
	"oburn/rpc/client/odrive"
	"oburn/rpc/internal/svc"
	"oburn/rpc/oburn"
	"oburn/utils"
	"os"
	"path/filepath"

	"github.com/tal-tech/go-zero/core/logx"
)

type PutObjectLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPutObjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PutObjectLogic {
	return &PutObjectLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PutObjectLogic) PutObject(in *oburn.PutObjectRequest) (*oburn.PutObjectResponse, error) {
	// todo: add your logic here and delete this line
	resp, err := l.svcCtx.Odrive.TestUnityReady(l.ctx, &odrive.EmptyRequest{})
	if err != nil {
		return &oburn.PutObjectResponse{
			Message: &oburn.Message{
				IsSuccess: false,
				Code:      errorcode.OdriveRequestError.Int(),
				Message:   errorcode.OdriveRequestError.String(),
			},
		}, err
	}

	if !resp.IsReady {
		return &oburn.PutObjectResponse{
			Message: &oburn.Message{
				IsSuccess: false,
				Code:      resp.GetMessage().Code,
				Message:   resp.GetMessage().GetMessage(),
			},
		}, err
	}

	taskInfo := data.TaskInfo{
		ObjectName:  in.GetObjectName(),
		BucketName:  in.GetBucketName(),
		DiscPath:    in.GetDiscPath(),
		StorageType: data.StorageType(in.GetStorageType()),
		OffSetStart: in.GetOffset(),
		OffSetEnd:   in.GetOffset() + in.GetCount(),
		CheckCode:   in.GetCheckCode(),
		FilePath:    filepath.Join(in.GetBucketName(), in.GetObjectName()),
	}
	val, err := l.svcCtx.TaskInfoModel.Exists(taskInfo)
	if err != nil {
		return &oburn.PutObjectResponse{
			Message: &oburn.Message{
				IsSuccess: false,
				Code:      errorcode.DbQueryError.Int(),
				Message:   errorcode.DbQueryError.String(),
			},
		}, err
	}

	if val {
		return &oburn.PutObjectResponse{
			Message: &oburn.Message{
				IsSuccess: false,
				Code:      errorcode.BurnFileSamePath.Int(),
				Message:   errorcode.BurnFileSamePath.String(),
			},
		}, errors.New(errorcode.BurnFileSamePath.String())
	}

	err, ecode := l.InitWriteTaskMode(&taskInfo)
	if err != nil {
		return &oburn.PutObjectResponse{
			Message: &oburn.Message{
				IsSuccess: false,
				Code:      ecode.Int(),
				Message:   ecode.String(),
			},
		}, err
	}
	totalSize := taskInfo.SyncDiscCapacity()
	if l.svcCtx.DiscInfo.FreeSize < totalSize {
		return &oburn.PutObjectResponse{
			Message: &oburn.Message{
				IsSuccess: false,
				Code:      errorcode.InsufficientDiskCpacity.Int(),
				Message:   errorcode.InsufficientDiskCpacity.String(),
			},
		}, errors.New(errorcode.InsufficientDiskCpacity.String())
	}

	_, err = l.svcCtx.TaskInfoModel.Insert(taskInfo)
	if err != nil {
		return &oburn.PutObjectResponse{
			Message: &oburn.Message{
				IsSuccess: false,
				Code:      errorcode.DbCreateTaskError.Int(),
				Message:   errorcode.DbCreateTaskError.String(),
			},
		}, err
	}
	l.svcCtx.DiscInfo.SyncCapacity(totalSize)
	return &oburn.PutObjectResponse{
		Message: &oburn.Message{
			IsSuccess: true,
			Code:      errorcode.NoError.Int(),
			Message:   errorcode.NoError.String(),
		},
	}, nil
}

// InitWriteTaskMode TODO 创建刻录任务
func (l *PutObjectLogic) InitWriteTaskMode(cmd *data.TaskInfo) (error, errorcode.ObError) {
	if cmd.StorageType == data.Local {
		if cmd.FileType == data.File {
			val, err := utils.PathExists(cmd.FilePath)
			if err != nil || !val {
				return err, errorcode.FileIsNotExist
			}
			cmd.UUID = cmd.DiscPath
			cmd.Status = data.Waiting
			cmd.DiscMode = data.DiscWrite
			cmd.Expired = data.Valid
			info, err := os.Stat(cmd.FilePath)
			if err == nil {
				cmd.FileSize = info.Size()
			}
			return nil, errorcode.NoError
		} else if cmd.FileType == data.Directory {
			val, err := utils.PathExists(cmd.FilePath)
			if err != nil || !val {
				return err, errorcode.FileIsNotExist
			}
			cmd.UUID = cmd.DiscPath
			cmd.Status = data.Waiting
			cmd.DiscMode = data.DiscWrite
			cmd.Expired = data.Valid

			discDirPath := cmd.DiscPath
			err = filepath.WalkDir(cmd.FilePath, func(path string, info fs.DirEntry, err error) error {
				rel, err := filepath.Rel(cmd.FilePath, path)
				if err != nil {
					return err
				}
				if rel == "." {
					return nil
				}
				discFilePath := filepath.ToSlash(filepath.Join(discDirPath, rel))
				task := data.TaskInfo{}
				task.FilePath = path
				task.DiscPath = discFilePath
				task.StorageType = data.Local
				task.UUID = discFilePath
				task.Status = data.Waiting
				task.DiscMode = data.DiscWrite
				task.Expired = data.Valid
				if info.IsDir() {
					task.FileType = data.Directory
					task.FileSize = 0
				} else {
					fileInfo, _ := info.Info()
					task.FileType = data.File
					task.FileSize = fileInfo.Size()
				}
				cmd.SubTasks = append(cmd.SubTasks, task)
				return nil
			})
			if err != nil {
				return err, errorcode.FileWalkError
			}
		}
	} else if cmd.StorageType == data.MinIo {

	} else {

	}
	return nil, errorcode.NoError
}
