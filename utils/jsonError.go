package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type MapErrorMessage struct {
	HttpRespCode int    `json:"http_resp_code"`
	GrpcRespCode int    `json:"grpc_resp_code,omitempty"`
	ErrCode      int    `json:"error_code,omitempty"`
	Description  string `json:"description"`
}

func (m *MapErrorMessage) Error() string {
	return fmt.Sprintf("%s", m.Description)
}

func MappError(w http.ResponseWriter, err error, code int) {
	payloadError := MapErrorMessage{
		HttpRespCode: code,
		Description:  err.Error(),
	}
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(payloadError)
}
