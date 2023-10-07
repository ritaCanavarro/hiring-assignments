package http

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"DocumentKeeper/internal/api"
	"DocumentKeeper/internal/auxiliar"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/version"
	"github.com/sirupsen/logrus"
)

// -------------------- Global Vars & Const -----------------------

var router = mux.NewRouter()

// -------------------- Functions -----------------------
func NewHttpServer(port string) (*http.Server, error) {
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

	router.HandleFunc("/document/{id}", api.GetDocument).Methods("GET")

	server := &http.Server{
		Addr:     fmt.Sprintf(":%v", port),
		Handler:  router,
		ErrorLog: &log.Logger{},
	}

	return server, nil
}

func StartDocumentFetcher(httpPort string) {
	httpServer, err := NewHttpServer(httpPort)

	if err != nil {
		logrus.Errorf("http server error %v", err)
		os.Exit(1)
	}

	go func() {
		if err := httpServer.ListenAndServe(); err != nil {
			logrus.Errorf("http server error %v", err)
		}
	}()
}
