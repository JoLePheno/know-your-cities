package service

import (
	"context"
	"net/http"

	"github.com/JoLePheno/know-your-cities/internal/controller"
	"github.com/JoLePheno/know-your-cities/pkg/utils"
	"github.com/gorilla/mux"
)

type ReaderService struct {
	ReaderController *controller.ReaderController
}

func (s *ReaderService) Router(ctx context.Context) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", HomeHandler).Methods("GET")
	r.Methods("POST").Name("Read").Handler(s.Read(ctx)).Path("/read")

	return r
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	message := utils.Message(true, "CSV reader")
	utils.Respond(w, message, 200)
}
