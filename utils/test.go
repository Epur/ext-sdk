package utils

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"math/rand"
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

func GetRandLimitInt(s int, e int) int {
	if s == e {
		return s
	}
	rand.Seed(time.Now().UnixNano())
	a := rand.Intn(e-s) + s
	return a
}

func EncodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))
	return hex.EncodeToString(m.Sum(nil))
}

func HmacSha256(key string, data string) []byte {
	mac := hmac.New(sha256.New, []byte(key))
	_, _ = mac.Write([]byte(data))

	return mac.Sum(nil)
}

func HmacSha256ToBase64(key string, data string) string {
	return Base64UrlEncode(HmacSha256(key, data))
}

func Base64UrlEncode(data []byte) string {
	return base64.URLEncoding.EncodeToString(data)
}

func Base64Encode(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}
