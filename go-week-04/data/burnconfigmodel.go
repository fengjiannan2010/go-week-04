package data

import (
	"gorm.io/gorm"
)

type (
	BurnConfigModel interface {
		Insert(data BurnConfig) (BurnConfig, error)
		FindOne(id int64) (*BurnConfig, error)
		Update(data BurnConfig) error
		Delete(id int64) error
		DeleteAll() error
	}

	defaultBurnConfigModel struct {
		conn *gorm.DB
	}

	BurnConfig struct {
		BaseModel
		DiscLabel  string     `gorm:"column:disc_label;type:varchar(255)" json:"disc_label"`
		DiscPasswd string     `gorm:"column:disc_passwd;type:varchar(255)" json:"disc_passwd"`
		FsType     DiscFsType `gorm:"column:fs_type;type:int(11)" json:"fs_type"`
		RecordMode RecordMode `gorm:"column:record_mode;type:int(11)" json:"record_mode"`
		IsVerify   bool       `gorm:"column:verify;type:int(11)" json:"is_verify"`
	}
)

func NewBurnConfigModel(conn *gorm.DB) BurnConfigModel {
	return &defaultBurnConfigModel{
		conn: conn,
	}
}

func (b *BurnConfig) TableName() string {
	return "burn_config"
}

func (b *BurnConfig) Reset() {
	b = &BurnConfig{}
}

func (m *defaultBurnConfigModel) Insert(data BurnConfig) (BurnConfig, error) {
	err := m.conn.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&data).Error; err != nil {
			// 遇到错误时回滚事务
			return err
		}
		return nil
	})
	return data, err
}

func (m *defaultBurnConfigModel) FindOne(id int64) (*BurnConfig, error) {
	var resp BurnConfig
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

func (m *defaultBurnConfigModel) Update(data BurnConfig) error {
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

func (m *defaultBurnConfigModel) Delete(id int64) error {
	err := m.conn.Delete(&BurnConfig{}, id).Error
	return err
}

func (m *defaultBurnConfigModel) GetDiscBurnConfig() (*BurnConfig, error) {
	var resp BurnConfig
	err := m.conn.First(&resp).Error
	switch err {
	case nil:
		return &resp, nil
	case gorm.ErrRecordNotFound:
		return nil, gorm.ErrRecordNotFound
	default:
		return nil, err
	}
}

func (m *defaultBurnConfigModel)DeleteAll() error {
	return m.conn.Exec("delete from burn_config").Error
}
