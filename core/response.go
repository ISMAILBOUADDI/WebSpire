package core

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	http.ResponseWriter
}

func NewResponse(w http.ResponseWriter) *Response {
	return &Response{ResponseWriter: w}
}

func (res *Response) SendJSON(data interface{}, statusCode int) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(statusCode)
	json.NewEncoder(res).Encode(data)
}

func (res *Response) SendText(data string, statusCode int) {
	res.WriteHeader(statusCode)
	res.Write([]byte(data))
}
