package logic

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/tal-tech/go-zero/core/logx"
	"io"
	"oburn/data"
	"oburn/errorcode"
	"oburn/rpc/client/odrive"
	"oburn/rpc/internal/svc"
	"oburn/utils"
	"os"
	"path/filepath"
	"time"
)

type ReadLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewReadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReadLogic {
	return &ReadLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ReadLogic) ReadTaskJob() (bool, error) {
	count, err := l.svcCtx.TaskInfoModel.Count(data.TaskInfo{Expired: data.Valid})
	if err != nil {
		return false, err
	}
	if count > 0 {
		response, err := l.svcCtx.Odrive.GetCacheInfo(l.ctx, &odrive.EmptyRequest{})
		if err != nil {
			return false, err
		}
		if !response.GetMessage().GetIsSuccess() {

		}
		if response.GetWriteErrorCode() != errorcode.OdNoError.Int() {
			if response.GetWriteErrorCode() == errorcode.OdDataBaseError.Int() {
				l.svcCtx.Odrive.Reset(l.ctx, &odrive.EmptyRequest{})
			}
		}
		blockSize := response.GetWriteBlockSize()
		task, err := l.svcCtx.TaskInfoModel.FindOneByStatus(data.TaskInfo{DiscMode: data.DiscRead, Status: data.ReadingFail, Expired: data.Valid})
		if err != nil {
			task, err = l.svcCtx.TaskInfoModel.FindOneByStatus(data.TaskInfo{DiscMode: data.DiscRead, Status: data.Waiting, Expired: data.Valid})
		}
		if err != nil {
			return true, err
		}
		if task.Status == data.Waiting {
			l.ReadFile(task, int64(blockSize))
		}
	}
	return false, nil
}

func (l *ReadLogic) ReadFile(task *data.TaskInfo, blockSize int64) (int32, error) {
	ecode, err := l.SetReadMode()
	if err != nil {
		return ecode, err
	}
	task.Status = data.Reading
	l.svcCtx.TaskInfoModel.Update(*task)
	if task.FileType == data.File {
		dataFile := DataFile{}
		dataFile.OffSetStart = task.OffSetStart
		dataFile.OffSetEnd = task.OffSetEnd
		if task.StorageType == data.Local {
			destDir := filepath.Dir(task.FilePath)
			if !utils.FileExists(destDir) {
				err = utils.Mkdir(destDir, 0755)
				if err != nil {
					return 0, err
				}
			}
			file, err := os.OpenFile(task.FilePath, os.O_CREATE|os.O_WRONLY, 0755)
			if err != nil {
				task.Status = data.Fail
				task.Error = err.Error()
				task.ErrorCode = errorcode.CreateFileError.Int()
				l.svcCtx.TaskInfoModel.Update(*task)
				return 0, errors.Wrap(err, "create file error")
			}
			dataFile.FileWriter = file
			fileInfo := NewObjectFileInfo(file.Name(), task.OffSetEnd, time.Now())
			dataFile.FileInfo = fileInfo
			dataFile.DiscPath = task.DiscPath
		} else if task.StorageType == data.MinIo {

		}
		defer dataFile.FileWriter.Close()
		filecrc32, ecode, err := l.GetObject(task, dataFile, blockSize)
		if err != nil {
			return ecode, err
		}
		task.FileCrc32 = filecrc32
	} else if task.FileType == data.Directory {
		if !utils.FileExists(task.FilePath) {
			err = utils.Mkdir(task.FilePath, 0755)
			if err != nil {
				return 0, err
			}
		}
	} else {
		return 0, err
	}
	task.Status = data.ReadingSuccess
	l.svcCtx.TaskInfoModel.Update(*task)
	return 0, nil
}

func (l *ReadLogic) SetReadMode() (int32, error) {
	burnMode := odrive.SetBurnModeResquest{}
	burnMode.Mode = data.ReadOnly.Int()
	response, err := l.svcCtx.Odrive.SetBurnMode(l.ctx, &burnMode)
	if err != nil {
		return 0, err
	}
	return response.GetMessage().GetCode(), nil
}

func (l *ReadLogic) GetObject(task *data.TaskInfo, dataFile DataFile, blockSize int64) (uint32, int32, error) {
	startTime := time.Now()
	//offSet := dataFile.OffSetStart
	totalSize := dataFile.OffSetEnd
	if totalSize == 0 {
		totalSize = dataFile.FileInfo.Size()
	}
	request := odrive.GetObjectRequest{
		//FilePathHash: utils.MD5V([]byte(filepath.ToSlash(dataFile.DiscPath))),
		//FilePath:     dataFile.DiscPath,
		//OffSetStart:  offSet,
		//OffSetEnd:    totalSize,
	}
	reader, err := l.svcCtx.Odrive.GetObject(l.ctx, &request)
	if err != nil {
		return 0, 0, err
	}
	readSize := int64(0)
	for {
		select {
		case <-l.ctx.Done():
			break
		default:
		}
		response, err := reader.Recv()
		if err != nil {
			if err == io.EOF {
				endTime := time.Now()
				totalSec := endTime.Sub(startTime).Seconds()
				mb := float64(readSize) / 1024.00 / 1024.00
				message := fmt.Sprintf("totalsize:%d,totalsec:%f,%f MB/s", totalSize, totalSec, mb/totalSec)
				l.Logger.Info(message)
				return 0, 0, nil
			}
			return 0, 0, err
		}
		if response.GetMessage().GetIsSuccess() {
			//dataFile.FileWriter.Seek(response.GetOffset()-offSet, 0)
			//crc32Code := crc32.ChecksumIEEE(response.GetDataBuffer())
			//if crc32Code != uint32(response.GetCheckCode()) {
			//
			//}
			count, err := dataFile.FileWriter.Write(response.GetDataBuffer())
			if err != nil {
				return 0, 0, err
			}
			readSize += int64(count)
		} else {
			return 0, response.GetMessage().GetCode(), errors.New("read file error")
		}
	}
}
