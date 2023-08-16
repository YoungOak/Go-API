package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRouter_AddHandler(t *testing.T) {
	r := NewRouter("")
	handlerCalled := false
	testHandler := func(w http.ResponseWriter, r *http.Request) {
		handlerCalled = true
	}
	r.AddHandler("/test", testHandler)

	rType, _ := r.(*router)
	server := httptest.NewServer(rType.router)
	defer server.Close()

	http.Get(server.URL + "/test")

	if !handlerCalled {
		t.Fatal("expected testHandler to be called")
	}
}

// Optional: You can potentially skip this test, as it's more of an
// assertion about the Go standard library's functionality than your own code.
func TestRouter_Serve(t *testing.T) {
	r := NewRouter(":8080")
	rType, _ := r.(*router)
	server := httptest.NewServer(rType.router)
	defer server.Close()

	resp, err := http.Get(server.URL)
	if err != nil {
		t.Fatalf("could not make GET request: %v", err)
	}

	if resp.StatusCode != http.StatusNotFound {
		t.Fatalf("expected status 404, got: %d", resp.StatusCode)
	}
}
