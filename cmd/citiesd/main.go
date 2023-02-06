package main

import (
	"context"
	"log"
	"net/http"

	"github.com/JoLePheno/know-your-cities/internal/adapter/city"
	"github.com/JoLePheno/know-your-cities/internal/adapter/hook"

	"github.com/JoLePheno/know-your-cities/internal/adapter/noop"
	"github.com/JoLePheno/know-your-cities/internal/adapter/restclient"
	"github.com/JoLePheno/know-your-cities/internal/controller"
	"github.com/JoLePheno/know-your-cities/internal/service"
	"github.com/JoLePheno/know-your-cities/pkg/logger"
)

func main() {
	ctx := logger.WithContext(context.Background())
	s := service.ReaderService{
		ReaderController: &controller.ReaderController{
			Reader: nil,
			CityController: &controller.CityController{
				HookAdapter: &hook.HookAdapter{},
				Store:       &noop.Store{},
				CitiesAdapter: &city.CitiesAdapter{
					RestClient: restclient.NewRestClient("https://apicarto.ign.fr/api/codes-postaux/communes", "", http.DefaultClient),
				},
			},
		},
	}

	log.Fatal(http.ListenAndServe(":8080", s.Router(ctx)))
}
