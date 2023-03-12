package homepage

import (
	"encoding/json"
	"fmt"
	"go-docker/helper"
	"net/http"
)

type HomePageResponse struct {
	Page    string `json:"page"`
	Message string `json:"message"`
}

func HomePageEntryPoint(responseWriter http.ResponseWriter, request *http.Request) {

	helper.SugarObj.Info(request)

	resp := &HomePageResponse{
		Page:    "Homepage",
		Message: "Success",
	}

	fmt.Println(" ==== HomePageEntryPoint ==== ", request)

	responseWriter.Header().Set("Content-Type", "application/json")

	json.NewEncoder(responseWriter).Encode(resp)

}
