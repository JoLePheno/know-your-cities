package controller

import (
	"net/http"
	"os"
	"testing"

	"github.com/JoLePheno/know-your-cities/internal/adapter/city"
	"github.com/JoLePheno/know-your-cities/internal/adapter/csv"
	"github.com/JoLePheno/know-your-cities/internal/adapter/hook"
	"github.com/JoLePheno/know-your-cities/internal/adapter/restclient"
	"github.com/stretchr/testify/require"
)

func TestReaderControllerGetFileName(t *testing.T) {
	fileName := "../../pkg/test/data.csv"
	f, err := os.Open(fileName)
	require.NoError(t, err)
	// remember to close the file at the end of the program
	defer f.Close()

	c := ReaderController{
		Reader: csv.NewReader(f),
		CityController: &CityController{
			HookAdapter: &hook.HookAdapter{},
			Store:       nil,
			CitiesAdapter: &city.CitiesAdapter{
				RestClient: restclient.NewRestClient("https://apicarto.ign.fr/api/codes-postaux/communes", "", http.DefaultClient),
			},
		},
	}
	require.EqualValues(t, fileName, c.Reader.(*csv.CSVAdapter).GetFileName())
}
