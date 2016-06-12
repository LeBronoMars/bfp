package models

type User struct {
	BaseModel
	FirstName  string `json:"first_name" form:"first_name" binding:"required"`
	LastName  string `json:"last_name" form:"last_name" binding:"required"`	
	Email string `json:"email" form:"email" binding:"required"`
	ContactNo string `json:"contact_no" form:"contact_no" binding:"required"`
	Status string `json:"status"`
	Userrole string `json:"user_role" form:"user_role" binding:"required"`
	Userlevel string `json:"user_level" form:"user_level" binding:"required"`
	Password string `json:"password"`
	IsPasswordDefault bool `json:"is_password_default"`
}

func (u *User) BeforeCreate() (err error) {
	u.IsPasswordDefault = true
	u.Status = "active"
	return
}