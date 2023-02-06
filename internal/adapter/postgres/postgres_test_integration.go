package postgres

import (
	"testing"

	"github.com/JoLePheno/know-your-cities/internal/model"
	"github.com/JoLePheno/know-your-cities/internal/port"
	"github.com/stretchr/testify/require"
)

func TestIntegrationInsertCity(t *testing.T) {
	newCity01 := &model.CityModel{
		ID:      1234,
		ZipCode: "57870",
		Name:    "Hartzviller",
	}

	newCity02 := &model.CityModel{
		ID:      12345,
		ZipCode: "57870",
		Name:    "Walscheid",
	}
	store := NewCitiesStore()

	defer func() {
		_, err := store.db.Exec(`truncate "cities" CASCADE`)
		require.NoError(t, err)
	}()

	err := store.InsertCity(newCity01)
	require.NoError(t, err)

	err = store.InsertCity(newCity02)
	require.NoError(t, err)

	storedCity01, err := store.GetCityByName(newCity01.Name)
	require.NoError(t, err)
	require.EqualValues(t, newCity01, storedCity01)

	storedCity02, err := store.GetCityByName(newCity02.Name)
	require.NoError(t, err)
	require.EqualValues(t, newCity02, storedCity02)
}

func TestIntegrationUpdateExistingCity(t *testing.T) {
	newCity01 := &model.CityModel{
		ID:      1234,
		ZipCode: "67000",
		Name:    "Strasbourg",
	}

	store := NewCitiesStore()

	defer func() {
		_, err := store.db.Exec(`truncate "cities" CASCADE`)
		require.NoError(t, err)
	}()

	err := store.InsertCity(newCity01)
	require.NoError(t, err)

	storedCity01, err := store.GetCityByName(newCity01.Name)
	require.NoError(t, err)
	require.EqualValues(t, newCity01, storedCity01)

	newCity01.ID = 12345
	err = store.InsertCity(newCity01)
	require.NoError(t, err)

	storedCity01, err = store.GetCityByName(newCity01.Name)
	require.NoError(t, err)
	require.EqualValues(t, newCity01, storedCity01)
}

func TestIntegrationNotFound(t *testing.T) {
	newCity01 := &model.CityModel{
		ID:      1234,
		ZipCode: "67000",
		Name:    "Strasbourg",
	}

	store := NewCitiesStore()

	_, err := store.GetCityByName(newCity01.Name)
	require.Error(t, err, port.ErrPgCityNotFound)
}


