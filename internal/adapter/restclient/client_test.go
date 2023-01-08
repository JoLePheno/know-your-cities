package restclient

import (
	"net/http"
	"testing"

	"github.com/JoLePheno/know-your-cities/internal/model"
	"github.com/JoLePheno/know-your-cities/internal/port"
	"github.com/stretchr/testify/require"
)

func TestRestClientOK(t *testing.T) {
	expectedCity := []*model.RestCity{
		{
			PostCode: "67000",
			CodeCity: "67482",
			CityName: "Strasbourg",
			Name:     "STRASBOURG",
		},
	}

	c := NewRestClient("https://apicarto.ign.fr/api/codes-postaux/communes", "", http.DefaultClient)
	currentCity, err := c.Do(expectedCity[0].PostCode, map[string]string{
		"accept": "application/json",
	})
	require.NoError(t, err)
	require.EqualValues(t, expectedCity, currentCity)
}

func TestRestClientError(t *testing.T) {
	c := NewRestClient("https://almost_the_good_url/api/codes-postaux/communes", "", http.DefaultClient)
	_, err := c.Do("123", map[string]string{
		"accept": "application/json",
	})
	require.Error(t, err)
}

func TestRestClientInvalidZipCode(t *testing.T) {
	c := NewRestClient("https://apicarto.ign.fr/api/codes-postaux/communes", "", http.DefaultClient)
	_, err := c.Do("1234", map[string]string{
		"accept": "application/json",
	})
	require.Error(t, err, port.ErrZipCodeNotFound)
}
