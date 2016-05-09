package models

import (
	"time"
)

type Incident struct {
	Id int `json:"id"`
	ReportedBy int `json:"reported_by"`
	Status string `json:"status"`
	AlarmLevel string `json:"alarm_level"`
	Latitude float32 `json:"latitude"`
	Longitude float32 `json:"Longitude"`
	Address string `json:"address"`
	Remarks string `json:"remarks"`
	DateCreated time.Time `json:"date_created"`
}