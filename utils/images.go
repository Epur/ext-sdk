package utils

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

// 获取网络图片转base64
func GetUrlImgBase64(path string) (baseImg string, err error) {
	//获取网络图片
	client := &http.Client{
		Timeout: time.Second * 5, //超时时间
	}
	var bodyImg io.Reader
	request, err := http.NewRequest("GET", path, bodyImg)
	if err != nil {
		err = errors.New("获取网络图片失败")
		return
	}
	respImg, _ := client.Do(request)
	defer respImg.Body.Close()
	imgByte, _ := ioutil.ReadAll(respImg.Body)

	// 判断文件类型，生成一个前缀，拼接base64后可以直接粘贴到浏览器打开，不需要可以不用下面代码
	//取图片类型
	mimeType := http.DetectContentType(imgByte)
	switch mimeType {
	case "image/jpeg":
		baseImg = "data:image/jpeg;base64," + base64.StdEncoding.EncodeToString(imgByte)
	case "image/png":
		baseImg = "data:image/png;base64," + base64.StdEncoding.EncodeToString(imgByte)
	}

	return
}

// 本地图片
func GetLocalImgBase64(path string) (baseImg string, err error) {

	//获取本地文件
	file, err := os.Open(path)
	if err != nil {
		err = errors.New("获取本地图片失败")
		return
	}
	fi, _ := file.Stat()
	fmt.Println("file size is:", fi.Size())

	defer file.Close()
	imgByte, _ := ioutil.ReadAll(file)

	// 判断文件类型，生成一个前缀，拼接base64后可以直接粘贴到浏览器打开，不需要可以不用下面代码
	//取图片类型
	mimeType := http.DetectContentType(imgByte)
	switch mimeType {
	case "image/jpeg":
		baseImg = "data:image/jpeg;base64," + base64.StdEncoding.EncodeToString(imgByte)
	case "image/png":
		baseImg = "data:image/png;base64," + base64.StdEncoding.EncodeToString(imgByte)
	}
	return
}
