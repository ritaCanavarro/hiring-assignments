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

// -------------------- Metrics -----------------------

var sucessCountDocumentKeeper = promauto.NewCounter(
	prometheus.CounterOpts{
		Namespace: "documentKeeper",
		Name:      "document_keeper_documents_success_total",
		Help:      "Total number of documents served successfully.",
	},
)

var corruptedCountDocumentKeeper = promauto.NewCounter(
	prometheus.CounterOpts{
		Namespace: "documentKeeper",
		Name:      "document_keeper_documents_corrupted_total",
		Help:      "Total number of documents that were corrupted.",
	},
)

var errorCountDocumentKeeper = promauto.NewCounter(
	prometheus.CounterOpts{
		Namespace: "documentKeeper",
		Name:      "document_keeper_documents_error_total",
		Help:      "Total number of documents that weren't served due to errors.",
	},
)

// -------------------- Auxiliar Functions -----------------------

// processIdentifier verifies that the user
// sent a proper positive integer when
// asking for a document
func processIdentifier(url string) int {
	param := strings.Split(url, "/document/")
	id, err := strconv.Atoi(param[1])

	if err != nil || id < 0 {
		return -1
	}

	return id
}

// generateFilename takes into account the mimeType
// of the document sent by the dummy service and
// defines a proper filename with the Id provided
// by the user
func generateFilename(mimeType string, id int) string {
	if mimeType == "application/pdf" {
		return fmt.Sprintf("%d.pdf", id)
	} else {
		return fmt.Sprintf("%d.png", id)
	}
}

// validateContentType validates if the mimeType
// of the document sent by the dummy service is
// either a PDF or a PNG
func validateContentType(mimeType string, rw http.ResponseWriter) bool {
	if mimeType != "application/pdf" && mimeType != "image/png" {
		sendErrorMessage("Api served an unexpected document type.", http.StatusUnsupportedMediaType, rw)
		return false
	}

	return true
}

// sendErrorMessage takes an error message, a status code,
// and a Response writer and sends an HTTP response back
// to the user. Additionally, also increases the error metric counter.
func sendErrorMessage(errorMsg string, statusCode int, rw http.ResponseWriter) {
	errorCountDocumentKeeper.Inc()

	logrus.Errorf(errorMsg)
	auxiliar.ConfigureHttpResponse(rw, statusCode, errorMsg)
}

// fetchDocument takes the Id that the user sent, the request URL of the dummy server
// and the Response writer and fetches a random document. After reading the document
// it will ensure that it is either a PDF or a PNG and that it is not a corrupted
// document. If all is well, it will send the document back to the user.
func fetchDocument(rw http.ResponseWriter, id int, requestURL string) (string, bool) {
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
		corruptedCountDocumentKeeper.Inc()
		msg := "Api served a corrupted document."

		logrus.Errorf(msg)
		auxiliar.ConfigureHttpResponse(rw, http.StatusUnprocessableEntity, msg)
		return "", false
	}

	return filename, true
}

// -------------------- Functions -----------------------

// GetDocument processes the request of a Document by a user
// and will serve him back either a PDF or a PNG
func GetDocument(rw http.ResponseWriter, r *http.Request) {
	id := processIdentifier(r.URL.Path)
	if id == -1 {
		sendErrorMessage("Bad request because ID was not a valid positive integer.", http.StatusBadRequest, rw)
		return
	}

	serverPort, err := strconv.Atoi(os.Getenv("externalPort"))
	if err != nil {
		sendErrorMessage("Internal service Port is wrongly configured.", http.StatusInternalServerError, rw)
		return
	}

	requestURL := fmt.Sprintf("http://%s:%d", os.Getenv("externalDNS"), serverPort)

	filename, sucess := fetchDocument(rw, id, requestURL)
	if !sucess {
		return
	}

	http.ServeFile(rw, r, filename)
	sucessCountDocumentKeeper.Inc()
}
