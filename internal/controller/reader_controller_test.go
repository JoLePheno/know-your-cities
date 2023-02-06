package controller

import (
	"context"
	"net/http"
	"strings"
	"testing"

	"github.com/JoLePheno/know-your-cities/internal/adapter/city"
	"github.com/JoLePheno/know-your-cities/internal/adapter/csv"
	"github.com/JoLePheno/know-your-cities/internal/adapter/hook"
	"github.com/JoLePheno/know-your-cities/internal/adapter/noop"
	"github.com/JoLePheno/know-your-cities/internal/adapter/restclient"
	"github.com/stretchr/testify/require"
)

var in string = `1;75000
foo;bar
2;75001
4;baz
123456;67000
98;57870
3;69125
45;69620
34;69380
26;69420
89;69490
87;69430
27;69690
67;69840
7;69500
654;69770
987;69620
98765;69410
89;69840
76543;69680`

func TestReaderControllerGetFileName(t *testing.T) {
	fileName := "test"
	c := ReaderController{
		Reader: csv.NewReader(strings.NewReader(in), fileName),
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

func TestReaderController(t *testing.T) {
	fileName := "test"
	c := ReaderController{
		Reader: csv.NewReader(strings.NewReader(in), fileName),
		CityController: &CityController{
			HookAdapter: &hook.HookAdapter{},
			Store:       &noop.Store{},
			CitiesAdapter: &city.CitiesAdapter{
				RestClient: restclient.NewRestClient("https://apicarto.ign.fr/api/codes-postaux/communes", "", http.DefaultClient),
			},
		},
	}

	resp, err := c.RunReaderController(context.Background())
	require.NoError(t, err)
	totalLine, err := c.Reader.(*csv.CSVAdapter).LineCounter(strings.NewReader(in))
	require.NoError(t, err)
	require.Len(t, resp, totalLine)
}
