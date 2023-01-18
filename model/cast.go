package model

import (
	"time"
)

type Cast struct {
	Id       	int     `json:"id" gorm:"primary_key, AUTO_INCREMENT"`
	Name    	string  `json:"name"`
	Birth_place string `json:"birth_place"`
	Birthday  	time.Time `json:"birthday" gorm:"birthday;type:datetime"`
	Rating  	int `json:"rating"`
	// Movies	[]Movie	`gorm:"many2many:movie_casts"`
}