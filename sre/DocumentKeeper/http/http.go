package http

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"DocumentKeeper/api"
	"DocumentKeeper/auxiliar"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/version"
	"github.com/sirupsen/logrus"
)

// -------------------- Functions -----------------------

// NewHttpServer will starts an HTTP Server
// which accepts requests in the following handlers:
// /metrics - So Prometheus can scrape this application metrics
// /-/health - For the liveliness probing
// /-/ready - For the readiness probing
// /document/{id} - For the user to retrieve a random document
func NewHttpServer(port int) (*http.Server, error) {
	router := mux.NewRouter()
	err := prometheus.DefaultRegisterer.Register(version.NewCollector("documentkeeper"))

	if err != nil {
		return nil, err
	}

	router.Handle("/metrics", promhttp.Handler())

	router.HandleFunc("/-/health", func(rw http.ResponseWriter, _ *http.Request) {
		auxiliar.ConfigureHttpResponse(rw, http.StatusOK, "Healthy")
	})

	router.HandleFunc("/-/ready", func(rw http.ResponseWriter, _ *http.Request) {
		auxiliar.ConfigureHttpResponse(rw, http.StatusOK, "Ready")
	})

	router.HandleFunc("/ping", func(rw http.ResponseWriter, _ *http.Request) {
		auxiliar.ConfigureHttpResponse(rw, http.StatusOK, "Ready")
	})

	router.HandleFunc("/document/{id}", api.GetDocument).Methods("GET")

	server := &http.Server{
		Addr:     fmt.Sprintf(":%d", port),
		Handler:  router,
		ErrorLog: &log.Logger{},
	}

	return server, nil
}

// StartDocumentFetcher will define an HTTP Server
// And will be listening for requests in the Port
// defined in an Environment Variable, which in
// turn is defined in the Deployment Helm Chart
func StartDocumentFetcher() {
	httpPort, err := strconv.Atoi(os.Getenv("internalPort"))
	if err != nil {
		logrus.Errorf("Error fetching httpPort %v", err)
		os.Exit(1)
	}

	httpServer, err := NewHttpServer(httpPort)

	if err != nil {
		logrus.Errorf("http server error %v", err)
		os.Exit(1)
	}

	if err := httpServer.ListenAndServe(); err != nil {
		logrus.Errorf("http server error %v", err)
	}
}
