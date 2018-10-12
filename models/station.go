package models

type Station struct {
	BaseModel
	Region string `json:"region" form:"region" binding:"required"`
	RegionalOffice string `json:"regional_office" form:"regional_office" binding:"required"`
	Provice float64 `json:"province" form:"province" binding:"required"`
	Municipality float64 `json:"municipality" form:"municipality" binding:"required"`
	SubStation string `json:"sub_station" form:"sub_station"`
	UnitCode string `json:"unit_code" form:"unit_code" binding:"required"`
	Description string `json:"description" form:"description"`
	Remarks string `json:"remarks" form:"remarks"`
}
