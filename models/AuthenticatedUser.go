package models

import "time"

type AuthenticatedUser struct {
	Id int `json:"id"`
	FirstName  string `json:"first_name"`
	LastName  string `json:"last_name"`	
	Status string `json:"status"`
	UserRole string `json:"user_role"`
	UserLevel string `json:"user_level"`
	Email string `json:"email"`
	IsPasswordDefault bool `json:"is_password_default"`
	DateCreated time.Time `json:"date_created"`
	DateUpdated time.Time `json:"date_updated"`
	Token string `json:"token"`
}