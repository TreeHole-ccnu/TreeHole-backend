package model

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/garyburd/redigo/redis"
	"log"

	"math/rand"
	"strconv"
	"time"
	"fmt"
)

// 向手机发送验证码
func SendMsg(tel string, code string) string {
	client, err := dysmsapi.NewClientWithAccessKey("cn-hangzhou", "<accesskeyId>", "<accessSecret>")
	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"
	request.PhoneNumbers = tel //手机号变量值
	request.SignName = "TreeHole树洞" //签名
	request.TemplateCode = "SMS_19586XXXX" //模板编码
	request.TemplateParam = "{\"code\":\"" + code + "\"}"
	response, err := client.SendSms(request)
	//fmt.Println(response.Code)

	if response.Code == "isv.BUSINESS_LIMIT_CONTROL" {
		return "frequency_limit"
	}
	if err != nil {
		fmt.Print(err.Error())
		return "failed"
	}
	return "success"
}

//随机生成六位验证码
func Code() string {
	rand.Seed(time.Now().UnixNano())
	code := rand.Intn(899999) + 100000
	res := strconv.Itoa(code)
	return res
}

func SetRedis(phone string, code string) bool {
	newRedis := RedisDb.Self.Get()
	_, err := newRedis.Do("SET", phone, code)
	if err != nil {
		log.Println("redis set error:", err)
		return false
	}
	_, err = newRedis.Do("expire", phone, 300)
	if err != nil {
		log.Println("set expire error: ", err)
		return false
	}
	return true
}

func GetRedis(phone string) string {
	newRedis := RedisDb.Self.Get()
	code, err := redis.String(newRedis.Do("GET", phone))
	if err != nil {
		log.Println("redis get error:", err)
	}
	return code
}