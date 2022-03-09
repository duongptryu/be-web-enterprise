package common

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type File struct {
	SQLModelCreate
	Url        string  `json:"url,omitempty" gorm:"column:url;"`
	Name       string  `json:"name,omitempty" gorm:"column:name"`
	NameOrigin string  `json:"name_origin,omitempty" gorm:"column:name_origin"`
	Size       float64 `json:"size,omitempty" gorm:"column:size"`
	Ext        string  `json:"ext,omitempty" gorm:"column:ext"`
}

func (File) TableName() string {
	return "files"
}

func (f *File) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value", value))
	}

	var file File
	if err := json.Unmarshal(bytes, &file); err != nil {
		return err
	}
	*f = file
	return nil
}

func (f *File) Value() (driver.Value, error) {
	if f == nil {
		return nil, nil
	}
	return json.Marshal(f)
}

type Files []File

func (fs *Files) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value", value))
	}

	var files []File
	if err := json.Unmarshal(bytes, &files); err != nil {
		return err
	}
	*fs = files
	return nil
}

func (f *Files) Value() (driver.Value, error) {
	if f == nil {
		return nil, nil
	}
	return json.Marshal(f)
}
