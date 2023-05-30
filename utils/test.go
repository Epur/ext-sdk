package utils

import (
	"encoding/json"
	"errors"
	"time"
)

func Err(e string) error {
	return errors.New(e)
}

func PString(e string) *string {
	return &e
}

func TimestampSecond() int64 {
	return time.Now().UnixMilli() / 1000
}

func IsEmpty(val string) bool {
	return val == ""
}

func ToJson(data interface{}) string {
	e, _ := json.Marshal(data)
	return string(e)
}

func ConvertToString(v interface{}) (str string) {
	if v == nil {
		return ""
	}
	var (
		bs  []byte
		err error
	)
	if bs, err = json.Marshal(v); err != nil {
		return ""
	}
	str = string(bs)
	return
}
