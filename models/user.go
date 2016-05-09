package models

import (
	"time"
)

type User struct {
	Id int `json:"id"`
	FirstName  string `json:"first_name"`
	LastName  string `json:"last_name"`	
	Status string `json:"status"`
	Userrole string `json:"user_role"`
	Userlevel string `json:"user_level"`
	Username string `json:"username"`
	Password string `json:"password"`
	DateCreated time.Time `json:"date_created"`
	DateUpdated time.Time `json:"date_updated"`
}