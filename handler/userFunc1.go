package handler

import (
	"github.com/TreeHole-ccnu/TreeHole-backend/model"
	"github.com/TreeHole-ccnu/TreeHole-backend/pkg"
	"github.com/TreeHole-ccnu/TreeHole-backend/middleware"
	"github.com/gin-gonic/gin"
)

func UserResetting (c *gin.Context) {
	var data model.RegisterInfo

	Token := c.Request.Header.Get("Token")
	if data.phone,err := middleware.verifyToken(Token); err != nil{
		SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	} 

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


func UserNormalInfo (c *gin.Context) {
	var data model.User

	Token := c.Request.Header.Get("Token")
	if data.phone,err := middleware.verifyToken(Token); err != nil{
		SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	} 

	if err := c.BindJSON(&data); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	if err := model.ResetPassword(data); err != nil {
		SendServerError(c, errno.InternalServerError, nil, err.Error())
		return
	}

	SendResponse(c, errno.OK, nil)
}