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
func (handler IncidentHandler) Show(c *gin.Context) {
	if IsTokenValid(c) {
		incident_id := c.Param("incident_id")
		qryIncident := m.FetchIncidents{}
		incident := m.Incident{}
		query := handler.db.Where("id = ?",incident_id).First(&incident)
		if query.RowsAffected > 0 {
			statuses := []m.QryIncidents{}
			handler.db.Where("incident_id = ?",incident_id).Find(&statuses)
			qryIncident.Incident = incident
			qryIncident.Status = statuses
			c.JSON(http.StatusOK,qryIncident)
		} else {
			respond(http.StatusBadRequest,"Unable to find incident record!",c,true)
		}
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


