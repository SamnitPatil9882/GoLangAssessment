package service_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"serviceTesting"
)

func TestWebsiteStatuses(t *testing.T) {
	request, err := http.NewRequest("GET", "/websites", nil)
	if err != nil {
		t.Fatal(err)
		return
	}

	response := httptest.NewRecorder()
	handler := http.HandlerFunc(service.WebsiteStatuses)

	handler.ServeHTTP(response, request)

	if status := response.Code; status != http.StatusOK {
		t.Errorf("error !! got error due to : want %v, got %v", status, http.StatusOK)
	}

}

func TestWebsiteHandler(t *testing.T) {
	jsonData := []byte(`{"websites":["https://pkg.go.dev/github.com/gorilla/mux","https://www.youtube.com","https://www.facebook.com"]
	}
	`)
	request, err := http.NewRequest("POST", "websites", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatal("error !! while posting data !!")
		return
	}
	response := httptest.NewRecorder()
	handler := http.HandlerFunc(service.WebsiteHandler)

	handler.ServeHTTP(response, request)

	if status := response.Code; status != http.StatusOK {
		t.Fatal("error !! Error in POST /websites")
	}

	//For Empty List
	expectedOutput := "Empty list of websites provided."
	request, err = http.NewRequest("POST", "websites", bytes.NewBuffer([]byte("[]")))
	if err != nil {
		t.Fatal("error !! while posting data !!")
		return
	}
	response = httptest.NewRecorder()
	handler = http.HandlerFunc(service.WebsiteHandler)

	handler.ServeHTTP(response, request)

	if body := response.Body.String(); body != "" {
		t.Errorf("Expected body '%s', but got '%s'", expectedOutput, body)
	}

	//For Invalid URLs
	expectedOutput = "Invalid URL: www.facebook.com"
	request, err = http.NewRequest("POST", "websites", bytes.NewBuffer([]byte(`{
		"websites":["www.facebook.com"]
	  }`)))
	if err != nil {
		t.Fatal("error !! while posting data !!")
		return
	}
	response = httptest.NewRecorder()
	handler = http.HandlerFunc(service.WebsiteHandler)

	handler.ServeHTTP(response, request)

	if body := response.Body.String(); body != "Invalid URL: www.facebook.com" {
		t.Errorf("Expected body '%s', but got '%s'", expectedOutput, body)
	}

}

func TestShowStatus(t *testing.T) {
	//For Valid Case
	req, err := http.NewRequest("GET", "/website?name=https://www.youtube.com/", nil)
	if err != nil {
		t.Fatal(err)
	}

	response := httptest.NewRecorder()

	service.ShowStatus(response, req)

	if response.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, but got %d", http.StatusBadRequest, response.Code)
	}

	expectedBody := "https://www.youtube.com/ is now Down !!"
	if body := response.Body.String(); body != expectedBody {
		t.Errorf("Expected body '%s', but got '%s'", expectedBody, body)
	}

	//for Invalid Case
	req, err = http.NewRequest("GET", "/website?name=https://www.invalid.com", nil)
	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()
	service.ShowStatus(res, req)

	if res.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, but got %d", http.StatusBadRequest, res.Code)
	}

	expectedBody = "https://www.invalid.com is now Down !!"
	if body := res.Body.String(); body != expectedBody {
		t.Errorf("Expected body '%s', but got '%s'", expectedBody, body)
	}
}