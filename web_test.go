package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

// TestHandleHome_Success verifies that GET / returns the home page with status 200
func TestHandleHome_Success(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()

	handleHome(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status 200 OK, got %d", rr.Code)
	}

	body := rr.Body.String()
	if !strings.Contains(body, "ASCII Art Generator") {
		t.Errorf("expected body to contain logo/title, got: %s", body)
	}
}

// TestHandleHome_NotFound verifies that requesting a non-existent route returns status 404
func TestHandleHome_NotFound(t *testing.T) {
	req := httptest.NewRequest("GET", "/invalid-route-name", nil)
	rr := httptest.NewRecorder()

	handleHome(rr, req)

	if rr.Code != http.StatusNotFound {
		t.Errorf("expected status 404 Not Found, got %d", rr.Code)
	}
}

// TestHandleHome_BadRequest verifies that non-GET requests to / return status 400
func TestHandleHome_BadRequest(t *testing.T) {
	req := httptest.NewRequest("POST", "/", nil)
	rr := httptest.NewRecorder()

	handleHome(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("expected status 400 Bad Request, got %d", rr.Code)
	}
}

// TestHandleASCIIArt_Success verifies that POST /ascii-art-web with valid form fields returns status 200 and graphic representation
func TestHandleASCIIArt_Success(t *testing.T) {
	form := url.Values{}
	form.Add("text", "Go")
	form.Add("banner", "standard")

	req := httptest.NewRequest("POST", "/ascii-art-web", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rr := httptest.NewRecorder()

	handleASCIIArt(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status 200 OK, got %d", rr.Code)
	}

	body := rr.Body.String()
	// The word "Go" rendered in standard style should contain specific blocks
	if !strings.Contains(body, "____|") && !strings.Contains(body, "_____") {
		t.Errorf("expected rendered ASCII art for 'Go' in response body")
	}
}

// TestHandleASCIIArt_InvalidBanner verifies that POST /ascii-art-web with invalid banner returns 400
func TestHandleASCIIArt_InvalidBanner(t *testing.T) {
	form := url.Values{}
	form.Add("text", "Hello")
	form.Add("banner", "invalid-banner-style")

	req := httptest.NewRequest("POST", "/ascii-art-web", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rr := httptest.NewRecorder()

	handleASCIIArt(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("expected status 400 Bad Request, got %d", rr.Code)
	}
}

// TestHandleASCIIArt_InvalidCharacters verifies that POST /ascii-art-web with non-ASCII returns 400
func TestHandleASCIIArt_InvalidCharacters(t *testing.T) {
	form := url.Values{}
	form.Add("text", "hello ©")
	form.Add("banner", "standard")

	req := httptest.NewRequest("POST", "/ascii-art-web", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rr := httptest.NewRecorder()

	handleASCIIArt(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("expected status 400 Bad Request, got %d", rr.Code)
	}
}

// TestHandleASCIIArt_WrongMethod verifies that GET /ascii-art-web returns 400
func TestHandleASCIIArt_WrongMethod(t *testing.T) {
	req := httptest.NewRequest("GET", "/ascii-art-web", nil)
	rr := httptest.NewRecorder()

	handleASCIIArt(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("expected status 400 Bad Request, got %d", rr.Code)
	}
}
