package models

import (
	"time"
	"fmt"
)

type FireStatus struct {
	BaseModel
	IncidentId int `json:"incident_id" form:"incident_id" binding:"required"`
	Status string `json:"status" form:"status" binding:"required"`
	ReportedBy int `json:"reported_by" form:"reported_by" binding:"required"`
}

func (f *FireStatus) AfterSave() (err error) {
	fmt.Printf("\nBEFORE PARSE FIRE STATUS --> %v",f.CreatedAt.Format(time.RFC3339))
	loc,_ := time.LoadLocation("Asia/Manila")
	newCreatedAt,err1 := time.ParseInLocation(time.RFC3339,f.CreatedAt.Format(time.RFC3339),loc)
	newUpdatedAt,err2 := time.ParseInLocation(time.RFC3339,f.UpdatedAt.Format(time.RFC3339),loc)

	if err1 == nil && err2 == nil {
		fmt.Printf("\nAFTER PARSE FIRE STATUS ---> %v\n\n",newCreatedAt)
		f.CreatedAt = newCreatedAt
		f.UpdatedAt = newUpdatedAt
	} else {
		fmt.Printf("\nERROR IN PARSING FIRE STATUS ---> %v\n\n",err1)
	}
	return
}