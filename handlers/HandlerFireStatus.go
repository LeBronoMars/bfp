package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	m "bfp/avi/api/models"
)

type FireStatusHandler struct {
	db *gorm.DB
}

func NewFireStatusHandler(db *gorm.DB) *FireStatusHandler {
	return &FireStatusHandler{db}
}

func (handler FireStatusHandler) Create(c *gin.Context) {
	if IsTokenValid(c) {
		var newFireStatus m.FireStatus
		c.Bind(&newFireStatus)
		result := handler.db.Create(&newFireStatus)
		if result.RowsAffected > 0 {
			status := m.QryIncidents{}
			handler.db.Where("fire_status_id = ?",newFireStatus.Id).First(&status)
			c.JSON(http.StatusCreated,status)
		} else {
			respond(http.StatusBadRequest,result.Error.Error(),c,true)
		}
	} else {
		respond(http.StatusUnauthorized,"Sorry, but your session has expired!",c,true)	
	}
	return
}

func (handler FireStatusHandler) Update(c *gin.Context) {
	if IsTokenValid(c) {
		fireStatus := m.FireStatus{}
		handler.db.Where("id = ?",c.Param("id")).First(&fireStatus)
		fireStatus.Status = c.PostForm("status")
		save := handler.db.Save(&fireStatus)
		if save.RowsAffected > 0 {
			c.JSON(http.StatusOK,&fireStatus)
		} else {
			respond(http.StatusBadRequest,save.Error.Error(),c,true)
		}
	} else {
		respond(http.StatusUnauthorized,"Sorry, but your session has expired!",c,true)	
	}
	return
}