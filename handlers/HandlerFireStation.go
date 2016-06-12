package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	m "bfp/avi/api/models"
)

type FireStationHandler struct {
	db *gorm.DB
}

func NewFireStationHandler(db *gorm.DB) *FireStationHandler {
	return &FireStationHandler{db}
}

func (handler FireStationHandler) Index(c *gin.Context) {
	if IsTokenValid(c) {
		stations := []m.FireStation{}
		handler.db.Order("created_at desc").Find(&stations)
		c.JSON(http.StatusOK,stations)
	} else {
		respond(http.StatusUnauthorized,"Sorry, but your session has expired!",c,true)	
	}
	return
}

func (handler FireStationHandler) Create(c *gin.Context) {
	if IsTokenValid(c) {
		var newFireStation m.FireStation
		fs := m.FireStation{}
		c.Bind(&newFireStation)
		query := handler.db.Where("station_code = ? AND station_name = ?",newFireStation.StationCode,newFireStation.StationName).First(&fs)
		if query.RowsAffected > 0 {
			respond(http.StatusBadRequest,"Fire station already existing",c,true)
		} else {
			result := handler.db.Create(&newFireStation)
			if result.RowsAffected > 0 {
				c.JSON(http.StatusCreated,newFireStation)
			} else {
				respond(http.StatusBadRequest,result.Error.Error(),c,true)
			}
		}
	} else {
		respond(http.StatusUnauthorized,"Sorry, but your session has expired!",c,true)	
	}
	return
}

func (handler FireStationHandler) Update(c *gin.Context) {
	if IsTokenValid(c) {
		
	} else {
		respond(http.StatusUnauthorized,"Sorry, but your session has expired!",c,true)	
	}
	return
}