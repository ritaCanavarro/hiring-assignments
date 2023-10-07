package api

import (
	"fmt"
	"math/rand"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Requests struct {
	Url                string
	Id                 int
	ExpectedStatusCode int
	ExpectedResponse   string
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

func TestGetDocument_DocumentIsOfInvalidType_ReturnsError(t *testing.T) {
	mimeType := "image/jpeg"
	response := httptest.NewRecorder()

	result := validateContentType(mimeType, response)

	assert.False(t, result)
}

func TestGetDocument_DocumentIsOfValidType_ReturnsOk(t *testing.T) {
	mimeType := "image/png"
	response := httptest.NewRecorder()

	result := validateContentType(mimeType, response)

	assert.True(t, result)
}

func TestGetDocument_DocumentIsPNG_ReturnsPNG(t *testing.T) {
	// server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
	// 	http.ServeFile(rw, r, "./dummyFiles/golang.png")
	// }))
	// defer server.Close()

	// response := httptest.NewRecorder()
	// id := rand.Int()
	// url := fmt.Sprintf("http://127.0.0.1:4096/document/%d", id)
	// request, _ := http.NewRequest(http.MethodGet, url, nil)

	// GetDocument(response, request)

	// assert.Equal(t, http.StatusUnsupportedMediaType, response.Result().StatusCode)
	// assert.Equal(t, "Api served an unexpected document type.", response.Body.String())
}

func TestGetDocument_DocumentIsPDF_ReturnsPDF(t *testing.T) {

}
