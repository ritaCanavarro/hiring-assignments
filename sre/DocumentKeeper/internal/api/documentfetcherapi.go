package api

import (
	"DocumentKeeper/internal/auxiliar"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/sirupsen/logrus"
)

// -------------------- Global Vars & Const -----------------------

const serverPort = 3000

// -------------------- Metrics -----------------------

var sucessCountDocumentFetcher = promauto.NewCounter(
	prometheus.CounterOpts{
		Namespace: "documentKeeper",
		Name:      "document_fetcher_documents_success_total",
		Help:      "Total number of documents served successfully.",
	},
)

var corruptedCountDocumentFetcher = promauto.NewCounter(
	prometheus.CounterOpts{
		Namespace: "documentKeeper",
		Name:      "document_fetcher_documents_corrupted_total",
		Help:      "Total number of documents that were corrupted.",
	},
)

var errorCountDocumentFetcher = promauto.NewCounter(
	prometheus.CounterOpts{
		Namespace: "documentKeeper",
		Name:      "document_fetcher_documents_error_total",
		Help:      "Total number of documents that weren't served due to errors.",
	},
)

// -------------------- Auxiliar Functions -----------------------

func processIdentifier(params map[string]string) int {
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		errorCountDocumentFetcher.Inc()
		logrus.Errorf("Identifier is not a valid integer %v.", err)
		return -1
	}

	if id < 0 {
		errorCountDocumentFetcher.Inc()
		logrus.Errorf("Identifier is not a positive integer.")
		return -1
	}

	return id
}

// -------------------- Functions -----------------------

func GetDocument(rw http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := processIdentifier(params)

	if id == -1 {
		auxiliar.ConfigureHttpResponse(rw, http.StatusBadRequest, "Bad request because ID was not a valid positive integer.")
		return
	}

	requestURL := fmt.Sprintf("http://localhost:%d", serverPort)
	resp, err := http.Get(requestURL)

	if err != nil {
		errorCountDocumentFetcher.Inc()

		logrus.Errorf("Failed to fetch document from %s", requestURL)
		auxiliar.ConfigureHttpResponse(rw, http.StatusFailedDependency, "Failed to fetch document from API.")
		return
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		errorCountDocumentFetcher.Inc()

		logrus.Errorf("Failed to fetch document from %s", requestURL)
		auxiliar.ConfigureHttpResponse(rw, http.StatusUnprocessableEntity, "Failed to read document from API.")
		return
	}

	mimeType := http.DetectContentType(body)

	if mimeType != "application/pdf" && mimeType != "image/png" {
		errorCountDocumentFetcher.Inc()

		logrus.Errorf("Api served an unexpected document type.")
		auxiliar.ConfigureHttpResponse(rw, http.StatusUnsupportedMediaType, "Document is not supported.")
		return
	}

	var filename = ""

	if mimeType == "application/pdf" {
		filename = fmt.Sprintf("%d.pdf", id)
	} else {
		filename = fmt.Sprintf("%d.png", id)
	}

	err = os.WriteFile(filename, body, 0644)

	if err != nil {
		corruptedCountDocumentFetcher.Inc()

		logrus.Errorf("Api served a corrupted document.")
		auxiliar.ConfigureHttpResponse(rw, http.StatusUnprocessableEntity, "Document is corrupted.")
		return
	}

	http.ServeFile(rw, r, filename)
	sucessCountDocumentFetcher.Inc()
}
