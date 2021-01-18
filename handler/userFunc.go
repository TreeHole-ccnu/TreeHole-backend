package handler

import (
	"github.com/TreeHole-ccnu/TreeHole-backend/middleware"
	"github.com/TreeHole-ccnu/TreeHole-backend/model"
	"github.com/TreeHole-ccnu/TreeHole-backend/pkg"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SendVer (c *gin.Context){
	phone := c.Query("phone")

	code := model.Code()
	sendRes := model.SendMsg(phone, code)

	if sendRes == "failed" || sendRes == "frequency_limit" {
		SendServerError(c, errno.InternalServerError, nil, "send message failed")
		return
	} else {
		if model.SetRedis(phone, code) {
			SendServerError(c, errno.InternalServerError, nil, "set redis failed")
			return
		}
	}

	SendResponse(c, errno.OK, nil)
	return
}

func UserLogin (c *gin.Context) {
	var userLoginInfo model.LoginInfo
	if err := c.BindJSON(&userLoginInfo); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	//i = 1时，该用户还未注册、i = 2时，该用户输入密码错误、i = 3时，该用户输入密码正确
	i, err := model.ConfirmUserPhone(userLoginInfo.Phone, userLoginInfo.Password)
	if err != nil {
		SendServerError(c, errno.InternalServerError, nil, err.Error())
		return
	}
	if i == 1 {
		SendUnauthorizedError(c, errno.ErrUserNotFound, nil, "User didn't exist ！")
		return
	} else if i == 2 {
		SendUnauthorizedError(c, errno.ErrPasswordIncorrect, nil, "The password is incorrect ！")
		return
	} else if i == 3 {
		c.JSON(http.StatusOK, gin.H{
			"message": "Login successfully.",
			"token": middleware.ProduceToken(userLoginInfo.Phone),
		})
		return
	}
}

func UserRegister (c *gin.Context) {
	var data model.RegisterInfo
	if err := c.BindJSON(&data); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	if model.ConfirmUserVcd(data.Phone, data.Vcd) == 0 {
		SendUnauthorizedError(c, errno.ErrValidation, nil, "The vaildation code is not correct")
		return
	}

	if err := model.CreateUserRegisterInfo(data.Phone, data.Password); err != nil {
		SendServerError(c, errno.InternalServerError, nil, err.Error())
		return
	}

	SendResponse(c, errno.OK, nil)
}

