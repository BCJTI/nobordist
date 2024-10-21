package nobordist

import "net/http"

type GenericMessage struct {
	StatusCode int         `json:"statusCode"`
	Code       string      `json:"code"`
	Err        string      `json:"error"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

type ErrorMessage struct {
	Err *Error `json:"error"`
}

type Error struct {
	Message string `json:"message"`
	Data    *Data  `json:"data"`
}

type Data struct {
	Fields map[string]string
}

func (m GenericMessage) Error() string {

	if m.Err != "" {
		return m.Err
	}

	if m.Message != "" {
		return m.Message
	}

	if m.StatusCode != 0 {
		http.StatusText(m.StatusCode)
	}

	return ""

}

func (m ErrorMessage) Error() string {

	if m.Err != nil {
		return m.Err.Message
	}

	return ""

}
