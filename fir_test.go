package firGo

import "testing"

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
