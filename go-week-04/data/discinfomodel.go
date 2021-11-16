package data

import (
	"gorm.io/gorm"
)

type (
	DiscInfoModel interface {
		Insert(data DiscInfo) (DiscInfo, error)
		FindOne(id int64) (*DiscInfo, error)
		Update(data DiscInfo) error
		Delete(id int64) error
	}

	defaultDiscInfoModel struct {
		conn *gorm.DB
	}

	DiscInfo struct {
		BaseModel
		Existed           int32      `gorm:"column:existed;type:int(11)" json:"existed"`
		IsBlank           bool       `gorm:"column:is_blank;type:int(11)" json:"is_blank"`
		IsCompleted       bool       `gorm:"column:is_completed;type:int(11)" json:"is_completed"`
		SerialNo          string     `gorm:"column:serial_no;type:varchar(255)" json:"serial_no"`
		Mid               string     `gorm:"column:m_id;type:varchar(255)" json:"m_id"`
		DiscType          DiscType   `gorm:"column:disc_type;type:int(11)" json:"disc_type"`
		TotalSize         int64      `gorm:"column:total_size;type:bigint(20)" json:"total_size"`
		FreeSize          int64      `gorm:"column:free_size;type:bigint(20)" json:"free_size"`
		UsedSize          int64      `gorm:"column:used_size;type:bigint(20)" json:"used_size"`
		TrackNum          int32      `gorm:"column:track_num;type:int(11)" json:"track_num"`
		UserDefinedId     string     `gorm:"column:user_defined_id;type:varchar(255)" json:"user_defined_id"`
		FsType            DiscFsType `gorm:"column:fs_type;type:int(11)" json:"fs_type"`
		DiscLabel         string     `gorm:"column:disc_label;type:varchar(255)" json:"disc_label"`
		DiscPasswd        string     `gorm:"column:disc_passwd;type:varchar(255)" json:"disc_passwd"`
		MediaStatus       int32      `gorm:"column:media_status;type:int(11)" json:"media_status"`
		WriteSpeed        []int32    `gorm:"-" json:"write_speed"`
		TrackStatusList   []int32    `gorm:"-" json:"track_status_list"`
		TrackSizeList     []int32    `gorm:"-" json:"track_size_list"`
		TrackUsedSizeList []int32    `gorm:"-" json:"track_used_size_list"`
		TrackNwaList      []int32    `gorm:"-" json:"track_nwa_list"`
	}
)

func NewDiscInfoModel(conn *gorm.DB) DiscInfoModel {
	return &defaultDiscInfoModel{
		conn: conn,
	}
}
func (d *DiscInfo) TableName() string {
	return "disc_info"
}

func (d *DiscInfo) IsExisted() bool {
	return d.Existed == 1
}

func (d *DiscInfo) SyncCapacity(size int64) {
	d.UsedSize += size
	d.FreeSize -= size
}

func (d *DiscInfo) ConvertSize() {
	d.TotalSize = d.TotalSize * BlockSize
	d.UsedSize = d.UsedSize * BlockSize
	d.FreeSize = d.FreeSize * BlockSize
}

func (m *defaultDiscInfoModel) Insert(data DiscInfo) (DiscInfo, error) {
	err := m.conn.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&data).Error; err != nil {
			// 遇到错误时回滚事务
			return err
		}
		return nil
	})
	return data, err
}

func (m *defaultDiscInfoModel) FindOne(id int64) (*DiscInfo, error) {
	var resp DiscInfo
	err := m.conn.First(&resp, id).Error
	switch err {
	case nil:
		return &resp, nil
	case gorm.ErrRecordNotFound:
		return nil, gorm.ErrRecordNotFound
	default:
		return nil, err
	}
}

func (m *defaultDiscInfoModel) Update(data DiscInfo) error {
	err := m.conn.Transaction(func(tx *gorm.DB) error {
		err := tx.Updates(&data).Error
		if err != nil {
			// 遇到错误时回滚事务
			return err
		}
		return nil
	})
	return err
}

func (m *defaultDiscInfoModel) Delete(id int64) error {
	err := m.conn.Delete(&DiscInfo{}, id).Error
	return err
}

func (m *defaultDiscInfoModel) DeleteAll() error {
	return m.conn.Exec("delete from disc_info").Error
}
