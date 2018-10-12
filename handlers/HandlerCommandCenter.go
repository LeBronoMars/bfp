package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	m "bfp/avi/api/models"
)

type CommandCenterHandler struct {
	db *gorm.DB
}

func NewCommandCenterHandler(db *gorm.DB) *CommandCenterHandler {
	return &CommandCenterHandler{db}
}

//get all command center
func (handler CommandCenterHandler) Index(c *gin.Context) {
	commanCenters := []m.CommandCenter{}
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

	query.Find(&commanCenters)

	c.JSON(http.StatusOK, commanCenters)
}

func (handler CommandCenterHandler) Create(c *gin.Context) {
	var commandCenter m.CommandCenter
	err := c.Bind(&commandCenter)

	if err == nil {
		existingByName := m.CommandCenter{}
		existingByContactNo := m.CommandCenter{}
		existingByEmail := m.CommandCenter{}

		if handler.db.Where("name = ?", commandCenter.Name).First(&existingByName).RowsAffected > 0 {
			respond(http.StatusUnprocessableEntity, "Command center name already taken.", c, true)
		} else if handler.db.Where("contact_no = ?", commandCenter.ContactNo).First(&existingByContactNo).RowsAffected > 0 {
			respond(http.StatusUnprocessableEntity, "Command center contact no already taken.", c, true)
		} else if handler.db.Where("email = ?", commandCenter.Email).First(&existingByEmail).RowsAffected > 0 {
			respond(http.StatusUnprocessableEntity, "Command center email already taken.", c, true)
		} else {
			result := handler.db.Create(&commandCenter)

			if result.RowsAffected > 0 {
				c.JSON(http.StatusCreated, commandCenter)
			} else {
				respond(http.StatusBadRequest, result.Error.Error(), c, true)
			}
		}
	} else {
		respond(http.StatusBadRequest, err.Error(), c, true)
	}
}

func (handler CommandCenterHandler) Update(c *gin.Context) {
	id := c.Param("id")
	commandCenter := m.CommandCenter{}
	qry := handler.db.Where("id = ?", id).First(&commandCenter)

	if qry.RowsAffected > 0 {

		if (len(c.PostForm("name")) == 0 && len(c.PostForm("contact_no")) == 0 && len(c.PostForm("email")) == 0) {
			respond(http.StatusUnprocessableEntity, "Nothing to update, no parameters provided.", c, true)	
		} else {
			// check if command center name is available
			if (len(c.PostForm("name")) > 0) {
				existing := m.CommandCenter{}

				if handler.db.Where("name = ? and id != ?", c.PostForm("name"), commandCenter.Id).First(&existing).RowsAffected > 0 {
					respond(http.StatusUnprocessableEntity, "Command center name already taken.", c, true)
					return
				} else {
					// update command center name
					commandCenter.Name = c.PostForm("name")
				}
			}

			// check if command center contact no is available
			if (len(c.PostForm("contact_no")) > 0) {
				existing := m.CommandCenter{}

				if handler.db.Where("contact_no = ? and id != ?", c.PostForm("contact_no"), commandCenter.Id).First(&existing).RowsAffected > 0 {
					respond(http.StatusUnprocessableEntity, "Command center contact no already taken.", c, true)
					return
				} else {
					// update command center name
					commandCenter.ContactNo = c.PostForm("contact_no")
				}
			}

			// check if command center email is available
			if (len(c.PostForm("email")) > 0) {
				existing := m.CommandCenter{}

				if handler.db.Where("email = ? and id != ?", c.PostForm("email"), commandCenter.Id).First(&existing).RowsAffected > 0 {
					respond(http.StatusUnprocessableEntity, "Command center email already taken.", c, true)
					return
				} else {
					// update command center name
					commandCenter.Email = c.PostForm("email")
				}
			}

			result := handler.db.Save(&commandCenter)

			if result.RowsAffected > 0 {
				c.JSON(http.StatusOK, commandCenter)
			} else {
				respond(http.StatusBadRequest, result.Error.Error(), c, true)	
			}
		}
	} else {
		respond(http.StatusNotFound, "Record not found.", c, true)	
	}
}