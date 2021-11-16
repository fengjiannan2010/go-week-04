package data

import (
	"context"
	"oburn/errorcode"
	"oburn/rpc/client/odriveclient"
)

type BurnEnvironment struct {
	UnitReady        odriveclient.TestUnityReadyResponse
	DriveInfo        odriveclient.DriveInfo
	RealTimeDiscInfo odriveclient.DiscInfoResponse
	DiscInfo         odriveclient.DiscInfoResponse
	CacheInfo        odriveclient.GetCacheInfoResponse
	BurnMode         BurnConfig
	//BurnStatus       oburn.BurnStatus
	CancelFunc      context.CancelFunc
	TaskContext     context.Context
	SystemErrorCode errorcode.OdError
	SystemErrorMsg  string
}

func NewBurnEnvironment() BurnEnvironment {
	return BurnEnvironment{}
}
