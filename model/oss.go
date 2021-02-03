package model

import (
	"context"
	"errors"
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
	"github.com/spf13/viper"
	"io"
	"strconv"
	"strings"
	"time"
)

var (
	accessKey  string
	secretKey  string
	bucketName string
	domainName string
	upToken    string
	typeMap    map[string]bool
)

//不知道viper.getstring具体实现的原理，为初始化赋值与所需的基础属性
var initOSS = func() {
	accessKey = viper.GetString("oss.accessKey")
	secretKey = viper.GetString("oss.secretKey")
	bucketName = viper.GetString("oss.bucketName")
	domainName = viper.GetString("oss.domainName")
	//未知，推测是对文件格式的一种支持策略，表示对于这些格式的文件均支持收纳
	typeMap = map[string]bool{"jpg": true, "png": true, "bmp": true, "jpeg": true, "gif": true, "svg": true, "pdf": true, "ppt": true, "doc": true, "docx": true, "txt": true}
}

//根据传入文件名判断并获取文件格式
func getType(filename string) (string, error) {
	//strings.LastIndex作用为获取符号前字符串总数，例如在此处作用为从filename（“例如: xxxx.jpg”）获取“.”前字符串总数(“对应前面的例子为4”)
	i := strings.LastIndex(filename, ".")
	//对filename进行切割，得到文件格式，对应前例为jpg
	fileType := filename[i+1:]
	//判断是否为可支持格式的文件
	if !typeMap[strings.ToLower(fileType)] {
		return "", errors.New("the file type is not allowed")
	}
	return fileType, nil
}

//获得上传客户端的凭证
func getToken() {
	//设置最长有效时长
	var maxInt uint64 = 1 << 32
	initOSS()
	//经过initOSS后bucketName已有实质内容
	putPolicy := storage.PutPolicy{
		Scope:   bucketName,
		Expires: maxInt, //expires设置有效时长，不设置默认1小时
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken = putPolicy.UploadToken(mac)
}

//获取上传文件的路径
func getObjectName(filename string, id uint32) (string, error) {
	fileType, err := getType(filename)
	if err != nil {
		return "", err
	}
	//unix代表时间戳，time.Now().Unix()代表获取当前时间的时间戳，在不同平台具有良好的兼容性
	timeEpochNow := time.Now().Unix()
	objectName := strconv.FormatUint(uint64(id), 10) + "-" + strconv.FormatInt(timeEpochNow, 10) + "." + fileType
	return objectName, nil
}

func Uploadfile(filename string, fileid uint32, r io.ReaderAt, dataLen int64) (string, error) {
	if upToken == "" {
		getToken()
	}

	objectName, err := getObjectName(filename, fileid)
	if err != nil {
		return "", err
	}

	// 下面是七牛云的oss所需信息，objectName对应key是文件上传路径
	cfg := storage.Config{Zone: &storage.ZoneHuanan, UseHTTPS: false, UseCdnDomains: true}
	formUploader := storage.NewResumeUploader(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.RputExtra{Params: map[string]string{"x:name": "STACK"}}
	err = formUploader.Put(context.Background(), &ret, upToken, objectName, r, dataLen, &putExtra)
	//err = formUploader.PutFile(context.Background(), &ret, upToken, objectName, "/home/bowser/Pictures/maogai/1.jpg", &putExtra)
	if err != nil {
		return "", err
	}
	url := domainName + objectName
	return url, nil
}

func GetDownloadUrl(filename string) string {
	initOSS()
	key := filename
	mac := qbox.NewMac(accessKey, secretKey)
	deadline := time.Now().Add(time.Second * 3600).Unix()
	privateAccessURL := storage.MakePrivateURL(mac, domainName, key, deadline)
	return privateAccessURL
}
