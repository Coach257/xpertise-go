package utils

import (
	"database/sql/driver"

	jsoniter "github.com/json-iterator/go"
)

// customized type for string array store & read
type StrList []string

func (l StrList) Value() (driver.Value, error) {
	data, err := jsoniter.Marshal(l)
	return string(data), err
}

func (l *StrList) Scan(data interface{}) error {
	return jsoniter.Unmarshal(data.([]byte), l)
}

// customized type for uint array store & read
type UintList []uint

func (l UintList) Value() (driver.Value, error) {
	data, err := jsoniter.Marshal(l)
	return string(data), err
}

func (l *UintList) Scan(data interface{}) error {
	err := jsoniter.Unmarshal(data.([]byte), l)
	for i := 0; i < len(*l); i++ {
		print((*l)[i])
	}
	return err
}
