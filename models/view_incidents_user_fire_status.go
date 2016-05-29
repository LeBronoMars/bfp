package models

import "time"

type QryIncidents struct {
	FireStatusId int `json:"fire_status_id"`
	FireStatusReported time.Time `json:"fire_status_reported"`
	FireStatus string `json:"fire_status"`
	ReporterId int `json:"reporter_id"`
	ReporterFirstName string `json:"reporter_first_name"`
	ReporterLastName string `json:"reporter_last_name"`
	ReporterRole string `json:"reporter_role"`
}