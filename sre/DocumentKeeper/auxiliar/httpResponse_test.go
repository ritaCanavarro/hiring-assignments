package auxiliar

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigureHttpResponse_WhenNonEmptyMsg_ReturnsCorrestResponse(t *testing.T) {
	response := httptest.NewRecorder()
	statusCode := http.StatusOK
	msg := "Test was a success"

	ConfigureHttpResponse(response, statusCode, msg)

	assert.Equal(t, statusCode, response.Result().StatusCode)
	assert.Equal(t, "{\"message\":\"Test was a success\"}", response.Body.String())
}

func TestConfigureHttpResponse_WhenEmptyMsg_ReturnsCorrestResponse(t *testing.T) {
	response := httptest.NewRecorder()
	statusCode := http.StatusOK
	msg := ""

	ConfigureHttpResponse(response, statusCode, msg)

	assert.Equal(t, statusCode, response.Result().StatusCode)
	assert.Equal(t, "OK", response.Body.String())
}
