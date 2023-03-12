package server

import (
	"go-docker/app/healthcheck"
	"go-docker/app/homepage"
	"go-docker/app/redispkg"

	"github.com/gorilla/mux"
)

func InitializeRouter(router *mux.Router) {

	var apiV1 = router.PathPrefix("/v1").Subrouter()

	apiV1.HandleFunc("/health-check", healthcheck.HealthCheckEntryPoint).Methods("GET")

	apiV1.HandleFunc("/homepage", homepage.HomePageEntryPoint).Methods("GET")

	apiV1.HandleFunc("/redis-data", redispkg.RedisEntrypoint).Methods("GET")

}
