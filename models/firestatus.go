package models

type FireStatus struct {
	BaseModel
	IncidentId int `json:"incident_id" form:"incident_id" binding:"required"`
	Status string `json:"status form:"status" binding:"required"`
	ReportedBy int `json:"reported_by"`
}