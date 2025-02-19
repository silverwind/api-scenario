package model_test

import (
	"testing"
	"time"

	"github.com/sendgrid/rest"
	"github.com/silverwind/api-scenario/pkg/model"
	"github.com/silverwind/api-scenario/test"
)

func TestNewResponseCreateAValidResponse(t *testing.T) {
	restResp := rest.Response{
		StatusCode: 200,
		Body:       `{"hello":"world"}`,
		Headers: map[string][]string{
			"Accept": {"toto"},
		},
	}

	expectedDuration := time.Duration(1 * time.Second)
	response, err := model.NewResponse(restResp, expectedDuration)
	test.Ok(t, err)
	test.Equals(t, "Invalid duration", expectedDuration, response.TimeElapsed)
	test.Equals(t, "Invalid status code", restResp.StatusCode, response.StatusCode)
	expectedBody := `{"hello":"world"}`
	test.Equals(t, "Invalid body", expectedBody, response.Body)
}

func TestNewResponseEmptyBody(t *testing.T) {
	restResp := rest.Response{
		StatusCode: 200,
		Body:       "",
		Headers: map[string][]string{
			"Accept": {"toto"},
		},
	}

	expectedDuration := time.Duration(1 * time.Second)
	response, err := model.NewResponse(restResp, expectedDuration)
	test.Ok(t, err)
	test.Equals(t, "Invalid duration", expectedDuration, response.TimeElapsed)
	test.Equals(t, "Invalid status code", restResp.StatusCode, response.StatusCode)
	test.Equals(t, "Body should be an empty map", 0, len(response.Body))
}
