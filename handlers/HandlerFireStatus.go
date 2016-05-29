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
			c.JSON(http.StatusCreated,newFireStatus)
		} else {
			respond(http.StatusBadRequest,result.Error.Error(),c,true)
		}
	} else {
		respond(http.StatusUnauthorized,"Sorry, but your session has expired!",c,true)	
	}
	return
}