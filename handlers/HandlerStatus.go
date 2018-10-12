package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	m "bfp/avi/api/models"
)

type StatusHandler struct {
	db *gorm.DB
}

func NewStatusHandler(db *gorm.DB) *StatusHandler {
	return &StatusHandler{db}
}

//get all statuses
func (handler StatusHandler) Index(c *gin.Context) {
	statuses := []m.Status{}
	var query = handler.db

	startParam,startParamExist := c.GetQuery("start")
	limitParam,limitParamExist := c.GetQuery("limit")

	//start param exist
	if startParamExist {
		start,_ := strconv.Atoi(startParam)
		if start != 0 {
			query = query.Offset(start)				
		} else {
			query = query.Offset(0)
		}
	} 

	//limit param exist
	if limitParamExist {
		limit,_ := strconv.Atoi(limitParam)
		query = query.Limit(limit)
	} else {
		query = query.Limit(10)
	}

	query.Find(&statuses)

	c.JSON(http.StatusOK, statuses)
}

func (handler StatusHandler) Create(c *gin.Context) {
	var status m.Status
	err := c.Bind(&status)

	if err == nil {
		existingStatusByName := m.Status{}
		existingStatusByColor := m.Status{}

		if handler.db.Where("status_name = ?", status.StatusName).First(&existingStatusByName).RowsAffected > 0 {
			respond(http.StatusUnprocessableEntity, "Status name already taken.", c, true)
		} else if handler.db.Where("status_color = ?", status.StatusColor).First(&existingStatusByColor).RowsAffected > 0 {
			respond(http.StatusUnprocessableEntity, "Status color already taken.", c, true)
		} else {
			result := handler.db.Create(&status)

			if result.RowsAffected > 0 {
				c.JSON(http.StatusCreated, status)
			} else {
				respond(http.StatusBadRequest, result.Error.Error(), c, true)
			}
		}
	} else {
		respond(http.StatusBadRequest, err.Error(), c, true)
	}
}

func (handler StatusHandler) Update(c *gin.Context) {
	id := c.Param("id")
	status := m.Status{}
	qry := handler.db.Where("id = ?", id).First(&status)

	if qry.RowsAffected > 0 {

		if (len(c.PostForm("status_name")) == 0 && len(c.PostForm("status_color")) == 0) {
			respond(http.StatusUnprocessableEntity, "Nothing to update, no parameters provided.", c, true)	
		} else {
			// check if status color is available
			if (len(c.PostForm("status_name")) > 0) {
				existingStatusByName := m.Status{}
				
				if handler.db.Where("status_name = ?", c.PostForm("status_name")).First(&existingStatusByName).RowsAffected > 0 {
					respond(http.StatusUnprocessableEntity, "Status name already taken.", c, true)
					return
				} else {
					// update status name
					status.StatusName = c.PostForm("status_name")
				}
			}

			// check if status color is available
			if (len(c.PostForm("status_color")) > 0) {
				existingStatusByColor := m.Status{}
				
				if handler.db.Where("status_color = ?", c.PostForm("status_color")).First(&existingStatusByColor).RowsAffected > 0 {
					respond(http.StatusUnprocessableEntity, "Status color already taken.", c, true)
					return
				} else {
					// update status name
					status.StatusColor = c.PostForm("status_color")
				}
			}

			result := handler.db.Save(&status)

			if result.RowsAffected > 0 {
				c.JSON(http.StatusOK, status)
			} else {
				respond(http.StatusBadRequest, result.Error.Error(), c, true)	
			}
		}
	} else {
		respond(http.StatusNotFound, "Status record not found.", c, true)	
	}
}