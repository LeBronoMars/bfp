package models

type Incident struct {
	BaseModel
	Latitude float64 `json:"latitude" form:"latitude" binding:"required"`
	Longitude float64 `json:"longitude" form:"longitude" binding:"required"`
	Address string `json:"address" form:"address" binding:"required"`
	Remarks string `json:"remarks" form:"remarks" binding:"required"`
}