package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Epur/ext-sdk/logger"
	"net/url"
	"sort"
	"strings"
)

const (
	NULL = ""
)

type BodyMap map[string]interface{}

// 设置参数
func (bm BodyMap) Set(key string, value interface{}) BodyMap {
	bm[key] = value
	return bm
}

func (bm BodyMap) SetBodyMap(key string, value func(b BodyMap)) BodyMap {
	_bm := make(BodyMap)
	value(_bm)
	bm[key] = _bm
	return bm
}

//// 设置 FormFile
//func (bm BodyMap) SetFormFile(key string, file *util.File) BodyMap {
//	bm[key] = file
//	return bm
//}

// 获取参数，同 GetString()
func (bm BodyMap) Get(key string) string {
	return bm.GetString(key)
}

// 获取参数转换string
func (bm BodyMap) GetString(key string) string {
	if bm == nil {
		return NULL
	}
	value, ok := bm[key]
	if !ok {
		return NULL
	}
	v, ok := value.(string)
	if !ok {
		return convertToString(value)
	}
	return v
}

// 获取原始参数
func (bm BodyMap) GetInterface(key string) interface{} {
	if bm == nil {
		return nil
	}
	return bm[key]
}

// 删除参数
func (bm BodyMap) Remove(key string) {
	delete(bm, key)
}

// 置空BodyMap
func (bm BodyMap) Reset() {
	for k := range bm {
		delete(bm, k)
	}
}

func (bm BodyMap) JsonBody() (jb string) {
	bs, err := json.Marshal(bm)
	if err != nil {
		return ""
	}
	jb = string(bs)
	return jb
}

// Unmarshal to struct or slice point
func (bm BodyMap) Unmarshal(ptr interface{}) (err error) {
	bs, err := json.Marshal(bm)
	if err != nil {
		return err
	}
	return json.Unmarshal(bs, ptr)
}

func (bm BodyMap) EncodeURLParams() string {
	if bm == nil {
		return NULL
	}
	var (
		buf  strings.Builder
		keys []string
	)
	for k := range bm {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		if v := bm.GetString(k); v != NULL {
			buf.WriteString(url.QueryEscape(k))
			buf.WriteByte('=')
			buf.WriteString(url.QueryEscape(v))
			buf.WriteByte('&')
		}
	}
	if buf.Len() <= 0 {
		return NULL
	}
	return buf.String()[:buf.Len()-1]
}

func (bm BodyMap) CheckEmptyError(keys ...string) error {
	var emptyKeys []string
	for _, k := range keys {
		if v := bm.GetString(k); v == NULL {
			emptyKeys = append(emptyKeys, k)
		}
	}
	if len(emptyKeys) > 0 {
		logger.SdkLogger.Errorf("[%w], %v", errors.New("missing required parameter"), strings.Join(emptyKeys, ", "))
		return fmt.Errorf("[%w], %v", errors.New("missing required parameter"), strings.Join(emptyKeys, ", "))
	}
	return nil
}

func convertToString(v interface{}) (str string) {
	if v == nil {
		return NULL
	}
	var (
		bs  []byte
		err error
	)
	if bs, err = json.Marshal(v); err != nil {
		return NULL
	}
	str = string(bs)
	return
}
