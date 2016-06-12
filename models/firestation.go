package models

import (
	"strings"
	"errors"
)

type FireStation struct {
	BaseModel
	StationName string `json:"station_name" form:"station_name" binding:"required"`
	StationCode string `json:"station_code" form:"station_code" binding:"required"`
	Latitude float64 `json:"latitude" form:"latitude" binding:"required"`
	Longitude float64 `json:"longitude" form:"longitude" binding:"required"`
	Address string `json:"address" form:"address" binding:"required"`
	ContactNo string `json:"contact_no" form:"contact_no" binding:"required"`
	Email string `json:"email" form:"email" binding:"required"`
}

func (f *FireStation) BeforeCreate() (err error) {
	if len(strings.TrimSpace(f.Email)) < 1 {
		f.Email = "Not available"
	} else if len(strings.TrimSpace(f.ContactNo)) < 1 {
		err = errors.New("Contact no is required")
	}
	return
}	

func (f *FireStation) BeforeSave() (err error) {
	if len(strings.TrimSpace(f.Email)) < 1 {
		f.Email = "Not available"
	} else if len(strings.TrimSpace(f.ContactNo)) < 1 {
		err = errors.New("Contact no is required")
	}
	return
}	
