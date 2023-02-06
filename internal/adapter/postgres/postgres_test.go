package postgres

import (
	"testing"

	"github.com/JoLePheno/know-your-cities/internal/model"
	"github.com/JoLePheno/know-your-cities/internal/port"
	"github.com/JoLePheno/know-your-cities/internal/port/mock_port"
	"github.com/golang/mock/gomock"
)

func TestInsertCity(t *testing.T) {
	newCity01 := &model.CityModel{
		ID:      1234,
		ZipCode: "57870",
		Name:    "Hartzviller",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockStore := mock_port.NewMockStore(ctrl)
	mockStore.EXPECT().InsertCity(newCity01).Return(nil)
	mockStore.InsertCity(newCity01)
}

func TestInsertCity2(t *testing.T) {
	newCity01 := &model.CityModel{
		ID:      1234,
		ZipCode: "57870",
		Name:    "Hartzviller",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockStore := mock_port.NewMockStore(ctrl)
	mockStore.EXPECT().InsertCity(newCity01).Return(nil)
	mockStore.InsertCity(newCity01)
}

func TestUpdateExistingCity(t *testing.T) {
	newCity01 := &model.CityModel{
		ID:      1234,
		ZipCode: "67000",
		Name:    "Strasbourg",
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockStore := mock_port.NewMockStore(ctrl)

	mockStore.EXPECT().InsertCity(newCity01).Return(nil)
	mockStore.InsertCity(newCity01)

	mockStore.EXPECT().GetCityByName(newCity01.Name).Return(newCity01, nil)
	mockStore.GetCityByName(newCity01.Name)

	newCity01.ID = 12345
	mockStore.EXPECT().InsertCity(newCity01).Return(nil)
	mockStore.InsertCity(newCity01)

	mockStore.EXPECT().GetCityByName(newCity01.Name).Return(newCity01, nil)
	mockStore.GetCityByName(newCity01.Name)
}

func TestNotFound(t *testing.T) {
	newCity01 := &model.CityModel{
		ID:      1234,
		ZipCode: "67000",
		Name:    "Strasbourg",
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockStore := mock_port.NewMockStore(ctrl)

	mockStore.EXPECT().GetCityByName(newCity01.Name).Return(nil, port.ErrPgCityNotFound)
	mockStore.GetCityByName(newCity01.Name)
}
