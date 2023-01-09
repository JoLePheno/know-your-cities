package postgres

import (
	"time"

	"github.com/JoLePheno/know-your-cities/internal/model"
	"github.com/JoLePheno/know-your-cities/internal/port"
	"github.com/go-pg/pg"
)

type CityDBO struct {
	tableName struct{} `sql:"cities"`

	ID         uint64 `sql:"id"`
	ExternalID uint64 `sql:"external_id"`

	CreatedAt time.Time `sql:"created_at,notnull"`
	UpdatedAt time.Time `sql:"updated_at,notnull"`

	Name     string `sql:"name"`
	PostCode string `sql:"post_code"`
}

func (s *CitiesStore) InsertCity(city *model.CityModel) error {
	err := s.db.RunInTransaction(func(tx *pg.Tx) error {
		c := &CityDBO{}
		currentTime := time.Now()
		err := s.db.Model(c).Where("name = ?", city.Name).Select()
		if err != nil {
			if err == pg.ErrNoRows {
				c.CreatedAt = currentTime
				c.UpdatedAt = currentTime
				c.Name = city.Name
				c.ExternalID = city.ID
				c.PostCode = city.ZipCode
				return tx.Insert(c)
			}

			storedCity := convertCityModelDBOToModel(c)
			if storedCity.ZipCode == city.ZipCode && storedCity.Name == city.Name && storedCity.ID != city.ID {
				return port.ErrCityAlreadyStored
			}
			return err
		}

		_, err = tx.Model(c).Where("name = ?", city.Name).
			Set("updated_at = ?", currentTime).
			Set("external_id = ?", city.ID).Update()
		return err
	})

	if err != nil {
		return err
	}
	return nil
}

func (s *CitiesStore) GetCityByName(name string) (*model.CityModel, error) {
	var cityDBO CityDBO
	err := s.db.RunInTransaction(func(tx *pg.Tx) error {
		return tx.Model(&cityDBO).Where("name = ?", name).Select()
	},
	)
	if err != nil {
		if err == pg.ErrNoRows {
			return nil, port.ErrPgCityNotFound
		}
		return convertCityModelDBOToModel(&cityDBO), err
	}
	return convertCityModelDBOToModel(&cityDBO), nil
}

func convertCityModelDBOToModel(cityDBO *CityDBO) *model.CityModel {
	if cityDBO == nil {
		return nil
	}
	return &model.CityModel{
		ID:      cityDBO.ExternalID,
		ZipCode: cityDBO.PostCode,
		Name:    cityDBO.Name,
	}
}
