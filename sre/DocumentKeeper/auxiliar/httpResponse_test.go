package auxiliar

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Requests struct {
	Message         string
	StatusCode      int
	ExpectedMessage string
}

func TestConfigureHttpResponse_WhenMsgReceived_ReturnsCorrectResponse(t *testing.T) {
	requests := make([]Requests, 2)
	requests[0] = Requests{
		"Test was a success",
		200,
		"{\"message\":\"Test was a success\"}",
	}

	requests[1] = Requests{
		"",
		200,
		"OK",
	}

	for _, request := range requests {
		response := httptest.NewRecorder()

		ConfigureHttpResponse(response, request.StatusCode, request.Message)

		assert.Equal(t, request.StatusCode, response.Result().StatusCode)
		assert.Equal(t, request.ExpectedMessage, response.Body.String())
	}
}
