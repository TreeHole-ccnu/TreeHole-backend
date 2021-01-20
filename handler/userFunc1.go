package handler

import (
	"github.com/TreeHole-ccnu/TreeHole-backend/model"
	"github.com/TreeHole-ccnu/TreeHole-backend/pkg"
	"net/http"

	//"github.com/TreeHole-ccnu/TreeHole-backend/middleware"
	"github.com/gin-gonic/gin"
	"strconv"
	"log"
)

func UserResetting (c *gin.Context) {
	var data model.RegisterInfo

	if err := c.BindJSON(&data); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	if model.ConfirmUser(data.Phone) == 0 {
		SendUnauthorizedError(c, errno.ErrUserNotFound, nil, "The user didn't exist ! ")
		return
	}

	if model.ConfirmUserVcd(data.Phone, data.Vcd) == 0 {
		SendUnauthorizedError(c, errno.ErrValidation, nil, "The verification code is not correct")
		return
	}

	if err := model.ResetPassword(data.Phone, data.Password); err != nil {
		SendServerError(c, errno.InternalServerError, nil, err.Error())
		return
	}

	SendResponse(c, errno.OK, nil)
}


func InfoResetting (c *gin.Context) {
	var data model.User

	phone,_ := c.Get("phone")
	data.Phone = phone.(string)

	if err := c.BindJSON(&data); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}
	
	if err := model.ResetNormalInfo(data); err != nil {
		SendServerError(c, errno.InternalServerError, nil, err.Error())
		return
	}

	SendResponse(c, errno.OK, nil)
}

func InfoGetting (c *gin.Context) {
	var data model.User
	var	err	error

	phone,_ := c.Get("phone")
	data.Phone = phone.(string)

	
	if data,err = model.GetInfo(data.Phone); err != nil {
		SendServerError(c, errno.InternalServerError, nil, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message" : "Getting information successfully ! ",
		"name" : data.Name,
		"birth" : data.Birth,
		"sex" : data.Sex,
		"nation" : data.Nation,
		"native_place" : data.NativePlace,
		"identity_number" : data.IdentityNumber,
		"phone" : data.Phone,
		"email" : data.Email,
		"image_url" : data.ImageUrl,
		"level" : data.Level,
	})
	return
}

//传url进数据库
func UserImage (c *gin.Context) {
	var data model.User

	phone,_ := c.Get("phone")
	data.Phone = phone.(string)
	if err := c.BindJSON(&data); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	//model.Image_modify(data.User_id, data.Image_url)
	if err := model.Image_modify(data.Phone, data.ImageUrl); err != nil {
		log.Println(err)
		SendServerError(c, errno.InternalServerError, nil, err.Error())
		return
	}

	SendResponse(c, errno.OK, nil)
	return
}

//通用接口，传图片，返回url
func Image (c *gin.Context) {
	var data model.User
	var	err	error

	phone,_ := c.Get("phone")
	data.Phone = phone.(string)

	fileid, _ := strconv.Atoi(c.Param("fileid"))
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		SendBadRequest(c, errno.ErrBindFormFile, nil, err.Error())
		return
	}
	dataLen := header.Size

	data.ImageUrl, err = model.Uploadfile(header.Filename, uint32(fileid), file, dataLen)
  //log.Print(fileid)

  if err != nil {
		SendBadRequest(c, errno.ErrAddr, nil, err.Error())
	return
}

	SendResponse(c, errno.OK, data.ImageUrl)
	return
}