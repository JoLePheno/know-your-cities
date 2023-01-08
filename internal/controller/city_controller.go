package controller

import (
	"fmt"
	"strings"

	"github.com/JoLePheno/know-your-cities/internal/adapter/city"
	"github.com/JoLePheno/know-your-cities/internal/adapter/hook"
	"github.com/JoLePheno/know-your-cities/internal/port"
)

type CityController struct {
	HookAdapter   *hook.HookAdapter
	Store         port.Store
	CitiesAdapter *city.CitiesAdapter
}

func (c *CityController) IsAValideCity(str string) error {
	if str == "" {
		return fmt.Errorf("%w: missing data", port.ErrInvalidData)
	}

	splitedData := strings.Split(str, ";")
	if len(splitedData) != 2 {
		return fmt.Errorf("%w: missing data or valid separator", port.ErrInvalidData)
	}
	id := splitedData[0]
	zipCode := splitedData[1]

	idConverted, err := c.HookAdapter.CheckId(id)
	if err != nil {
		return fmt.Errorf("%w: invalid ID", port.ErrInvalidData)
	}
	if !c.HookAdapter.CheckZipCode(zipCode) {
		return fmt.Errorf("%w: invalid Zip code", port.ErrInvalidData)
	}
	cities, err := c.CitiesAdapter.GetCitiesFromZipCode(idConverted, zipCode)
	if err != nil {
		return err
	}
	for _, city := range cities {
		err = c.Store.InsertCity(city)
		if err != nil {
			return err
		}
	}
	return nil
}
