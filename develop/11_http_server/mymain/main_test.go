package mymain

import (
	"config"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var app Application

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

	got := response.Body.String()
	want := "[]"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestDataToAddNewEvent_isValid(t *testing.T) {
	tests := []struct {
		name string
		d    DataToAddNewEvent
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.isValid(); got != tt.want {
				t.Errorf("DataToAddNewEvent.isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_application_AddEvent(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		app  *Application
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.app.AddEvent(tt.args.w, tt.args.r)
		})
	}
}

func TestDataToUpdateEvent_isValid(t *testing.T) {
	tests := []struct {
		name string
		d    DataToUpdateEvent
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.isValid(); got != tt.want {
				t.Errorf("DataToUpdateEvent.isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_application_UpdateEvent(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		app  *Application
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.app.UpdateEvent(tt.args.w, tt.args.r)
		})
	}
}

func Test_application_DeleteEvent(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		app  *Application
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.app.DeleteEvent(tt.args.w, tt.args.r)
		})
	}
}

func Test_application_getEventsForDay(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		app  *Application
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.app.getEventsForDay(tt.args.w, tt.args.r)
		})
	}
}

func Test_application_getEventsForWeek(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		app  *Application
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.app.getEventsForWeek(tt.args.w, tt.args.r)
		})
	}
}

func Test_application_getEventsForMonth(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		app  *Application
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.app.getEventsForMonth(tt.args.w, tt.args.r)
		})
	}
}

func Test_application_getEventsForYear(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		app  *Application
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.app.getEventsForYear(tt.args.w, tt.args.r)
		})
	}
}

func Test_sendResponse(t *testing.T) {
	type args struct {
		w    http.ResponseWriter
		code int
		data interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sendResponse(tt.args.w, tt.args.code, tt.args.data)
		})
	}
}

func Test_sendError(t *testing.T) {
	type args struct {
		w    http.ResponseWriter
		code int
		data interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sendError(tt.args.w, tt.args.code, tt.args.data)
		})
	}
}
