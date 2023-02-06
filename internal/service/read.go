package service

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/JoLePheno/know-your-cities/internal/adapter/csv"
	"github.com/JoLePheno/know-your-cities/pkg/logger"
	"github.com/JoLePheno/know-your-cities/pkg/utils"
)

type readRequest struct {
	Body string `json:"body"`
}

func (s *ReaderService) Read(ctx context.Context) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Ctx(ctx).Info().Msg("post /read")
		req := &readRequest{}

		err := json.NewDecoder(r.Body).Decode(req) //decode the request body into struct, failed if any error occured
		if err != nil {
			logger.Ctx(ctx).Error().Msgf("An error occurred while decoding request, err: " + err.Error())
			utils.Respond(w, utils.Message(false, "Invalid request"), 400)
			return
		}
		s.ReaderController.Reader = csv.NewReader(strings.NewReader(req.Body), "")
		resp, err := s.ReaderController.RunReaderController(ctx)
		if err != nil {
			logger.Ctx(ctx).Error().Msgf("an error occurred: " + err.Error())
			utils.Respond(w, utils.Message(false, "Internal error"), 500)
			return
		}

		utils.Respond(w, resp, 200)
	})
}
