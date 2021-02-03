package handler

import (
	"github.com/TreeHole-ccnu/TreeHole-backend/model"
	"github.com/TreeHole-ccnu/TreeHole-backend/pkg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func LevelChanging(c *gin.Context) {
	idTmp := c.Query("id")
	id, _ := strconv.Atoi(idTmp)
	if !model.CheckUser(id) {
		SendBadRequest(c, errno.ErrUserNotFound, nil, "the user didn't exist ! ")
		return
	}

	if err := model.ChangeLevel(id); err != nil {
		SendServerError(c, errno.InternalServerError, nil, err.Error())
		return
	}

	SendResponse(c, errno.OK, nil)
}

func StatusChanging(c *gin.Context) {
	var id int

	idTmp := c.Query("id")
	id, err := strconv.Atoi(idTmp)
	if err != nil {
		SendBadRequest(c, errno.ErrQuery, nil, "The id is wrong.")
		return
	}
	if err := model.ChangeStatus(id); err != nil {
		SendServerError(c, errno.InternalServerError, nil, err.Error())
		return
	}

	SendResponse(c, errno.OK, nil)
}

func Verification(c *gin.Context) {
	var info []model.CheckingInfo
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		SendBadRequest(c, errno.ErrQuery, nil, "The page is wrong.")
		return
	}
	limit, _ := strconv.Atoi(c.Query("limit"))
	if err != nil {
		SendBadRequest(c, errno.ErrQuery, nil, "The limit is wrong.")
		return
	}

	if info, err = model.VerificationInfo(page, limit); err != nil {
		SendServerError(c, errno.InternalServerError, nil, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":      "success ! ",
		"informations": info,
	})

	return
}
