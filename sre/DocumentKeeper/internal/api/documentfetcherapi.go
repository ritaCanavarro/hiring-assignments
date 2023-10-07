package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/version"
	"github.com/sirupsen/logrus"
)

// -------------------- Structures ----------------------

// -------------------- Functions -----------------------

func NewHttpServer(port string) (*http.Server, error) {
	err := prometheus.DefaultRegisterer.Register(version.NewCollector("documentkeeper"))

	if err != nil {
		return nil, err
	}

	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())

	mux.HandleFunc("/-/health", func(rw http.ResponseWriter, _ *http.Request) {
		rw.WriteHeader(http.StatusOK)
		rw.Header().Set("Content-Type", "application/json")

		resp := map[string]string{
			"message": "Healthy",
		}

		jsonResp, err := json.Marshal(resp)

		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}

		rw.Write(jsonResp)
	})

	mux.HandleFunc("/-/ready", func(rw http.ResponseWriter, _ *http.Request) {
		rw.WriteHeader(http.StatusOK)
		rw.Header().Set("Content-Type", "application/json")

		resp := map[string]string{
			"message": "Ready",
		}

		jsonResp, err := json.Marshal(resp)

		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}

		rw.Write(jsonResp)
	})

	mux.Handle("/document/{id}", GetDocument).Methods("GET")

	server := &http.Server{
		Addr:     fmt.Sprintf(":%v", port),
		Handler:  mux,
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

	wg.Go(func() error {
		if err := httpServer.ListenAndServer(); err != net_http.ErrServerClosed {
			logrus.Errorf("http server error %v", err)
			return err
		}

		return nil
	})
}
