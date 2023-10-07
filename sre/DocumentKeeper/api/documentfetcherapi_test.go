package api

import (
	"math/rand"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIdentifierDefined_AsString_ReturnsFalse(t *testing.T) {
	params := make(map[string]string)
	params["id"] = "identifier"

	processedId := processIdentifier(params)

	assert.Less(t, processedId, 0)
}

func TestIdentifierDefined_AsNegativeNumber_ReturnsFalse(t *testing.T) {
	params := make(map[string]string)
	params["id"] = "-1"

	processedId := processIdentifier(params)

	assert.Less(t, processedId, 0)
}

func TestIdentifierDefined_AsValidNumber_ReturnsTrue(t *testing.T) {
	params := make(map[string]string)
	id := rand.Int()
	params["id"] = strconv.Itoa(id)

	processedId := processIdentifier(params)

	assert.NotEqual(t, processedId, -1)
	assert.Equal(t, processedId, id)
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
