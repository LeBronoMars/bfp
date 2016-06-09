package models

import (
	"time"
	"github.com/jinzhu/gorm"
)

type Incident struct {
	BaseModel
	Latitude float32 `json:"latitude" form:"latitude" binding:"required"`
	Longitude float32 `json:"longitude" form:"longitude" binding:"required"`
	Address string `json:"address" form:"address" binding:"required"`
	Remarks string `json:"remarks" form:"remarks" binding:"required"`
}

func (i *Incident) AfterCreate(tx *gorm.DB) (err error) {
	loc,_ := time.LoadLocation("Asia/Manila")
	newCreatedAt,_ := time.ParseInLocation(time.RFC3339,i.CreatedAt.Format(time.RFC3339),loc)
	newUpdatedAt,_ := time.ParseInLocation(time.RFC3339,i.UpdatedAt.Format(time.RFC3339),loc)
    tx.Model(i).Update("CreatedAt", newCreatedAt)
    tx.Model(i).Update("UpdatedAt", newUpdatedAt)
    return
}