package handlers

import (
	"net/http"
	"strings"
	"time"

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
		respond(http.StatusBadRequest,"Sorry, but your session has expired!",c,true)	
	}
}

func (handler IncidentHandler) Create(c *gin.Context) {
	report_id := c.PostForm("report_id")
	reported_by := c.PostForm("reported_by")
	alarm_level := c.PostForm("alarm_level")
	latitude := c.PostForm("latitude")
	longitude := c.PostForm("longitude")
	address := c.PostForm("address")
	remarks := c.PostForm("remarks")

	if strings.TrimSpace(report_id) == "" {
		respond(http.StatusBadRequest,"Please specify the report id",c,true)
	} else if (strings.TrimSpace(reported_by) == "") {
		respond(http.StatusBadRequest,"Please specify the reportee of the incident report",c,true)
	} else if (strings.TrimSpace(alarm_level) == "") {
		respond(http.StatusBadRequest,"Please specify the incident's alarm level",c,true)
	} else if (strings.TrimSpace(address) == "") {
		respond(http.StatusBadRequest,"Please specify the exact address of the fire incident",c,true)
	} else {
		now := time.Now().UTC()
		if (strings.TrimSpace(remarks) == "") {
			remarks = "None"
		}
		result := handler.db.Exec("INSERT INTO incident VALUES (null,?,?,?,?,?,?,?,?,?)",reported_by,"active",alarm_level,latitude,longitude,address,remarks,now,report_id)
		if result.RowsAffected == 1 {
			incident := m.Incident{}
			handler.db.Table("incident").Where("report_id = ?",report_id).Last(&incident)
			c.JSON(http.StatusCreated,incident)
		} else {
			respond(http.StatusBadRequest,"Unable to create incident report, Please try again",c,true)
		}
	}
}


