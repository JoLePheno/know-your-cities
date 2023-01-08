package city

import (
	"github.com/JoLePheno/know-your-cities/internal/adapter/restclient"
	"github.com/JoLePheno/know-your-cities/internal/model"
)

type CitiesAdapter struct {
	RestClient *restclient.Client
}

func (c *CitiesAdapter) GetCitiesFromZipCode(id uint64, zip string) ([]*model.CityModel, error) {
	restCities, err := c.RestClient.Do(zip, map[string]string{
		"accept": "application/json",
	})
	if err != nil {
		return nil, err
	}

	cities := []*model.CityModel{}
	for _, city := range restCities {
		cities = append(cities, &model.CityModel{
			ID:      id,
			ZipCode: city.PostCode,
			Name:    city.CityName,
		})
	}
	return cities, nil
}
