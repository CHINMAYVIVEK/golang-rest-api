package healthcheck

import (
	"encoding/json"
	"fmt"
	helper "go-docker/helper"
	"net/http"
)

type HealthCheckResponse struct {
	HttpStatus int    `json:"httpStatus"`
	Message    string `json:"message"`
}

func HealthCheckEntryPoint(responseWriter http.ResponseWriter, request *http.Request) {

	helper.SugarObj.Info(request)

	resp := &HealthCheckResponse{
		HttpStatus: http.StatusOK,
		Message:    "Success",
	}

	fmt.Println(" ==== HealthCheckEntryPoint ==== ", request)

	responseWriter.Header().Set("Content-Type", "application/json")

	json.NewEncoder(responseWriter).Encode(resp)

}
