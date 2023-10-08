package api

import (
	"DocumentKeeper/auxiliar"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

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

func processIdentifier(url string) int {
	param := strings.Split(url, "/document/")
	id, err := strconv.Atoi(param[1])

	if err != nil || id < 0 {
		return -1
	}

	return id
}

func generateFilename(mimeType string, id int) string {
	if mimeType == "application/pdf" {
		return fmt.Sprintf("%d.pdf", id)
	} else {
		return fmt.Sprintf("%d.png", id)
	}
}

func validateContentType(mimeType string, rw http.ResponseWriter) bool {
	if mimeType != "application/pdf" && mimeType != "image/png" {
		sendErrorMessage("Api served an unexpected document type.", http.StatusUnsupportedMediaType, rw)
		return false
	}

	return true
}

func sendErrorMessage(errorMsg string, statusCode int, rw http.ResponseWriter) {
	errorCountDocumentFetcher.Inc()

	logrus.Errorf(errorMsg)
	auxiliar.ConfigureHttpResponse(rw, statusCode, errorMsg)
}

func fetchDocument(rw http.ResponseWriter, id int) (string, bool) {
	requestURL := fmt.Sprintf("http://localhost:%d", serverPort)
	resp, err := http.Get(requestURL)
	defaultErrorMsg := fmt.Sprintf("Failed to fetch document from %s", requestURL)

	if err != nil {
		sendErrorMessage(defaultErrorMsg, http.StatusFailedDependency, rw)
		return "", false
	}

	body, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		sendErrorMessage(defaultErrorMsg, http.StatusUnprocessableEntity, rw)
		return "", false
	}

	mimeType := http.DetectContentType(body)
	valideMimeType := validateContentType(mimeType, rw)

	if !valideMimeType {
		return "", false
	}

	var filename = generateFilename(mimeType, id)
	err = os.WriteFile(filename, body, 0644)

	if err != nil {
		corruptedCountDocumentFetcher.Inc()
		msg := "Api served a corrupted document."

		logrus.Errorf(msg)
		auxiliar.ConfigureHttpResponse(rw, http.StatusUnprocessableEntity, msg)
		return "", false
	}

	return filename, true
}

// -------------------- Functions -----------------------

func GetDocument(rw http.ResponseWriter, r *http.Request) {
	id := processIdentifier(r.URL.Path)
	if id == -1 {
		sendErrorMessage("Bad request because ID was not a valid positive integer.", http.StatusBadRequest, rw)
		return
	}

	filename, sucess := fetchDocument(rw, id)
	if !sucess {
		return
	}

	http.ServeFile(rw, r, filename)
	sucessCountDocumentFetcher.Inc()
}
