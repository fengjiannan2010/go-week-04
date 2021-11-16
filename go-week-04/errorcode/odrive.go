package errorcode

type OdError int

const (
	OdNoError                  OdError = 0x0000
	OdErrCachefull             OdError = 0x0002
	OdErrFrameformat           OdError = 0x0003
	OdErrUnknowcmd             OdError = 0x0004
	OdErrNotconfig             OdError = 0x0005
	OdErrNotready              OdError = 0x0006
	OdErrTaskclose             OdError = 0x0007
	OdErrFramelentoolong       OdError = 0x0008
	OdErrConnectbroken         OdError = 0x0009
	OdErrUnknown               OdError = 0x000A
	OdErrRecevtimeout          OdError = 0x000B
	OdErrInternalerr           OdError = 0x000C
	OdErrParmerr               OdError = 0x000D
	OdErrDbqueryerr            OdError = 0x000E
	OdErrExtfilesize           OdError = 0x000F
	OdErrConfigureerror        OdError = 0x0010
	OdErrConnectodataerror     OdError = 0x0011
	OdErrCmdformaterror        OdError = 0x0012
	OdErrSenddatatoodataerror  OdError = 0x0013
	OdErrVerifytimeout         OdError = 0x0014
	OdErrNoverifyitem          OdError = 0x0015
	OdErrPatherror             OdError = 0x0016
	OdErrGpioerror             OdError = 0x0017
	OdErrGpiotrayopenerror     OdError = 0x0018
	OdErrGpiotraycloseerror    OdError = 0x0019
	OD_Err_RebootInit          OdError = 0x001A
	OdErrReadioexception       OdError = 0x001B
	OdErrChunkdatacrcexception OdError = 0x001C
	OdErrNotverifymode         OdError = 0x001D
	OD_Err_Busy                OdError = 0x001E
	OdErrDiscnotblank          OdError = 0x001F
	OdErrVerify                OdError = 0x0020
	OD_Err_FromOdataFrame      OdError = 0x0021
	OD_Err_FileIoException     OdError = 0xFFFF01
	OD_Err_FileStateException  OdError = 0xFFFF02
	OdErrFileReadIoException   OdError = 0xFFFF03
	OdErrSocketsendexception   OdError = 0xFFFF04
	OD_Err_NoDiscOrUnsupported OdError = 0xFFFF05
	OD_Err_MkFileException     OdError = 0xFFFF06
	OdDataBaseError            OdError = 0x01FF0013
	DiscCapacityShortage       OdError = 0x01FF0037
)

func (err OdError) String() string {
	switch err {
	case OdNoError:
		return ""
	default:
		return ""
	}
}

func (err OdError) Uint32() uint32{
	return uint32(err)
}

func (err OdError) Uint64() uint64{
	return uint64(err)
}

func (err OdError) Int() int32 {
	return int32(err)
}

//func ErrToStr(errCode int) (errStr string) {
//	switch errCode {
//	case OdNoError:
//		errStr = "正确"
//	case OdErrCachefull:
//		errStr = "缓冲区满"
//	case OdErrFrameformat:
//		errStr = "帧格式错误"
//	case OdErrUnknowcmd:
//		errStr = "无效指令"
//	case OdErrNotconfig:
//		errStr = "发送本指令之前，未发送设置指令"
//	case OdErrNotready:
//		errStr = "未准备号"
//	case OdErrTaskclose:
//		errStr = "任务已关闭"
//	case OdErrFramelentoolong:
//		errStr = "帧数据长度太长"
//	case OdErrConnectbroken:
//		errStr = "连接被断开"
//	case OdErrUnknown:
//		errStr = "未知错误"
//	case OdErrRecevtimeout:
//		errStr = "从 oData接收数据超时"
//	case OdErrInternalerr:
//		errStr = "内部错误"
//	case OdErrParmerr:
//		errStr = "参数错误"
//	case OdErrDbqueryerr:
//		errStr = "数据库访问错误"
//	case OdErrExtfilesize:
//		errStr = "访问时大小超过文件大小"
//	case OdErrConfigureerror:
//		errStr = "获取配置信息错误"
//	case OdErrConnectodataerror:
//		errStr = "连接 oData 服务失败"
//	case OdErrCmdformaterror:
//		errStr = "指令格式错误"
//	case OdErrSenddatatoodataerror:
//		errStr = "向 oData 发送数据失败"
//	case OdErrVerifytimeout:
//		errStr = "校验超时"
//	case OdErrNoverifyitem:
//		errStr = "未发现需要校验的有效信息"
//	case OdErrPatherror:
//		errStr = "路径错误"
//	case OdErrGpioerror:
//		errStr = "GPIO 错误"
//	case OdErrGpiotrayopenerror:
//		errStr = "打开托盘失败"
//	case OdErrGpiotraycloseerror:
//		errStr = "关闭托盘失败"
//	case OdErrReadioexception:
//		errStr = "读取文件流失败"
//	case OdErrChunkdatacrcexception:
//		errStr = "分片数据CRC32错误"
//	case OdErrNotverifymode:
//		errStr = "未设置校验模式"
//	case OdErrDiscnotblank:
//		errStr = "刻录光盘非空"
//	case OdErrVerify:
//		errStr = "校验错误"
//	default:
//		errStr = fmt.Sprintf("未知错误：%X", errCode)
//	}
//	return errStr
//}
