package handler

import (
	"github.com/TreeHole-ccnu/TreeHole-backend/model"
	"github.com/TreeHole-ccnu/TreeHole-backend/pkg"
	"github.com/TreeHole-ccnu/TreeHole-backend/middleware"
	"github.com/gin-gonic/gin"
)

func SendVer (c *gin.Context){
	phone := c.Query("phone")

	code := model.Code()
	sendRes := model.SendMsg(phone, code)

	if sendRes == "failed" {
		SendServerError(c, errno.InternalServerError, nil, "failed")
	} else {
		if !model.SetRedis(phone, code) {
			SendServerError(c, errno.InternalServerError, nil, )
		}
	}
}

func UserLogin (c *gin.Context) {
	var userLoginInfo model.LoginInfo
	if err := c.BindJSON(&userLoginInfo); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	//i = 1时，该用户还未注册、i = 2时，该用户输入密码错误、i = 3时，该用户输入密码正确
	i, err := model.ComfirmUserPhone(userLoginInfo.Phone, userLoginInfo.Passoword)
	if err != nil {
		SendServerError(c, errno.InternalServerError, nil, err.Error())
		return
	}

	if i == 1 {
		SendUnauthorizedError(c, errno.ErrUserNotFound, nil, err.Error())
		return
	} else if i == 2 {
		SendUnauthorizedError(c, errno.ErrPasswordIncorrect, nil, err.Error())
		return
	} else if i == 3 {
		c.JSON(200, gin.H{
			"message": "Login successfully.",
			"token": middleware.ProduceToken(userLoginInfo.Phone)
		})
		return
	}
}