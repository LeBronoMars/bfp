package models

type Status struct {
	BaseModel
	StatusName string `json:"status_name" form:"status_name" binding:"required"`
	StatusColor string `json:"status_color" form:"status_color" binding:"required"`
}
