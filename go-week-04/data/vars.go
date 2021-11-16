package data

import (
	"database/sql/driver"
	"fmt"
)

const BlockSize int64 = 2048

type PartStatus int32

const (
	None  PartStatus = 1
	Disc  PartStatus = 2
	Cache PartStatus = 3
)

func (s PartStatus) String() string {
	switch s {
	case None:
		return "None"
	case Disc:
		return "Disc"
	case Cache:
		return "Cache"
	default:
		return "UNKNOWN"
	}
}

func (s PartStatus) Int() int32 {
	return int32(s)
}

func (s PartStatus) Int64() int64 {
	return int64(s)
}

type TaskStatus int32

const (
	Waiting TaskStatus = iota + 1
	Burning
	BurningSuccess
	BurningFail
	AppendPart
	CheckPart
	Reading
	ReadingSuccess
	ReadingFail
	Verifying
	VerifyingSuccess
	VerifyingFail
	Success
	Fail
)

func (s TaskStatus) String() string {
	switch s {
	case Waiting:
		return "Waiting"
	case Burning:
		return "Burning"
	case Fail:
		return "Fail"
	default:
		return "UNKNOWN"
	}
}

func (s TaskStatus) Int() int32 {
	return int32(s)
}

func (s TaskStatus) Int64() int64 {
	return int64(s)
}

type RecordMode int32

const (
	NoSet       RecordMode = 0
	SyncVerify  RecordMode = 1
	NoVerify    RecordMode = 2
	ReadOnly    RecordMode = 4
	VerifyOnly  RecordMode = 5
	ImageVerify RecordMode = 7

	AfterVerify RecordMode = 3
	Idle        RecordMode = 6
)

func (b RecordMode) String() string {
	switch b {
	case NoSet:
		return "NoSet"
	case SyncVerify:
		return "SyncVerify"
	case NoVerify:
		return "NoVerify"
	case AfterVerify:
		return "AfterVerify"
	case ReadOnly:
		return "ReadOnly"
	case VerifyOnly:
		return "VerifyOnly"
	case Idle:
		return "Idle"
	case ImageVerify:
		return "ImageVerify"
	default:
		return "NoSet"
	}
}

func (s RecordMode) Int() int32 {
	return int32(s)
}

func (s RecordMode) Int64() int64 {
	return int64(s)
}

type DiscMode int32

const (
	DiscRead   DiscMode = 1
	DiscWrite  DiscMode = 2
	DiscVerify DiscMode = 3
)

func (d DiscMode) String() string {
	switch d {
	case DiscRead:
		return "DiscRead"
	case DiscWrite:
		return "DiscWrite"
	case DiscVerify:
		return "DiscVerify"
	default:
		return "UNKNOWN"
	}
}

func (d DiscMode) Int() int32 {
	return int32(d)
}

func (d DiscMode) Int64() int64 {
	return int64(d)
}

type ExpiredType int32

const (
	Valid   ExpiredType = 1
	InValid ExpiredType = 2
)

func (s ExpiredType) String() string {
	switch s {
	case Valid:
		return "Valid"
	case InValid:
		return "InValid"
	default:
		return "UNKNOWN"
	}
}

func (s ExpiredType) Int() int32 {
	return int32(s)
}

func (s ExpiredType) Int64() int64 {
	return int64(s)
}

type FileType int32

const (
	File         FileType = 0
	Directory    FileType = 1
	IsoImageInfo FileType = 2
)

func (s FileType) String() string {
	switch s {
	case File:
		return "FileReader"
	case Directory:
		return "Directory"
	default:
		return "UNKNOWN"
	}
}

func (s FileType) Int() int32 {
	return int32(s)
}

func (s FileType) Int64() int64 {
	return int64(s)
}

type StorageType int32

const (
	Local StorageType = 0
	MinIo StorageType = 1
)

func (s StorageType) String() string {
	switch s {
	case Local:
		return "Local"
	case MinIo:
		return "MinIo"
	default:
		return "UNKNOWN"
	}
}

func (s StorageType) Int() int32 {
	return int32(s)
}

func (s StorageType) Int64() int64 {
	return int64(s)
}


type DiscFsType int

const (
	NotSet     DiscFsType = 0
	ImageBurn  DiscFsType = 1
	StreamBurn DiscFsType = 2
	UdfBurn    DiscFsType = 3
	UnKnown    DiscFsType = 6
	ExtImage   DiscFsType = 7
)

func (r DiscFsType) String() string {
	switch r {
	case NotSet:
		return "NotSet"
	case ImageBurn:
		return "ImageBurn"
	case StreamBurn:
		return "StreamBurn"
	case UdfBurn:
		return "UdfBurn"
	case UnKnown:
		return "UNKNOWN"
	case ExtImage:
		return "ExtImage"
	default:
		return "UNKNOWN"
	}
}

func (s DiscFsType) Int() int {
	return int(s)
}

func (s DiscFsType) Int32() int32 {
	return int32(s)
}

func (s DiscFsType) Int64() int64 {
	return int64(s)
}

func (s DiscFsType) Value() (driver.Value, error) {
	return s.Int64(), nil
}

func (s *DiscFsType) Scan(v interface{}) error {
	value, ok := v.(int64)
	if ok {
		*s = DiscFsType(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

type DiscType int32

const (
	DISC_TYPE_CD           DiscType = 1
	DISC_TYPE_DVD          DiscType = 2
	DISC_TYPE_BD           DiscType = 3
	DISC_TYPE_NO_MEDIA     DiscType = 0x02
	DISC_TYPE_CDROM        DiscType = 0x08
	DISC_TYPE_CDR          DiscType = 0x09
	DISC_TYPE_CDRW         DiscType = 0x0A
	DISC_TYPE_DVDROM       DiscType = 0x10
	DISC_TYPE_DVDR         DiscType = 0x11
	DISC_TYPE_DVDRAM       DiscType = 0x12
	DISC_TYPE_DVDRWRO      DiscType = 0x13
	DISC_TYPE_DVDRWSR      DiscType = 0x14
	DISC_TYPE_DVDR_DL      DiscType = 0x15
	DISC_TYPE_DVDPLUSR_DL  DiscType = 0x16
	DISC_TYPE_DVDPLUSRW    DiscType = 0x1A
	DISC_TYPE_DVDPLUSR     DiscType = 0x1B
	DISC_TYPE_DVDPLUSR_DLD DiscType = 0x2B
	DISC_TYPE_BDROM        DiscType = 0x40
	DISC_TYPE_BDR          DiscType = 0x41
	DISC_TYPE_BDRE         DiscType = 0x43
	DISC_TYPE_UNKNOWN      DiscType = 0x00
)

func (d DiscType) String() string {
	switch d {
	case DISC_TYPE_BDR:
		return "BD-R"
	case DISC_TYPE_BDRE:
		return "BD-RE"
	case DISC_TYPE_BDROM:
		return "BD-ROM"
	case DISC_TYPE_CDR:
		return "CD-R"
	case DISC_TYPE_CDROM:
		return "CD-ROM"
	case DISC_TYPE_CDRW:
		return "CD-RW"
	//case DISC_TYPE_DDCDR:
	//	return "DDCD-R"
	//case DISC_TYPE_DDCDROM:
	//	return "DDCD-ROM"
	//case DISC_TYPE_DDCDRW:
	//	return "DDCD-RW"
	case DISC_TYPE_DVDPLUSR:
		return "DVD+R"
	case DISC_TYPE_DVDPLUSR_DL:
		return "DVD+R_DL"
	case DISC_TYPE_DVDPLUSR_DLD:
		return "DVD+R_DL"
	case DISC_TYPE_DVDPLUSRW:
		return "DVD+RW"
	case DISC_TYPE_DVDR:
		return "DVD-R"
	case DISC_TYPE_DVDR_DL:
		return "DVD-R_DL"
	case DISC_TYPE_DVDRAM:
		return "DVD-RAM"
	case DISC_TYPE_DVDROM:
		return "DVD-ROM"
	case DISC_TYPE_DVDRWRO:
		return "DVD-RWRO"
	case DISC_TYPE_DVDRWSR:
		return "DVD-RWSR"
	//case DISC_TYPE_HDDVDR:
	//	return "HDDVD-R"
	//case DISC_TYPE_HDDVDROM:
	//	return "HDDVD-ROM"
	//case DISC_TYPE_HDDVDRW:
	//	return "HDDVD-RW"
	case DISC_TYPE_NO_MEDIA:
		return "NO_MEDIA"
	case DISC_TYPE_UNKNOWN:
		return "UNKNOWN"
	default:
		return "UNKNOWN"
	}
}
