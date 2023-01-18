package model

type Movie_cast struct {
	Id       int `json:"id"`
	Movie_id int `json:"movie_id"`
	Cast_id  int `json:"cast_id"`
}