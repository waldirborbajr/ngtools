package getngrokurl_test

import (
	"fmt"
	"localhost/ngtools/internal/getngrokurl"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetngrokURL(t *testing.T) {
	expected := "https://abcdef123.ngrok.io"
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"tunnels":[{"name":"command_line","uri":"/api/tunnels/command_line","public_url":"https://abcdef123.ngrok.io","proto":"https","config":{"addr":"http://localhost:80","inspect":true},"metrics":{"conns":{"count":0,"gauge":0,"rates":{"1m":0,"5m":0,"15m":0}},"http":{"count":0,"rate1":0,"rate5":0,"rate15":0},"https":{"count":0,"rate1":0,"rate5":0,"rate15":0}},"id":"some_id","public_addr":"abcdef123.ngrok.io:12345"}],"uri":"/api/tunnels"}`)
	})
	server := httptest.NewServer(handler)
	defer server.Close()

	// url, err := getngrokurl.GetngrokURL("/usr/bin/curl", "-s", server.URL+"/api/tunnels")
	url, err := getngrokurl.GetngrokURL("/usr/bin/curl")
	if err != nil {
		t.Fatalf("GetngrokURL returned error: %v", err)
	}

	if url != expected {
		t.Errorf("GetngrokURL returned %q, want %q", url, expected)
	}
}
