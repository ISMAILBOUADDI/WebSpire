package core

import (
	"encoding/json"
	"net/http"
)

type Request struct {
	*http.Request
}

func NewRequest(r *http.Request) *Request {
	return &Request{Request: r}
}

func (req *Request) GetBody() (map[string]interface{}, error) {
	var body map[string]interface{}
	err := json.NewDecoder(req.Body).Decode(&body)
	return body, err
}
