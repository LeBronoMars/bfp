package models

type RegionalOffice struct {
	BaseModel
	Name string `json:"name" form:"name" binding:"required"`
}
