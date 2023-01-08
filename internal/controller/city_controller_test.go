package controller

import (
	"net/http"
	"testing"

	"github.com/JoLePheno/know-your-cities/internal/adapter/city"
	"github.com/JoLePheno/know-your-cities/internal/adapter/hook"
	"github.com/JoLePheno/know-your-cities/internal/adapter/noop"
	"github.com/JoLePheno/know-your-cities/internal/adapter/restclient"
	"github.com/JoLePheno/know-your-cities/internal/port"
	"github.com/stretchr/testify/require"
)

type checkResp struct {
	isErrorExpected bool
	expectedError   error
}

var givenData = map[string]checkResp{
	"123;75010": {
		true,
		nil,
	},
	"123;1234": {
		false,
		port.ErrInvalidData,
	},
	"123,67000": {
		false,
		port.ErrInvalidData,
	},
	"123;75000": {
		false,
		port.ErrZipCodeNotFound,
	},
}

func TestCityController(t *testing.T) {
	restClient := restclient.NewRestClient("https://apicarto.ign.fr/api/codes-postaux/communes", "", http.DefaultClient)
	c := CityController{
		HookAdapter: &hook.HookAdapter{},
		Store:       &noop.Store{},
		CitiesAdapter: &city.CitiesAdapter{
			RestClient: restClient,
		},
	}

	for data, checkForError := range givenData {
		err := c.IsAValideCity(data)
		if checkForError.isErrorExpected {
			require.NoError(t, err)
		} else {
			require.Error(t, err, checkForError.expectedError)
		}
	}
}
