package models

type CommandCenter struct {
	BaseModel
	Name string `json:"name" form:"name" binding:"required"`
	ContactNo string `json:"contact_no" form:"contact_no" binding:"required"`
	Email string `json:"email" form:"email" binding:"required"`
}
