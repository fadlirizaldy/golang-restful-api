package service

import (
	"errors"
	"project_alterra/config"
	"project_alterra/graph/model"
)

type MovieService struct {}

func (m *MovieService) GetAllMovies() []*model.Movie{
	var movies []*model.Movie = []*model.Movie{}

	DB := config.InitDB()

	DB.Preload("Casts").Select("id","title").Find(&movies)
	return movies
}

func (m *MovieService) GetMovieById(id string) (model.Movie, error){
	var movie model.Movie

	DB := config.InitDB()
	res := DB.Preload("Casts").Where("id = ?", id).First(&movie)

	if res.RowsAffected == 0 {
		return model.Movie{}, errors.New("movie not found")
	}

	return movie, nil
}

func (m *MovieService) CreateMovie(input model.MovieInput) model.Movie{
	DB := config.InitDB()

	var newMovie model.Movie = model.Movie{
		Title: input.Title,
		Language: input.Language,
		Status: input.Status,
		Rating: input.Rating,
	}

	DB.Create(&newMovie)

	return newMovie
}

func (m *MovieService) DeleteMovie(id string) (model.Movie, error){
	DB := config.InitDB()
	
	var movie model.Movie

	res := DB.Where("id = ?", id).Delete(&movie)

	if res.RowsAffected == 0 {
		return model.Movie{}, errors.New("movie not found")
	}

	return movie, nil
}

func (m *MovieService) EditMovie(id string, input model.MovieInput) model.Movie{
	DB := config.InitDB()

	var newMovie model.Movie = model.Movie{
		Title: input.Title,
		Language: input.Language,
		Status: input.Status,
		Rating: input.Rating,
	}

	DB.Where("id = ?", id).Updates(&newMovie)

	return newMovie
	
}



