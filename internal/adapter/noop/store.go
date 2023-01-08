package noop

import "github.com/JoLePheno/know-your-cities/internal/model"

type Store struct{}

func (*Store) InsertCity(city *model.CityModel) error {
	return nil
}

func (*Store) GetCityByName(name string) (*model.CityModel, error) {
	return nil, nil
}
