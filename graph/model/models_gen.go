// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Cast struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	BirthPlace string `json:"birth_place"`
	Birthday   string `json:"birthday"`
	Rating     int    `json:"rating"`
}

type CastInput struct {
	Name       string `json:"name"`
	BirthPlace string `json:"birth_place"`
	Birthday   string `json:"birthday"`
	Rating     int    `json:"rating"`
}

type Movie struct {
	ID       string  `json:"id"`
	Title    string  `json:"title"`
	Language string  `json:"language"`
	Status   string  `json:"status"`
	Rating   float64 `json:"rating"`
	Casts    []*Cast `json:"casts" gorm:"many2many:movie_casts"`
}

type MovieInput struct {
	Title    string  `json:"title"`
	Language string  `json:"language"`
	Status   string  `json:"status"`
	Rating   float64 `json:"rating"`
}