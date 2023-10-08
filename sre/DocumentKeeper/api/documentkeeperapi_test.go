package api

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Requests struct {
	FilePath           string
	Id                 int
	ExpectedStatusCode int
}

func TestIdentifierDefined_AsString_ReturnsFalse(t *testing.T) {
	url := "http://127.0.0.1:4096/document/test"

	processedId := processIdentifier(url)

	assert.Less(t, processedId, 0)
}

func TestIdentifierDefined_AsNegativeNumber_ReturnsFalse(t *testing.T) {
	url := "http://127.0.0.1:4096/document/-10"

	processedId := processIdentifier(url)

	assert.Less(t, processedId, 0)
}

func TestIdentifierDefined_AsValidNumber_ReturnsTrue(t *testing.T) {
	id := rand.Int()
	url := fmt.Sprintf("http://127.0.0.1:4096/document/%d", id)

	processedId := processIdentifier(url)

	assert.NotEqual(t, -1, processedId)
	assert.Equal(t, id, processedId)
}

func TestGenerateFileName_DocumentIsPDF_ReturnsFilenameOfPDFType(t *testing.T) {
	mimeType := "application/pdf"
	id := rand.Int()

	filename := generateFilename(mimeType, id)

	assert.Contains(t, filename, ".pdf")
}

func TestGenerateFileName_DocumentIsPNG_ReturnsFilenameOfPNGType(t *testing.T) {
	mimeType := "image/png"
	id := rand.Int()

	filename := generateFilename(mimeType, id)

	assert.Contains(t, filename, ".png")
}

func TestValidateContentType_DocumentIsOfInvalidType_ReturnsError(t *testing.T) {
	mimeType := "image/jpeg"
	response := httptest.NewRecorder()

	result := validateContentType(mimeType, response)

	assert.False(t, result)
}

func TestValidateContentType_DocumentIsOfValidType_ReturnsOk(t *testing.T) {
	mimeType := "image/png"
	response := httptest.NewRecorder()

	result := validateContentType(mimeType, response)

	assert.True(t, result)
}

func TestFetchDocument_DocumentsAreValid_ReturnsOk(t *testing.T) {
	requests := make([]Requests, 2)
	requests[0] = Requests{
		"./dummyFiles/dummy.png",
		rand.Int(),
		200,
	}

	requests[1] = Requests{
		"./dummyFiles/dummy.pdf",
		rand.Int(),
		200,
	}

	for _, reqs := range requests {
		dummyServer := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			rw.WriteHeader(http.StatusOK)

			http.ServeFile(rw, r, reqs.FilePath)
		}))
		defer dummyServer.Close()

		response := httptest.NewRecorder()

		filename, sucess := fetchDocument(response, reqs.Id, dummyServer.URL)
		defer os.Remove(filename)

		assert.True(t, sucess)
		assert.Contains(t, filename, strconv.Itoa(reqs.Id))
		assert.Equal(t, reqs.ExpectedStatusCode, response.Result().StatusCode)
		assert.NotNil(t, response.Body.String())
	}
}
