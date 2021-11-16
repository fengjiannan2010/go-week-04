package data

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"time"
)

var TimeFormat = "2006-01-02 15:04:05"

type (
	BaseModel struct {
		Id        int64        `gorm:"column:id;type:bigint(20);primary_key" json:"id"`
		CreatedAt sql.NullTime `gorm:"column:created_at;type:varchar(255)" json:"created_at"`
		UpdatedAt sql.NullTime `gorm:"column:updated_at;type:varchar(255)" json:"updated_at"`
	}

	DateTime struct {
		time.Time
	}
)

// TODO: 重写 MarshaJSON 方法，在此方法中实现自定义格式的转换；程序中解析到JSON
func (d DateTime) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%s\"", d.Format(TimeFormat))
	return []byte(formatted), nil
}

// TODO: JSON中解析到程序中
func (t *DateTime) UnmarshalJSON(data []byte) (err error) {
	now, _ := time.ParseInLocation(`"`+TimeFormat+`"`, string(data), time.Local)
	*t = DateTime{Time: now}
	return
}

//TODO: 写入数据库时会调用该方法将自定义时间类型转换并写入数据库
func (t DateTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return time.Now().Format(TimeFormat), nil
	}
	return t.Time.Format(TimeFormat), nil
}

func (t DateTime) String() string {
	return t.Time.Format(TimeFormat)
}

//TODO: 读取数据库时会调用该方法将时间数据转换成自定义时间类型
func (t *DateTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = DateTime{Time: value}
		return nil
	}

	strvalue, ok := v.(string)
	if ok {
		now, _ := time.ParseInLocation(`"`+TimeFormat+`"`, strvalue, time.Local)
		*t = DateTime{Time: now}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
