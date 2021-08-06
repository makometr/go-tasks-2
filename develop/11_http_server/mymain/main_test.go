package mymain

import (
	"bytes"
	"config"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"
)

var app Application

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func TestMain(m *testing.M) {
	cfg := config.NewTestConfig()
	app.initDateStorage(cfg)
	app.initHTTPServer(cfg)
	defer app.Shutdown()

	// Run tests
	exitVal := m.Run()

	// Write code here to run after tests

	// Exit with exit value from tests
	os.Exit(exitVal)
}

func Test_application_getEvent(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "/event", nil)
	response := httptest.NewRecorder()

	app.getEvent(response, request)
	checkResponseCode(t, http.StatusOK, response.Code)

	want := bytes.TrimSuffix(response.Body.Bytes(), []byte{10})
	expected := []byte(`{"result":[]}`)

	if !reflect.DeepEqual(expected, want) {
		t.Errorf("Expected %v array. Got %v", expected, want)
	}
}
