package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	m "bfp/avi/api/models"
)

type RegionalOfficeHandler struct {
	db *gorm.DB
}

func NewRegionalOfficeHandler(db *gorm.DB) *RegionalOfficeHandler {
	return &RegionalOfficeHandler{db}
}

//get all regional offices
func (handler RegionalOfficeHandler) Index(c *gin.Context) {
	offices := []m.RegionalOffice{}
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

	query.Find(&offices)

	c.JSON(http.StatusOK, offices)
}

func (handler RegionalOfficeHandler) Create(c *gin.Context) {
	var regionalOffice m.RegionalOffice
	err := c.Bind(&regionalOffice)

	if err == nil {
		existingOfficeByName := m.RegionalOffice{}

		if handler.db.Where("name = ?", regionalOffice.Name).First(&existingOfficeByName).RowsAffected > 0 {
			respond(http.StatusUnprocessableEntity, "Regional office name already taken.", c, true)
		} else {
			result := handler.db.Create(&regionalOffice)

			if result.RowsAffected > 0 {
				c.JSON(http.StatusCreated, regionalOffice)
			} else {
				respond(http.StatusBadRequest, result.Error.Error(), c, true)
			}
		}
	} else {
		respond(http.StatusBadRequest, err.Error(), c, true)
	}
}

func (handler RegionalOfficeHandler) Update(c *gin.Context) {
	id := c.Param("id")
	regionalOffice := m.RegionalOffice{}
	qry := handler.db.Where("id = ?", id).First(&regionalOffice)

	if qry.RowsAffected > 0 {

		if (len(c.PostForm("name")) == 0) {
			respond(http.StatusUnprocessableEntity, "Nothing to update, no parameter provided.", c, true)	
		} else {
			// check if regional office name is available
			if (len(c.PostForm("name")) > 0) {
				existingOfficeByName := m.RegionalOffice{}
				
				if handler.db.Where("name = ? and id != ?", c.PostForm("name"), regionalOffice.Id).First(&existingOfficeByName).RowsAffected > 0 {
					respond(http.StatusUnprocessableEntity, "Regional office name already taken.", c, true)
					return
				} else {
					// update regional office  name
					regionalOffice.Name = c.PostForm("name")
				}
			}

			result := handler.db.Save(&regionalOffice)

			if result.RowsAffected > 0 {
				c.JSON(http.StatusOK, regionalOffice)
			} else {
				respond(http.StatusBadRequest, result.Error.Error(), c, true)	
			}
		}
	} else {
		respond(http.StatusNotFound, "Record not found.", c, true)	
	}
}