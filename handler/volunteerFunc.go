package handler

import (
	"github.com/TreeHole-ccnu/TreeHole-backend/model"
	errno "github.com/TreeHole-ccnu/TreeHole-backend/pkg"
	"github.com/gin-gonic/gin"
	"net/http"
)

func VolunteerInfo (c *gin.Context) {
	phone := c.GetString("phone")

	var volunteerInfo model.VolunteerInfo
	var user model.User
	var err error

	if err := c.BindJSON(&volunteerInfo); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	if user, err = model.GetInfo(phone); err != nil {
		SendServerError(c, errno.InternalServerError, nil, err.Error())
		return
	}

	volunteer := model.Volunteer{
		UserId:           user.Id,
		IsCheck:          2,
		Reference:        volunteerInfo.Reference,
		Name:             user.Name,
		Political:        volunteerInfo.Political,
		Sex:              user.Sex,
		Nation:           user.Nation,
		NativePlace:      user.NativePlace,
		Education:        volunteerInfo.Education,
		Nationality:      volunteerInfo.Nationality,
		IdentityNumber:   user.IdentityNumber,
		WorkPhone:        volunteerInfo.Workphone,
		Phone:            user.Phone,
		Email:            user.Email,
		Job:              volunteerInfo.Job,
		SocialJob:        volunteerInfo.SocialJob,
		MedicalHistory:   volunteerInfo.MedicalHistory,
		TreatmentHistory: volunteerInfo.TreatmentHistory,
		Medicine:         volunteerInfo.Medicine,
		Reason:           volunteerInfo.Reason,
		FrontUrl:         volunteerInfo.FrontUrl,
		ContraryUrl:      volunteerInfo.ContraryUrl,
	}

	if err := model.CreateVolunteer(volunteer); err != nil {
		SendServerError(c, errno.InternalServerError, nil, err.Error())
		return
	}

	SendResponse(c, errno.OK, nil)
	return
}

func VolunteerCheck(c *gin.Context) {
	phone := c.GetString("phone")
	var i int
	var err error

	if model.ConfirmVolunteer(phone) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"check_id" : 1,
			"message" : "success ! ",
		})
	}

	if i, err = model.GetCheckId(phone); err != nil {
		//log.Println(err)
		SendServerError(c, errno.InternalServerError, nil, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"check_id" : i,
		"message" : "success ! ",
	})
	return
}