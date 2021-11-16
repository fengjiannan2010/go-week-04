package logic

import (
	"context"
	"oburn/rpc/internal/svc"
)

var Burn *BurnLogic
var Check *CheckLogic
var Minio *MinioLogic
var Read *ReadLogic
var Task *TaskLogic

func InitLogic(ctx context.Context, svcCtx *svc.ServiceContext) {
	Burn = NewBurnLogic(ctx, svcCtx)
	Check = NewCheckLogic(ctx, svcCtx)
	Minio = NewMinioLogic(ctx, svcCtx)
	Read = NewReadLogic(ctx, svcCtx)
	Task = NewTaskLogic(ctx, svcCtx)
}
