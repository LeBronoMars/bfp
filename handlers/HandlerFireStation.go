package handlers

import (
	"net/http"
	"strconv"
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
	var newFireStation m.FireStation
	fs := m.FireStation{}
	err := c.Bind(&newFireStation)

	if err == nil {
		query := handler.db.Where("station_code = ? AND station_name = ?",newFireStation.StationCode,newFireStation.StationName).First(&fs)
		if query.RowsAffected > 0 {
			respond(http.StatusBadRequest, "Fire station already existing", c, true)
		} else {
			result := handler.db.Create(&newFireStation)
			if result.RowsAffected > 0 {
				c.JSON(http.StatusCreated,newFireStation)
			} else {
				respond(http.StatusBadRequest, result.Error.Error(), c, true)
			}
		}
	} else {
		respond(http.StatusUnprocessableEntity, err.Error(), c, true)
	}
}

func (handler FireStationHandler) Update(c *gin.Context) {
station_id := c.Param("station_id")
	fs := m.FireStation{}
	
	if handler.db.Where("id = ? ",station_id).First(&fs).RowsAffected > 0 {
		otherStation := m.FireStation{}

		if handler.db.Where("(station_code = ? OR station_name = ?) AND id != ?",fs.StationCode,fs.StationName,station_id).First(&otherStation).RowsAffected < 1 {
			newLatitude,_ := strconv.ParseFloat(c.PostForm("latitude"), 64)
			newLongitude,_ := strconv.ParseFloat(c.PostForm("longitude"), 64)

			fs.Address = c.PostForm("address")
			fs.Latitude = newLatitude
			fs.Longitude = newLongitude
			//fs.ContactNo = c.PostForm("contact_no")
			fs.Email = c.PostForm("email")
			save := handler.db.Save(&fs)
			if save.RowsAffected > 0 {
				c.JSON(http.StatusOK,fs)
			} else {
				respond(http.StatusBadRequest,save.Error.Error(),c,true)
			}
		} else {
			respond(http.StatusBadRequest,"Sorry but station code or name was already used by other station",c,true)
		}
	} else {
		respond(http.StatusBadRequest,"Fire station unable to find!",c,true)
	}
	return
}