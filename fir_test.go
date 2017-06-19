package firGo

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewFIRClient(t *testing.T) {
	tok := "test-token"
	client := NewFIRClient("http://foobar", tok)
	want := "http://foobar/api"
	if client.BaseURL.String() != want {
		t.Errorf("BaseURL did not match. Wanted %s, got %s", want, client.BaseURL.String())
	}

	expectedToken := "Token test-token"
	if client.Token != expectedToken {
		t.Errorf("Token did not match expected. Wanted %s, got %s", expectedToken, client.Token)
	}
}

func TestNewRequest(t *testing.T) {
	// Setup test GET request
	client := NewFIRClient("http://foobar", "test")
	c, err := client.NewRequest("GET", "/api", nil)
	if err != nil {
		t.Error(err)
	}

	if c.Header.Get("Authorization") != "Token test" {
		t.Error(errors.New("authorization header incorrect"))
	}

	// Setup test POST request
	a := ArtifactRequest{Value: "TEST"}
	c, err = client.NewRequest("GET", "/api", a)
	if err != nil {
		t.Error(err)
	}
}

func TestDo(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{"fake json string"}`)
	}))
	defer ts.Close()

	// Setup request
	client := NewFIRClient(ts.URL, "test")
	req, _ := client.NewRequest("GET", "/", nil)
	res, err := client.Do(req)
	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 200 {
		t.Error("Bad status")
	}
}
