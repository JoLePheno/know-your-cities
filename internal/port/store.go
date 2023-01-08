package port

import "github.com/JoLePheno/know-your-cities/internal/model"

type Store interface {
	InsertCity(city *model.CityModel) error
	GetCityByName(name string) (*model.CityModel, error)
}
