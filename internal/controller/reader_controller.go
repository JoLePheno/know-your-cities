package controller

import (
	"context"
	"errors"
	"strings"
	"sync"

	"github.com/JoLePheno/know-your-cities/internal/port"
	"github.com/JoLePheno/know-your-cities/pkg/logger"
)

type ReaderController struct {
	Reader         port.Reader
	CityController *CityController
	TotalFileLine  int
}

type Resp struct {
	content map[string]interface{}
	mu      sync.Mutex
}

func (c *ReaderController) RunReaderController(ctx context.Context) (map[string]interface{}, error) {
	res := &Resp{content: map[string]interface{}{}}
	wg := sync.WaitGroup{}

	for {
		data, err := c.Reader.Read()
		if err != nil {
			if errors.Is(err, port.ErrEOF) {
				break
			}
			return nil, err
		}
		wg.Add(len(data))

		go func(data []string) {
			defer wg.Done()

			for _, str := range data {
				err = c.CityController.IsAValideCity(strings.TrimSpace(str))
				res.mu.Lock()
				if err != nil {
					logger.Ctx(ctx).Warn().Msgf("an error occured during city validation: %s", err.Error())
					res.content[str] = err
				} else {
					res.content[str] = "OK"
				}
				res.mu.Unlock()
			}
		}(data)
	}

	wg.Wait()
	for k, v := range res.content {
		logger.Ctx(ctx).Info().Msgf("%s : %v", k, v)
	}

	return res.content, nil
}
