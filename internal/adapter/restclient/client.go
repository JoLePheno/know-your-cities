package restclient

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/JoLePheno/know-your-cities/internal/model"
	"github.com/JoLePheno/know-your-cities/internal/port"
)

type Client struct {
	keyAPI  string
	baseURL string
	client  *http.Client
}

type city struct {
	PostCode string `json:"codePostal"`
	CodeCity string `json:"codeCommune"`
	CityName string `json:"nomCommune"`
	Name     string `json:"libelleAcheminement"`
}

func NewRestClient(baseURL, keyAPI string, httpClient *http.Client) *Client {
	return &Client{
		keyAPI:  keyAPI,
		baseURL: baseURL,
		client:  httpClient,
	}
}

func wrapError(code int, body io.ReadCloser) error {
	switch code {
	case 400:
		testError := &model.RestError{}
		err := json.NewDecoder(body).Decode(testError) // the error is part of the body
		if err != nil {
			return err
		}
		return fmt.Errorf("%w: %s", port.ErrInvalidZip, testError.Message.Content.Value)
	case 404:
		return port.ErrZipCodeNotFound
	default:
		return port.ErrInternalError
	}
}

func (c *Client) Do(code string, headers map[string]string) ([]*model.RestCity, error) {
	q := fmt.Sprintf("%s/%s", c.baseURL, code)

	req, err := http.NewRequest("GET", q, nil)
	if err != nil {
		return nil, err
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 { //The API only returns 200 if the code is found
		return nil, wrapError(resp.StatusCode, resp.Body)
	}

	cityResp := []*city{}
	err = json.NewDecoder(resp.Body).Decode(&cityResp) //decode the request body into struct, failed if any error occured
	if err != nil {
		return nil, err
	}

	var cities []*model.RestCity
	for _, city := range cityResp {
		cities = append(cities, &model.RestCity{
			PostCode: city.PostCode,
			CodeCity: city.CodeCity,
			CityName: city.CityName,
			Name:     city.Name,
		})
	}
	return cities, nil
}
