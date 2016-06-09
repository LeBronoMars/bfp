package models

import "time"

type Incident struct {
	BaseModel
	Latitude float32 `json:"latitude" form:"latitude" binding:"required"`
	Longitude float32 `json:"longitude" form:"longitude" binding:"required"`
	Address string `json:"address" form:"address" binding:"required"`
	Remarks string `json:"remarks" form:"remarks" binding:"required"`
}

func (i *Incident) BeforeCreate() (err error) {
	loc,_ := time.LoadLocation("Asia/Manila")
	newCreatedAt,_ := time.ParseInLocation(time.RFC3339,i.CreatedAt.String(),loc)
	newUpdatedAt,_ := time.ParseInLocation(time.RFC3339,i.UpdatedAt.String(),loc)

	i.CreatedAt = newCreatedAt
	i.UpdatedAt = newUpdatedAt
	return
}