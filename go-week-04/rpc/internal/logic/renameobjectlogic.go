package logic

import (
	"context"
	"oburn/rpc/client/odrive"
	"oburn/utils"
	"path/filepath"

	"oburn/rpc/internal/svc"
	"oburn/rpc/oburn"

	"github.com/tal-tech/go-zero/core/logx"
)

type RenameObjectLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRenameObjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RenameObjectLogic {
	return &RenameObjectLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RenameObjectLogic) RenameObject(in *oburn.RenameObjectRequest) (*oburn.RenameObjectResponse, error) {
	// todo: add your logic here and delete this line
	response, err := l.svcCtx.Odrive.RenameObject(l.ctx, &odrive.RenameObjectRequest{
		FilePathHash: utils.MD5V([]byte(filepath.ToSlash(in.GetOldDiscPath()))),
		BlockInfo: &odrive.BlockInfo{
			FileName:          filepath.ToSlash(filepath.Base(in.GetNewDiscPath())),
			FilePathHash:      utils.MD5V([]byte(filepath.ToSlash(in.GetNewDiscPath()))),
			DirectoryName:     filepath.ToSlash(filepath.Base(filepath.Dir(in.GetNewDiscPath()))),
			DirectoryPathHash: utils.MD5V([]byte(filepath.ToSlash(filepath.Dir(in.GetNewDiscPath())))),
		},
	})
	responseMessage := response.GetMessage()
	return &oburn.RenameObjectResponse{
		Message: &oburn.Message{
			IsSuccess: responseMessage.GetIsSuccess(),
			Code:      responseMessage.GetCode(),
			Message:   responseMessage.GetMessage(),
		},
	}, err
}
