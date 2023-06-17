package model

import (
	"encoding/json"
	"log"
	"net/http"
)

// errorResponse エラーレスポンスを表す。
type errorResponse struct {
	// Message エラーメッセージ
	Message string `json:"message"`
}

// WriteErrorResponse w にエラーレスポンスを書き込む。
func WriteErrorResponse(w http.ResponseWriter, code int, message string) {
	w.WriteHeader(code)

	resp := &errorResponse{
		Message: message,
	}

	if err := json.NewEncoder(w).Encode(&resp); err != nil {
		log.Printf("[ERROR] error response encoding failed: %+v\n", err)
	}
}
