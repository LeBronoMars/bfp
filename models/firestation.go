package models

import (
	"strings"
)

type FireStation struct {
	BaseModel
	StationName string `json:"station_name" form:"station_name" binding:"required"`
	StationCode string `json:"station_code" form:"station_code" binding:"required"`
	Latitude float64 `json:"latitude" form:"latitude" binding:"required"`
	Longitude float64 `json:"longitude" form:"longitude" binding:"required"`
	Region string `json:"region" form:"region" binding:"required"`
	RegionalOffice string `json:"regional_office" form:"regional_office" binding:"required"`
	Address string `json:"address" form:"address" binding:"required"`
	Email string `json:"email" form:"email" binding:"required"`
	ContactNo string `json:"contact_no" form:"contact_no" binding:"required"`
}

func (f *FireStation) BeforeCreate() (err error) {
	if len(strings.TrimSpace(f.Email)) < 1 {
		f.Email = "Not available"
	} 
	return
}	

func (f *FireStation) BeforeSave() (err error) {
	if len(strings.TrimSpace(f.Email)) < 1 {
		f.Email = "Not available"
	} 
	return
}	
