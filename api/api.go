package api

import (
	"encoding/json"
	"net/http"
)

// coins balance params
type CoinsBalanceParams struct {
	Username string
}

// coin balance response
type CoinsBalanceResponse struct {
	StatusCode int32 // the status code of response
	Balance    int64 // the balance
}

type Error struct {
	Code    int    // error code
	Message string // message
}

func writeError(writer http.ResponseWriter, message string , code int){
	resp:=Error{
		Code: code,
		Message: message,
	}
	writer.Header().Set("Content-Type","application/json")
	writer.WriteHeader(code)

	json.NewEncoder(writer).Encode(resp)
}

var (
	RequestErrorHandler =func(writer http.ResponseWriter, err error){
		writeError(writer,err.Error(),http.StatusBadRequest)
	}
	InternalErrorHandler = func(writer http.ResponseWriter, err error){
		writeError(writer,"An Unexpected Error occurred ",http.StatusInternalServerError)
	}
)

