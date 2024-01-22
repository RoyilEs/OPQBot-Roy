package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
)

type StringArray []string

// Scan 实现Scanner接口便于从数据库读取数据
func (s *StringArray) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("can not convert %v to []byte", value)
	}
	err := json.Unmarshal(bytes, &s)
	if err != nil {
		return err
	}
	return nil
}

// Value 实现Valuer接口以便向数据库写入数据
func (s StringArray) Value() (driver.Value, error) {
	jsonBytes, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}

	return jsonBytes, nil
}

type FriendTag struct {
	gorm.Model
	Name     string      `gorm:"column:name;type:varchar(255);not null;uniqueIndex"`
	TagsData StringArray `gorm:"column:tag_data;type:json"`
}
