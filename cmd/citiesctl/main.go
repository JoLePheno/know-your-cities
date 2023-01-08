package main

import (
	"context"
	"encoding/csv"
	"log"
	"net/http"
	"os"

	"github.com/JoLePheno/know-your-cities/internal/adapter/city"
	"github.com/JoLePheno/know-your-cities/internal/adapter/hook"
	"github.com/JoLePheno/know-your-cities/internal/adapter/postgres"
	"github.com/JoLePheno/know-your-cities/internal/adapter/restclient"
	"github.com/JoLePheno/know-your-cities/internal/controller"
	"github.com/JoLePheno/know-your-cities/pkg/logger"
	"github.com/caarlos0/env"
)

type config struct {
	RestClientEndpoint string `env:"REST_CLIENT_ENDPOINT" envDefault:"https://apicarto.ign.fr/api/codes-postaux/communes"`
	PathToFile         string `env:"FILE_PATH" envDefault:"pkg/test/data.csv"`
}

func main() {
	var cfg config
	if err := env.Parse(&cfg); err != nil {
		log.Fatal("parsing envs failed: %w", err)
	}

	f, err := os.Open(cfg.PathToFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	store := postgres.NewCitiesStore()
	c := controller.ReaderController{
		Reader: csv.NewReader(f),
		CityController: &controller.CityController{
			HookAdapter: &hook.HookAdapter{},
			Store:       store,
			CitiesAdapter: &city.CitiesAdapter{
				RestClient: restclient.NewRestClient(cfg.RestClientEndpoint, "", http.DefaultClient),
			},
		},
	}

	ctx := logger.WithContext(context.Background())
	c.RunReaderController(ctx)
}
