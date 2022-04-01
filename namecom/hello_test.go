package namecom

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHello(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v4/hello" {
			t.Errorf("Expected to request '/v4/hello', got: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		resp := &HelloResponse{
			ServerName: "api01",
			Username:   "username",
		}
		jData, err := json.Marshal(resp)
		if err != nil {
			t.Error("Could not marshal response")
		}
		w.Write(jData)
	}))
	defer server.Close()

	nc := Mock("username", "apitoken", server.URL)
	HelloFuncResponse, err := nc.HelloFunc(&HelloRequest{})
	if err != nil {
		t.Errorf("Expected 'example.org', got %s", err.Error())
	}
	if HelloFuncResponse.ServerName != "api01" {
		t.Errorf("Expected 'api01', got %s", HelloFuncResponse.ServerName)
	}

	if HelloFuncResponse.Username != "username" {
		t.Errorf("Expected 'username', got %s", HelloFuncResponse.ServerName)
	}
}
