package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIndexHandler(t *testing.T) {
	mockUserResp := `{"message":"hello world"}`

	ts := httptest.NewServer(SetupServer())
	defer ts.Close()

	res, err := http.Get(fmt.Sprintf("%s/", ts.URL))
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)

	assert.Nil(t, err)
	// check; did we get the right http status code
	assert.Equal(t, http.StatusOK, res.StatusCode)

	responseData, _ := io.ReadAll(res.Body)

	// check; did we get the right http response body
	assert.Equal(t, mockUserResp, string(responseData))
}
