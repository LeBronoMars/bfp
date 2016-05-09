package models

import (
	"time"
)

type MemberIncident struct {
	UserId int `json:"user_id"`
	ReportersName string `json:"reporters_name"`
	UserRole string `json:"user_role"`
	AccountLevel string `json:"account_level"`
	IncidentReportId int `json:"record_id"`
	IncidentStatus string `json:"incident_status"`
	AlarmLevel string `json:"alarm_level"`
	Latitude float32 `json:"latitude"`
	Longitude float32 `json:"Longitude"`
	Address string `json:"address"`
	Remarks string `json:"remarks"`
	DateCreated time.Time `json:"date_created"`
	ReportId string `json:"report_id"`
}