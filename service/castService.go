package service

import (
	"errors"
	"project_alterra/config"
	"project_alterra/graph/model"
)

type CastService struct {}

func (c *CastService) GetAllCasts() []*model.Cast{
	var casts []*model.Cast = []*model.Cast{}

	DB := config.InitDB()

	DB.Find(&casts)
	return casts
}

func (c *CastService) GetCastById(id string) (model.Cast, error){
	var cast model.Cast

	DB := config.InitDB()
	res := DB.First(&cast, "id = ?", id)

	if res.RowsAffected == 0 {
		return model.Cast{}, errors.New("cast not found")
	}

	return cast, nil
}

func (c *CastService) CreateCast(input model.CastInput) model.Cast{
	DB := config.InitDB()

	var newCast model.Cast = model.Cast{
		Name: input.Name,
		BirthPlace: input.BirthPlace,
		Birthday: input.Birthday,
		Rating: input.Rating,
	}

	DB.Create(&newCast)

	return newCast
}

func (c *CastService) DeleteCast(id string) (model.Cast, error){
	DB := config.InitDB()
	
	var cast model.Cast

	res := DB.Where("id = ?", id).Delete(&cast)

	if res.RowsAffected == 0 {
		return model.Cast{}, errors.New("cast not found")
	}

	return cast, nil
}

func (c *CastService) EditCast(id string, input model.CastInput) model.Cast{
	DB := config.InitDB()

	var newCast model.Cast = model.Cast{
		Name: input.Name,
		BirthPlace: input.BirthPlace,
		Birthday: input.Birthday,
		Rating: input.Rating,
	}

	DB.Where("id = ?", id).Updates(&newCast)

	return newCast
	
}



