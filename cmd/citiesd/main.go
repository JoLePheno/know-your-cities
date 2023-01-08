package main

import (
	"context"
	"encoding/csv"
	"log"
	"net/http"
	"os"

	"github.com/JoLePheno/know-your-cities/internal/adapter/city"
	"github.com/JoLePheno/know-your-cities/internal/adapter/hook"
	"github.com/JoLePheno/know-your-cities/internal/adapter/noop"
	"github.com/JoLePheno/know-your-cities/internal/adapter/restclient"
	"github.com/JoLePheno/know-your-cities/internal/controller"
	"github.com/JoLePheno/know-your-cities/pkg/logger"
)

func main() {
	fileName := "data.csv"
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	c := controller.ReaderController{
		Reader: csv.NewReader(f),
		CityController: &controller.CityController{
			HookAdapter: &hook.HookAdapter{},
			Store:       &noop.Store{},
			CitiesAdapter: &city.CitiesAdapter{
				RestClient: restclient.NewRestClient("https://apicarto.ign.fr/api/codes-postaux/communes", "", http.DefaultClient),
			},
		},
	}
	ctx := logger.WithContext(context.Background())
	c.RunReaderController(ctx)
}
