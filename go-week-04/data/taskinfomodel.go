package data

import (
	"gorm.io/gorm"
)

type (
	TaskInfoModel interface {
		Insert(data TaskInfo) (TaskInfo, error)
		FindOne(id int64) (*TaskInfo, error)
		Update(data TaskInfo) error
		Delete(id int64) error
		Exists(data TaskInfo) (bool, error)
		FindOneByStatus(data TaskInfo) (*TaskInfo, error)
		Count(data TaskInfo) (int64, error)
		UpdateAny(query TaskInfo, date TaskInfo) error
		DeleteAll() error
	}

	defaultTaskInfoModel struct {
		conn *gorm.DB
	}

	TaskInfo struct {
		BaseModel
		Expired        ExpiredType `gorm:"column:expired;type:int(11)" json:"expired"`
		UUID           string      `gorm:"column:uuid;type:varchar(255)" json:"uuid"`
		FileType       FileType    `gorm:"column:file_type;type:int(11)" json:"file_type"`
		StorageType    StorageType `gorm:"column:storage_type;type:int(11)" json:"storage_type"`
		FilePath       string      `gorm:"column:file_path;type:varchar(255)" json:"file_path"`
		ObjectName     string      `gorm:"column:object_name;type:varchar(255)" json:"object_name"`
		BucketName     string      `gorm:"column:bucket_name;type:varchar(255)" json:"bucket_name"`
		DiscPath       string      `gorm:"column:disc_path;type:varchar(255)" json:"disc_path"`
		CheckCode      string      `gorm:"comment:disc_path;type:varchar(255)" json:"check_code"`
		OffSetStart    int64       `gorm:"column:off_set_start;type:int(11)" json:"off_set_start"`
		OffSetEnd      int64       `gorm:"column:off_set_end;type:int(11)" json:"off_set_end"`
		Retries        int         `gorm:"column:retries;type:int(11)" json:"retries"`
		FileSize       int64       `gorm:"column:file_size;type:int(11)" json:"file_size"`
		DiscMode       DiscMode    `gorm:"column:disc_mode;type:int(11)" json:"disc_mode"`
		ParentId       int64       `gorm:"column:parent_id;type:int(11)" json:"parent_id"`
		Status         TaskStatus  `gorm:"column:status;type:int(11)" json:"status"`
		BurnProgress   int32       `gorm:"column:burn_progress;type:int(11)" json:"burn_progress"`
		VerifyProgress int32       `gorm:"column:verify_progress;type:int(11)" json:"verify_progress"`
		FileCrc32      uint32      `gorm:"column:file_crc32;type:int(11)" json:"file_crc32"`
		ErrorCode      int32         `gorm:"column:error_code;type:int(11)" json:"error_code"`
		ErrorMessage   string      `gorm:"column:error_message;type:varchar(255)" json:"error_message"`
		Error          string      `gorm:"column:error;type:varchar(255)" json:"error"`
		SubTasks       []TaskInfo  `gorm:"comment:sub_tasks;foreignkey:ParentId" json:"sub_tasks"`
	}
)

func NewTaskInfoModel(conn *gorm.DB) TaskInfoModel {
	return &defaultTaskInfoModel{
		conn: conn,
	}
}

func (t *TaskInfo) TableName() string {
	return "task_info"
}

func (t *TaskInfo) AddSubTask(task TaskInfo) {
	t.SubTasks = append(t.SubTasks, task)
}

func (t *TaskInfo) SyncTaskStatus() {
	status := t.Status
	if t.Status == Burning && len(t.SubTasks) > 0 {
		subStatus := t.Status
		for _, item := range t.SubTasks {
			if item.Status == BurningSuccess {
				subStatus = BurningSuccess
			} else {
				subStatus = t.Status
				break
			}
		}
		if subStatus == BurningSuccess {
			t.Status = subStatus
		} else {
			t.Status = status
		}
	}
}

// SyncDiscCapacity TODO 同步光盘容量
func (t *TaskInfo) SyncDiscCapacity() int64 {
	totalSize := int64(0)
	totalSize += t.OccupiedSize(t.FileSize)
	for _, item := range t.SubTasks {
		totalSize += t.OccupiedSize(item.FileSize)
	}
	return totalSize
}

// OccupiedSize TODO 计算文件光盘占用空间
func (t *TaskInfo) OccupiedSize(size int64) int64 {
	blocksize := int64(1024 * 128)
	if size < blocksize {
		size = blocksize
	} else {
		if size%blocksize != 0 {
			size = ((size / blocksize) + 1) * blocksize
		}
	}
	return size
}

func (m *defaultTaskInfoModel) Insert(data TaskInfo) (TaskInfo, error) {
	err := m.conn.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&data).Error; err != nil {
			// 遇到错误时回滚事务
			return err
		}
		return nil
	})
	return data, err
}

func (m *defaultTaskInfoModel) FindOne(id int64) (*TaskInfo, error) {
	var resp TaskInfo
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

func (m *defaultTaskInfoModel) Update(data TaskInfo) error {
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

func (m *defaultTaskInfoModel) Delete(id int64) error {
	err := m.conn.Delete(&TaskInfo{}, id).Error
	return err
}

func (m *defaultTaskInfoModel) Exists(data TaskInfo) (bool, error) {
	err := m.conn.Where(&TaskInfo{DiscPath: data.DiscPath}).First(&TaskInfo{}).Error
	switch err {
	case nil:
		return true, nil
	case gorm.ErrRecordNotFound:
		return false, nil
	default:
		return false, err
	}
}

func (m *defaultTaskInfoModel) FindOneByStatus(data TaskInfo) (*TaskInfo, error) {
	var resp TaskInfo
	err := m.conn.Where(&data).First(&resp).Error
	switch err {
	case nil:
		return &resp, nil
	case gorm.ErrRecordNotFound:
		return nil, gorm.ErrRecordNotFound
	default:
		return nil, err
	}
}

func (m *defaultTaskInfoModel) Count(data TaskInfo) (int64, error) {
	count := int64(0)
	err := m.conn.Where(&data).Count(&count).Error
	switch err {
	case nil:
		return count, nil
	case gorm.ErrRecordNotFound:
		return 0, gorm.ErrRecordNotFound
	default:
		return 0, err
	}
}

func (m *defaultTaskInfoModel) UpdateAny(query TaskInfo, date TaskInfo) error {
	err := m.conn.Transaction(func(tx *gorm.DB) error {
		err := tx.Where(&query).Updates(&date).Error
		if err != nil {
			// 遇到错误时回滚事务
			return err
		}
		return nil
	})
	return err
}

func (m *defaultTaskInfoModel) DeleteAll() error {
	return m.conn.Exec("delete from task_info").Error
}
