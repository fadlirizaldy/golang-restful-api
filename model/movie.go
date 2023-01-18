package model

type Movie struct {
	Id       int     `json:"id" gorm:"primary_key, AUTO_INCREMENT"`
	Title    string  `json:"title"`
	Language string  `json:"language"`
	Status   string  `json:"status"`
	Rating   float32 `json:"rating"`
	Casts    []Cast  `gorm:"many2many:movie_casts"`
}
