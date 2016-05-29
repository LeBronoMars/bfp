package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	m "bfp/avi/api/models"
)

type IncidentHandler struct {
	db *gorm.DB
}

func NewIncidentHandler(db *gorm.DB) *IncidentHandler {
	return &IncidentHandler{db}
}

//get all incident
func (handler IncidentHandler) Index(c *gin.Context) {
	if IsTokenValid(c) {
		incidents := []m.MemberIncident{}
		status := c.Query("status")
		handler.db.Table("qry_member_incident").Where("incident_status = ?",status).Order("incident_report_id desc").Find(&incidents)
		c.JSON(http.StatusOK, &incidents)
	} else {
		respond(http.StatusUnauthorized,"Sorry, but your session has expired!",c,true)	
	}
	return
}

func (handler IncidentHandler) Create(c *gin.Context) {
	if IsTokenValid(c) {
		var newIncident	m.Incident
		c.Bind(&newIncident)
		
		result := handler.db.Create(&newIncident)
		if result.RowsAffected > 0 {
			c.JSON(http.StatusCreated,newIncident)
		} else {
			respond(http.StatusBadRequest,result.Error.Error(),c,true)
		}
	} else {
		respond(http.StatusUnauthorized,"Sorry, but your session has expired!",c,true)	
	}
	return
}


