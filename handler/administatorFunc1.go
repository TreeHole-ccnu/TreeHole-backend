package handler

import (
	"github.com/TreeHole-ccnu/TreeHole-backend/model"
	errno "github.com/TreeHole-ccnu/TreeHole-backend/pkg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func VolunteerSearch(c *gin.Context) {
	var p model.PhoneInfo
	var name string
	var id int
	var err error

	if err := c.BindJSON(&p); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}
	phone := p.Phone

	if model.ConfirmVolunteer(phone) == 0 {
		SendBadRequest(c, errno.ErrUserNotFound, nil, "The volunteer didn't exist ! ")
		return
	}

	if name, id, err = model.SearchVolunteer(phone); err != nil {
		SendBadRequest(c, errno.InternalServerError, nil, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "successful ! ",
		"name":    name,
		"id":      id,
	})
	return
}

func GetDetailedInfo(c *gin.Context) {
	var v model.Volunteer
	var r []model.Resume
	var err error

	id, _ := strconv.Atoi(c.Query("id"))

	if v, err = model.GetVolunteerInfo(id); err != nil {
		SendServerError(c, errno.InternalServerError, nil, err.Error())
		return
	}

	if r, err = model.GetUserResume(v.UserId); err != nil {
		SendServerError(c, errno.InternalServerError, nil, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":           "successful ! ",
		"name":              v.Name,
		"reference":         v.Reference,
		"birth":             v.Birth,
		"political":         v.Political,
		"sex":               v.Sex,
		"nation":            v.Nation,
		"native_place":      v.NativePlace,
		"nationality":       v.Nationality,
		"identity_number":   v.IdentityNumber,
		"workphone":         v.WorkPhone,
		"phone":             v.Phone,
		"email":             v.Email,
		"job":               v.Job,
		"social_job":        v.SocialJob,
		"medical_history":   v.MedicalHistory,
		"treatment_history": v.TreatmentHistory,
		"medicine":          v.Medicine,
		"resume":            r,
		"reason":            v.Reason,
		"front_url":         v.FrontUrl,
		"contrary_url":      v.ContraryUrl,
	})
	return
}
