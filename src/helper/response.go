package helper

import (
	"encoding/json"
	"net/http"
)

const (
	statusSuccess = "success"
	statusFail    = "fail"
	statusError   = "error"
)

type StdResponse interface {
	reply(status string, data any, message string)
	ReplySuccess(data any)
	ReplyFail(data any)
	ReplyError(message string)
	SetHttpStatusCode(code int) StdResponse
}

type StdResp struct {
	w              http.ResponseWriter
	httpStatusCode int
	Status         string `json:"status,omitempty"`
	Data           any    `json:"data,omitempty"`
	Message        string `json:"message,omitempty"`
}

func (sr *StdResp) reply(status string, data any, message string) {
	sr.w.Header().Set("Content-Type", "application/json")
	if sr.httpStatusCode != 0 {
		sr.w.WriteHeader(sr.httpStatusCode)
	}

	sr.Status = status
	sr.Data = data
	sr.Message = message

	json.NewEncoder(sr.w).Encode(sr)
}

func (sr *StdResp) SetHttpStatusCode(code int) StdResponse {
	sr.httpStatusCode = code
	return sr
}

func (sr *StdResp) ReplySuccess(data any) {
	sr.reply(statusSuccess, data, "")
}

func (sr *StdResp) ReplyFail(data any) {
	sr.reply(statusFail, data, "")
}

func (sr *StdResp) ReplyError(message string) {
	sr.reply(statusError, nil, message)
}

func PlugResponse(w http.ResponseWriter) StdResponse {
	res := &StdResp{
		w:              w,
		httpStatusCode: 200,
		Status:         "",
		Data:           nil,
		Message:        "",
	}
	return res
}
