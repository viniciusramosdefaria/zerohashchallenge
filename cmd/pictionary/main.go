package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/viniciusramosdefaria/zerohashchallenge/pkg/metrics"
	"github.com/viniciusramosdefaria/zerohashchallenge/pkg/runner"
	"log"
	"net/http"
)

func main() {

	router := httprouter.New()
	runnerHandler := runner.NewRunnerHandler()
	metricsHandler := metrics.NewHTTPHandler()

	router.POST("/test", runnerHandler.Handler())
	router.GET("/metrics", metricsHandler.Handler())

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("static")))
	mux.Handle("/test/", router)
	mux.Handle("/metrics", router)

	log.Fatal(http.ListenAndServe(":3000", mux))


}
