package models

import (
	"time"
	"fmt"
	"github.com/jinzhu/gorm"
)

type Incident struct {
	BaseModel
	Latitude float32 `json:"latitude" form:"latitude" binding:"required"`
	Longitude float32 `json:"longitude" form:"longitude" binding:"required"`
	Address string `json:"address" form:"address" binding:"required"`
	Remarks string `json:"remarks" form:"remarks" binding:"required"`
}

func (i *Incident) AfterCreate(tx *gorm.DB) (err error) {
	loc,_ := time.LoadLocation("Asia/Manila")
	fmt.Printf("\n\nBEFORE CREATE --> %v\n",i.CreatedAt.Format(time.RFC3339))
	newCreatedAt,err1 := time.ParseInLocation(time.RFC3339,i.CreatedAt.Format(time.RFC3339),loc)
	newUpdatedAt,err2 := time.ParseInLocation(time.RFC3339,i.UpdatedAt.Format(time.RFC3339),loc)
	fmt.Printf("\n\nAFTER CREATE ---> %v\n",newCreatedAt)
	if err1 == nil && err2 == nil {
		if tx == nil {
				fmt.Printf("\n\nTX IS NULL!")
			} else {
				fmt.Printf("\n\nTX NOT NULL!")
				res1 := tx.Model(&i).Update("CreatedAt", newCreatedAt)
		    	res2 := tx.Model(&i).Update("UpdatedAt", newUpdatedAt)
		    	if res1 == nil && res2 == nil {
		    		fmt.Printf("\n\nSAVED!")
		    	}  else {
		    		fmt.Printf("\n ERROR IN PARSING DATE ---> %v\n\n",res1.Error.Error())
		    	}
			}
	} else {
		fmt.Printf("\n ERROR IN PARSING DATE ---> %v\n\n",err1,err2)
	}
    return
}