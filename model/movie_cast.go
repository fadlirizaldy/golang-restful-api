package model

type Movie_cast struct {
	Id       int `json:"id" gorm:"primary_key, AUTO_INCREMENT"`
	Movie_id int `json:"movieid"`
	Cast_id  int `json:"castid"`
}