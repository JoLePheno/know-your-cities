package model

type RestCity struct {
	PostCode string
	CodeCity string
	CityName string
	Name     string
}

type RestError struct {
	Code    int          `json:"code"`
	Message ErrorMessage `json:"message"`
}

type ErrorMessage struct {
	Content PostalCodeError `json:"codePostal"`
}

type PostalCodeError struct {
	Value    string `json:"value"`
	Msg      string `json:"msg"`
	Param    string `json:"param"`
	Location string `json:"location"`
}
