package logic

import (
	"context"

	"oburn/rpc/internal/svc"
	"oburn/rpc/oburn"

	"github.com/tal-tech/go-zero/core/logx"
)

type RecoverDiscLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRecoverDiscLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RecoverDiscLogic {
	return &RecoverDiscLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RecoverDiscLogic) RecoverDisc(in *oburn.EmptyRequest) (*oburn.RecoverDiscResponse, error) {
	// todo: add your logic here and delete this line

	return &oburn.RecoverDiscResponse{}, nil
}
