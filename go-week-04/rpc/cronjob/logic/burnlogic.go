package logic

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/tal-tech/go-zero/core/logx"
	"hash/crc32"
	"io"
	"oburn/data"
	"oburn/errorcode"
	"oburn/rpc/client/odrive"
	"oburn/rpc/internal/svc"
	"oburn/utils"
	"oburn/utils/fileattribute"
	"os"
	"path/filepath"
	"syscall"
	"time"
)

type BurnLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBurnLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BurnLogic {
	return &BurnLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BurnLogic) BurnTaskJob() (bool, error) {
	count, err := l.svcCtx.TaskInfoModel.Count(data.TaskInfo{Expired: data.Valid})
	if err != nil {
		return false, errors.Wrap(err, errorcode.DbQueryError.String())
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
		task, err := l.svcCtx.TaskInfoModel.FindOneByStatus(data.TaskInfo{DiscMode: data.DiscWrite, Status: data.AppendPart, Expired: data.Valid})
		if err != nil {
			task, err = l.svcCtx.TaskInfoModel.FindOneByStatus(data.TaskInfo{DiscMode: data.DiscWrite, Status: data.BurningFail, Expired: data.Valid})
		}
		if err != nil {
			task, err = l.svcCtx.TaskInfoModel.FindOneByStatus(data.TaskInfo{DiscMode: data.DiscWrite, Status: data.Waiting, Expired: data.Valid})
		}
		if err != nil {
			return true, err
		}

		if task.Status == data.Waiting {
			l.BurnFile(task, int64(blockSize))
		} else if task.Status == data.AppendPart {
			l.AppendPartFile(task, int64(blockSize))
		}
	}
	return true, nil
}

func (l *BurnLogic) BurnFile(task *data.TaskInfo, blockSize int64) (int32, error) {
	ecode, err := l.SetBurnMode()
	if err != nil {
		return ecode, err
	}
	task.Status = data.Burning
	l.svcCtx.TaskInfoModel.Update(*task)
	if task.FileType == data.File {
		dataFile := DataFile{}
		dataFile.OffSetStart = task.OffSetStart
		dataFile.OffSetEnd = task.OffSetEnd
		if task.StorageType == data.Local {
			file, err := os.Open(task.FilePath)
			if err != nil {
				task.Status = data.Fail
				task.Error = err.Error()
				task.ErrorCode = errorcode.FileOpenError.Int()
				l.svcCtx.TaskInfoModel.Update(*task)
				return 0, errors.Wrap(err, "open file error")
			}
			fileInfo, err := file.Stat()
			if err != nil {
				task.Status = data.Fail
				task.Error = err.Error()
				task.ErrorCode = errorcode.FileStatError.Int()
				l.svcCtx.TaskInfoModel.Update(*task)
				return 0, errors.Wrap(err, "file stat error")
			}
			dataFile.FileReader = file
			dataFile.FileInfo = fileInfo
			dataFile.DiscPath = task.DiscPath
		} else if task.StorageType == data.MinIo {
			file, err := Minio.GetObjectStream(task.BucketName, task.ObjectName)
			if err != nil {
				task.Status = data.Fail
				task.Error = err.Error()
				task.ErrorCode = errorcode.GetObjectStreamError.Int()
				l.svcCtx.TaskInfoModel.Update(*task)
				return 0, errors.Wrap(err, "get ogject stream error")
			}
			objectInfo, err := file.Stat()
			if err != nil {
				task.Status = data.Fail
				task.Error = err.Error()
				task.ErrorCode = errorcode.ObjectStatError.Int()
				l.svcCtx.TaskInfoModel.Update(*task)
				return 0, errors.Wrap(err, "获取文件属性失败")
			}
			fileInfo := NewObjectFileInfo(task.ObjectName, objectInfo.Size, objectInfo.LastModified)
			dataFile.FileReader = file
			dataFile.FileInfo = fileInfo
			dataFile.DiscPath = task.DiscPath
		}
		defer dataFile.FileReader.Close()
		filecrc32, ecode, err := l.PubObject(task, dataFile, blockSize)
		if err != nil {
			return ecode, err
		}
		task.FileCrc32 = filecrc32
	} else if task.FileType == data.Directory {
		response, err := l.svcCtx.Odrive.MakeBucket(l.ctx, &odrive.MakeBucketRequest{
			FilePathHash: utils.MD5V([]byte(filepath.ToSlash(task.DiscPath))),
		})
		if !response.GetMessage().GetIsSuccess() {
			return response.GetMessage().GetCode(), err
		}
	}
	task.Status = data.CheckPart
	l.svcCtx.TaskInfoModel.Update(*task)
	return 0, nil
}

func (l *BurnLogic) AppendPartFile(task *data.TaskInfo, blockSize int64) (int32, error) {
	ecode, err := l.SetBurnMode()
	if err != nil {
		return ecode, err
	}
	task.Status = data.Burning
	l.svcCtx.TaskInfoModel.Update(*task)
	if task.FileType == data.File {
		dataFile := DataFile{}
		if task.StorageType == data.Local {
			file, err := os.Open(task.FilePath)
			if err != nil {
				task.Status = data.Fail
				task.Error = err.Error()
				task.ErrorCode = errorcode.FileOpenError.Int()
				l.svcCtx.TaskInfoModel.Update(*task)
				return 0, errors.Wrap(err, "open file error")
			}
			fileInfo, err := file.Stat()
			if err != nil {
				task.Status = data.Fail
				task.Error = err.Error()
				task.ErrorCode = errorcode.FileStatError.Int()
				l.svcCtx.TaskInfoModel.Update(*task)
				return 0, errors.Wrap(err, "file stat error")
			}
			dataFile.FileReader = file
			dataFile.FileInfo = fileInfo
			dataFile.DiscPath = task.DiscPath
		} else if task.StorageType == data.MinIo {
			file, err := Minio.GetObjectStream(task.BucketName, task.ObjectName)
			if err != nil {
				task.Status = data.Fail
				task.Error = err.Error()
				task.ErrorCode = errorcode.GetObjectStreamError.Int()
				l.svcCtx.TaskInfoModel.Update(*task)
				return 0, errors.Wrap(err, "get ogject stream error")
			}
			objectInfo, err := file.Stat()
			if err != nil {
				task.Status = data.Fail
				task.Error = err.Error()
				task.ErrorCode = errorcode.ObjectStatError.Int()
				l.svcCtx.TaskInfoModel.Update(*task)
				return 0, errors.Wrap(err, "获取文件属性失败")
			}
			fileInfo := NewObjectFileInfo(task.ObjectName, objectInfo.Size, objectInfo.LastModified)
			dataFile.FileReader = file
			dataFile.FileInfo = fileInfo
			dataFile.DiscPath = task.DiscPath
		}
		defer dataFile.FileReader.Close()

		parts, err := l.svcCtx.FilePartInfoModel.FindAny(task.Id)
		if err != nil {
			return 0, err
		}
		for _, part := range *parts {
			dataFile.OffSetStart = part.OffSetStart
			dataFile.OffSetEnd = part.OffSetEnd
			l.PubObject(task, dataFile, blockSize)
		}
	} else if task.FileType == data.Directory {
		response, err := l.svcCtx.Odrive.MakeBucket(l.ctx, &odrive.MakeBucketRequest{
			FilePathHash: utils.MD5V([]byte(filepath.ToSlash(task.DiscPath))),
		})
		if !response.GetMessage().GetIsSuccess() {
			return response.GetMessage().GetCode(), err
		}
	}
	task.Status = data.CheckPart
	l.svcCtx.TaskInfoModel.Update(*task)
	return 0, nil
}

func (l *BurnLogic) SetBurnMode() (int32, error) {
	burnMode := odrive.SetBurnModeResquest{}
	if l.svcCtx.BurnConfig.IsVerify {
		burnMode.Mode = l.svcCtx.Config.VerifyMode
	} else {
		burnMode.Mode = int32(data.AfterVerify)
	}
	response, err := l.svcCtx.Odrive.SetBurnMode(l.ctx, &burnMode)
	if err != nil {
		return 0, err
	}
	return response.GetMessage().GetCode(), nil
}

func (l *BurnLogic) PubObject(task *data.TaskInfo, dataFile DataFile, blockSize int64) (uint32, int32, error) {
	startTime := time.Now()
	crcTable := crc32.IEEETable
	fileCrc32Val := crc32.ChecksumIEEE(nil)
	reader, err := l.svcCtx.Odrive.PutObject(l.ctx)
	if err != nil {
		return fileCrc32Val, 0, err
	}
	offSet := dataFile.OffSetStart
	totalSize := dataFile.OffSetEnd
	if totalSize == 0 {
		totalSize = dataFile.FileInfo.Size()
	}
	for i := int64(offSet); i < totalSize; {
		select {
		case <-l.ctx.Done():
			break
		default:
		}
		if totalSize-i < blockSize {
			blockSize = totalSize - i
		}

		request := odrive.PutObjectRequest{}
		request.DataBuffer = make([]byte, blockSize)
		dataFile.FileReader.Seek(i, 0)
		partSize, err := io.ReadFull(dataFile.FileReader, request.DataBuffer)
		if err != nil {
			if err != io.EOF && err != io.ErrUnexpectedEOF {
				return fileCrc32Val, errorcode.OdErrFileReadIoException.Int(), err
			}
		}
		crc32Code := crc32.ChecksumIEEE(request.DataBuffer)
		tmpCrc32 := crc32.Update(fileCrc32Val, crcTable, request.DataBuffer)
		request.BlockInfo = &odrive.BlockInfo{
			FileName:          filepath.ToSlash(filepath.Base(dataFile.DiscPath)),
			FilePathHash:      utils.MD5V([]byte(filepath.ToSlash(dataFile.DiscPath))),
			DirectoryName:     filepath.ToSlash(filepath.Base(filepath.Dir(dataFile.DiscPath))),
			DirectoryPathHash: utils.MD5V([]byte(filepath.ToSlash(filepath.Dir(dataFile.DiscPath)))),
			FileType:          int32(task.FileType),
			CreateTime:        fileattribute.CreateTime(dataFile.FileInfo).Unix(),
			FileMode:          uint32(os.ModePerm | syscall.S_IFREG),
			PartSize:          int32(partSize),
			FileSize:          dataFile.FileInfo.Size(),
			OffsetStart:       i,
			OffsetEnd:         i + int64(partSize),
			CheckCode:         tmpCrc32,
		}
		err = reader.Send(&request)
		if err != nil {
			return fileCrc32Val, 0, err
		}
		fileCrc32Val = tmpCrc32
		partInfo := data.FilePartInfo{
			TaskId:      task.Id,
			UUID:        fmt.Sprintf("%d_%d", request.BlockInfo.OffsetStart, request.BlockInfo.OffsetEnd),
			FilePath:    filepath.ToSlash(dataFile.DiscPath),
			FileType:    data.File,
			Crc32:       int32(crc32Code),
			OffSetStart: request.BlockInfo.OffsetStart,
			OffSetEnd:   request.BlockInfo.OffsetEnd,
		}
		_, err = l.svcCtx.FilePartInfoModel.Insert(partInfo)
		if err != nil {
			continue
		}
		i += int64(partSize)
		percent := int(float64(i) / float64(totalSize) * 100.00)
		task.BurnProgress = int32(percent)
	}
	response, err := reader.CloseAndRecv()
	if err != nil {
		return fileCrc32Val, response.GetMessage().GetCode(), err
	}
	endTime := time.Now()
	totalSec := endTime.Sub(startTime).Seconds()
	mb := float64(totalSize) / 1024.00 / 1024.00
	message := fmt.Sprintf("totalsize:%d,totalsec:%f,%f MB/s", totalSize, totalSec, mb/totalSec)
	l.Logger.Info(message)
	return fileCrc32Val, response.GetMessage().GetCode(), nil
}
