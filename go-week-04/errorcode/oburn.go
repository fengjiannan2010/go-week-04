package errorcode

type ObError int

const (
	NoError ObError = 0x000000
	Error   ObError = iota + 0x1A0000
	NotResponse
	ResponseJsonFormatError
	ResponseStatusError
	InterfaceToStructError
	FileIsNotExist
	FileExist
	FileWalkError
	InsufficientDiskCpacity
	MkDirError
	DiscModeError
	DbCreateTaskError
	DbCreateDiscConfigError
	BurnFileSamePath
	ReadFileSamePath
	DbQueryError
	DbQueryFilePartError
	FileOpenError
	CreateFileError
	FileStatError
	GetObjectStreamError
	ObjectStatError
	OdriveRequestError
)

func (err ObError) String() string {
	switch err {
	case NoError:
		return "NoError"
	case FileIsNotExist:
		return "FileIsNotExist"
	default:
		return "UNKNOWN"
	}
}

func (err ObError) Int() int32 {
	return int32(err)
}
