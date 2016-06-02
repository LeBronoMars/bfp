package handlers

import (
	"net/http"
	"strconv"

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
			handler.db.Where("incident_id = ?",incident_id).Order("fire_status_id desc").Find(&statuses)
			qryIncident.Incident = incident
			qryIncident.Status = statuses
			c.JSON(http.StatusOK,qryIncident)
		} else {
			respond(http.StatusBadRequest,"Unable to find incident record!",c,true)
		}
	} else {
		respond(http.StatusForbidden,"Sorry, but your session has expired!",c,true)	
	}
	return
}

func (handler IncidentHandler) Index(c *gin.Context) {
	if (IsTokenValid(c)) {
		incidents := []m.Incident{}
		handler.db.Order("created_at desc").Find(&incidents)
		var qryIncidents = make([]m.FetchIncidents,len(incidents))
		for i,incident := range incidents {
			statuses := []m.QryIncidents{}
			handler.db.Where("incident_id = ?",incident.Id).Order("fire_status_id desc").Find(&statuses)
			qryIncidents[i].Incident = incident
			qryIncidents[i].Status = statuses
		}
		c.JSON(http.StatusOK,qryIncidents)
	} else {
		respond(http.StatusForbidden,"Sorry, but your session has expired!",c,true)	
	}
}

func (handler IncidentHandler) Create(c *gin.Context) {
	if IsTokenValid(c) {
		var newIncident	m.Incident
		c.Bind(&newIncident)
		alarm_level := c.PostForm("alarm_level")
		reported_by,_ := strconv.Atoi(c.PostForm("reported_by"))

		result := handler.db.Create(&newIncident)
		if result.RowsAffected > 0 {
			incident_id := newIncident.Id
			//create the very first fire status of incident
			fireStatus := m.FireStatus{}
			fireStatus.IncidentId = incident_id
			fireStatus.Status = alarm_level
			fireStatus.ReportedBy = reported_by

			fireStatusResult := handler.db.Create(&fireStatus)

			if fireStatusResult.RowsAffected > 0 {
				qryIncident := m.FetchIncidents{}
				incident := m.Incident{}
				handler.db.Where("id = ?",incident_id).First(&incident)
				statuses := []m.QryIncidents{}
				handler.db.Where("incident_id = ?",incident_id).Order("fire_status_id desc").Find(&statuses)
				qryIncident.Incident = incident
				qryIncident.Status = statuses
				c.JSON(http.StatusCreated,qryIncident)
			} else {
				respond(http.StatusBadRequest,fireStatusResult.Error.Error(),c,true)
			}
		} else {
			respond(http.StatusBadRequest,result.Error.Error(),c,true)
		}
	} else {
		respond(http.StatusForbidden,"Sorry, but your session has expired!",c,true)	
	}
	return
}


