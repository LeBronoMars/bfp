package models

import (
	"time"
	"fmt"
)

type Incident struct {
	BaseModel
	Latitude float32 `json:"latitude" form:"latitude" binding:"required"`
	Longitude float32 `json:"longitude" form:"longitude" binding:"required"`
	Address string `json:"address" form:"address" binding:"required"`
	Remarks string `json:"remarks" form:"remarks" binding:"required"`
}

func (i *Incident) AfterSave() (err error) {
	fmt.Printf("\nBEFORE PARSE --> %v",i.CreatedAt.Format(time.RFC3339))
	loc,_ := time.LoadLocation("Asia/Manila")
	newCreatedAt,err1 := time.ParseInLocation(time.RFC3339,i.CreatedAt.Format(time.RFC3339),loc)
	newUpdatedAt,err2 := time.ParseInLocation(time.RFC3339,i.UpdatedAt.Format(time.RFC3339),loc)

	if err1 == nil && err2 == nil {
		fmt.Printf("\nAFTER PARSE FIRE STATUS ---> %v\n\n",newCreatedAt)
		i.CreatedAt = newCreatedAt
		i.UpdatedAt = newUpdatedAt
	} else {
		fmt.Printf("\nERROR IN PARSING ---> %v\n\n",err1)
	}
	return
}