package data

import (
	"errors"
	"gorm.io/gorm"
)

type (
	FilePartInfoModel interface {
		Insert(data FilePartInfo) (FilePartInfo, error)
		FindOne(id int64) (*FilePartInfo, error)
		Update(data FilePartInfo) error
		Delete(id int64) error
		FindAny(taskId int64) (*[]FilePartInfo, error)
		DeleteAll() error
	}

	defaultFilePartInfoModel struct {
		conn *gorm.DB
	}

	FilePartInfo struct {
		BaseModel
		TaskId       int64      `gorm:"column:task_id;type:int(11)" json:"task_id"`
		UUID         string     `gorm:"column:uuid;type:varchar(255)" json:"uuid"`
		FilePath     string     `gorm:"column:file_path;type:varchar(255)" json:"file_path"`
		FileNameHash string     `gorm:"column:file_name_hash;type:varchar(255)" json:"file_name_hash"`
		ParentHash   string     `gorm:"column:parent_hash;type:varchar(255)" json:"parent_hash"`
		FileType     FileType   `gorm:"column:file_type;type:int(11)" json:"file_type"`
		Status       PartStatus `gorm:"column:status;type:int(11)" json:"status"`
		Crc32        int32      `gorm:"column:crc32;type:int(11)" json:"crc32"`
		OffSetStart  int64      `gorm:"column:offset_start;type:bigint(20)" json:"offset_start"`
		OffSetEnd    int64      `gorm:"column:offset_end;type:bigint(20)" json:"offset_end"`
		IsExists     bool       `gorm:"column:exists;type:int(11)" json:"is_exists"`
	}
)

func NewFilePartInfoModel(conn *gorm.DB) FilePartInfoModel {
	return &defaultFilePartInfoModel{
		conn: conn,
	}
}

func (f *FilePartInfo) TableName() string {
	return "file_part_info"
}

func (m *defaultFilePartInfoModel) Insert(data FilePartInfo) (FilePartInfo, error) {
	var tmpPartInfo FilePartInfo
	err := m.conn.Transaction(func(tx *gorm.DB) error {
		err := tx.Where(&FilePartInfo{TaskId: data.TaskId, UUID: data.UUID}).First(&tmpPartInfo).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			if err = tx.Create(&data).Error; err != nil {
				// 遇到错误时回滚事务
				return err
			}
		} else {
			data.Id = tmpPartInfo.Id
			err := tx.Updates(&data).Error
			if err != nil {
				// 遇到错误时回滚事务
				return err
			}
		}
		return nil
	})
	return data, err
}

func (m *defaultFilePartInfoModel) FindOne(id int64) (*FilePartInfo, error) {
	var resp FilePartInfo
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

func (m *defaultFilePartInfoModel) Update(data FilePartInfo) error {
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

func (m *defaultFilePartInfoModel) Delete(id int64) error {
	err := m.conn.Delete(&FilePartInfo{}, id).Error
	return err
}

func (m defaultFilePartInfoModel) FindAny(taskId int64) (*[]FilePartInfo, error) {
	var resp []FilePartInfo
	err := m.conn.Where(FilePartInfo{TaskId: taskId}).Find(&resp).Error
	switch err {
	case nil:
		return &resp, nil
	case gorm.ErrRecordNotFound:
		return nil, gorm.ErrRecordNotFound
	default:
		return nil, err
	}
}
func (m defaultFilePartInfoModel)DeleteAll() error {
	return m.conn.Exec("delete from file_part_info").Error
}
