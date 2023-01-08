package postgres

import (
	"fmt"
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
		_, err := s.GetCityByName(city.Name)
		currentTime := time.Now()
		c := &CityDBO{
			UpdatedAt:  currentTime,
			Name:       city.Name,
			ExternalID: city.ID,
			PostCode:   city.ZipCode,
		}
		if err == port.ErrPgCityNotFound {
			c.CreatedAt = currentTime
			return tx.Insert(c)
		}
		_, err = tx.Model(c).Where("name = ?", city.Name).
			Set("updated_at = ?", currentTime).
			Set("external_id = ?", city.ID).Update()
		return err
	},
	)
	if err != nil {
		fmt.Println("error when inserting city, err %w", err)
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
		fmt.Println("error when fetching city, err %w", err)
		return nil, err
	}
	return convertCityModelDBOToModel(&cityDBO), nil
}

func convertCityModelDBOToModel(cityDBO *CityDBO) *model.CityModel {
	return &model.CityModel{
		ID:      cityDBO.ExternalID,
		ZipCode: cityDBO.PostCode,
		Name:    cityDBO.Name,
	}
}
