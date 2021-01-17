package handler

import (
	"github.com/TreeHole-ccnu/TreeHole-backend/model"
	"github.com/TreeHole-ccnu/TreeHole-backend/pkg"
	//"github.com/TreeHole-ccnu/TreeHole-backend/middleware"
	"github.com/gin-gonic/gin"
	"strconv"
	"log"
)

func UserResetting (c *gin.Context) {
	var data model.RegisterInfo

	// Token := c.Request.Header.Get("Token")
	// if token ,err := middleware.VerifyToken(Token); err != nil{
	// 	data.Phone = token.Claims.(jwt.MapClaims)["uid"]
	// 	SendBadRequest(c, errno.ErrBind, nil, err.Error())
	// 	return
	// } 
	phone,_ := c.Get("uid")
	data.Phone = phone.(string)

	if err := c.BindJSON(&data); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	if model.ConfirmUserVcd(data.Phone, data.Vcd) == 0 {
		SendUnauthorizedError(c, errno.ErrValidation, nil, "The vaildation code is not correct")
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

	// Token := c.Request.Header.Get("Token")
	// if data.phone,err := middleware.VerifyToken(Token); err != nil{
	// 	SendBadRequest(c, errno.ErrBind, nil, err.Error())
	// 	return
	// } 
	phone,_ := c.Get("uid")
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
	// Token := c.Request.Header.Get("Token")
	// if phone,_ := c.Get("uid"); err != nil{
	// 	SendBadRequest(c, errno.ErrBind, nil, err.Error())
	// 	return
	// } 
	phone,_ := c.Get("uid")
	data.Phone = phone.(string)

	// if err := c.BindJSON(&data); err != nil {
	// 	SendBadRequest(c, errno.ErrBind, nil, err.Error())
	// 	return
	// }
	
	if data,err = model.GetInfo(data.Phone); err != nil {
		SendServerError(c, errno.InternalServerError, nil, err.Error())
		return
	}

	SendResponse(c, errno.OK, data)
}

func UserImage (c *gin.Context) {
	var data model.User
	var	err	error
	// Token := c.Request.Header.Get("Token")
	// if phone,_ := c.Get("uid"); err != nil{
	// 	SendBadRequest(c, errno.ErrBind, nil, err.Error())
	// 	return
	// } 
	phone,_ := c.Get("uid")
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

	//model.Image_modify(data.User_id, data.Image_url)
	if err := model.Image_modify(data.Phone, data.ImageUrl); err != nil {
		log.Println(err)
		SendServerError(c, errno.InternalServerError, nil, err.Error())
		return
	}

	SendResponse(c, errno.OK, nil)
	return
}