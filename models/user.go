package models

type User struct {
	BaseModel
	FirstName  string `json:"first_name" form:"first_name" binding:"required"`
	LastName  string `json:"last_name" form:"last_name" binding:"required"`	
	Email string `json:"email" form:"email" binding:"required"`
	ContactNo string `json:"contact_no" form:"contact_no" binding:"required"`
	Status string `json:"status"`
	UserRole string `json:"user_role" form:"user_role" binding:"required"`
	UserLevel string `json:"user_level" form:"user_level" binding:"required"`
	Password string `json:"-"`
	IsPasswordDefault bool `json:"is_password_default"`
	StationId int `json:"station_id" form:"station_id"`
}

func (u *User) BeforeCreate() (err error) {
	u.IsPasswordDefault = true
	u.Status = "active"
	return
}