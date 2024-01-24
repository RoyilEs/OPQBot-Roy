package models

import (
	"gorm.io/gorm"
	"strings"
)

//type StringArray []string

//// Scan 实现Scanner接口便于从数据库读取数据
//func (s *StringArray) Scan(value interface{}) error {
//	bytes, ok := value.([]byte)
//	if !ok {
//		return fmt.Errorf("can not convert %v to []byte", value)
//	}
//	err := json.Unmarshal(bytes, &s)
//	if err != nil {
//		return err
//	}
//	return nil
//}
//
//// Value 实现Valuer接口以便向数据库写入数据
//func (s StringArray) Value() (driver.Value, error) {
//	jsonBytes, err := json.Marshal(s)
//	if err != nil {
//		return nil, err
//	}
//
//	return jsonBytes, nil
//}

type FriendTag struct {
	gorm.Model
	Name     string `gorm:"column:name;type:varchar(255);not null;uniqueIndex"`
	TagsData string `gorm:"column:tag_data"`
}

// ArrayToString 序列化[]string到CSV字符串
func ArrayToString(arr []string) (string, error) {
	return strings.Join(arr, ","), nil
}

// StringToArray 反序列化CSV字符串到[]string
func StringToArray(str string) ([]string, error) {
	return strings.Split(str, ","), nil
}
