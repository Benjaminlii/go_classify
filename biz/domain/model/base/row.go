package base

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type Row struct {
	Row map[string]interface{} `gorm:"column:row"`
}

func (row *Row) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}
	err := json.Unmarshal(bytes, row)
	return err
}

func (row *Row) Value() (driver.Value, error) {
	if len(row.Row) == 0 {
		return nil, nil
	}
	b, err := json.Marshal(*row)
	if err != nil {
		return nil, err
	}
	return string(b), nil
}
