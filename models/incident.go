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

func (i *Incident) BeforeCreate() (err error) {
	fmt.Printf("\nBEFORE PARSE --> %v",i.CreatedAt.String())
	loc,_ := time.LoadLocation("Asia/Manila")
	newCreatedAt,err1 := time.ParseInLocation(time.RFC3339,i.CreatedAt.String(),loc)
	newUpdatedAt,err2 := time.ParseInLocation(time.RFC3339,i.UpdatedAt.String(),loc)

	fmt.Printf("\nAFTER PARSE ---> %v\n\n",newCreatedAt)
	if err1 == nil && err2 == nil {
		i.CreatedAt = newCreatedAt
		i.UpdatedAt = newUpdatedAt
	} else {
		fmt.Printf("\nERROR IN PARSING ---> %v\n\n",err1)
	}
	return
}