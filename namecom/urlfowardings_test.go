package namecom

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListUrlfowardings(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v4/domains/example.org/url/forwarding" {
			t.Errorf("Expected to request '/v4/domains/example.org/url/forwarding', got: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		resp := &ListURLForwardingsResponse{
			URLForwarding: []*URLForwarding{
				{DomainName: "example.org"},
			},
		}
		jData, err := json.Marshal(resp)
		if err != nil {
			t.Error("Could not marshal response")
		}
		w.Write(jData)
	}))
	defer server.Close()

	nc := Mock("username", "apitoken", server.URL)
	ListURLForwardingsResponse, err := nc.ListURLForwardings(&ListURLForwardingsRequest{DomainName: "example.org"})
	if err != nil {
		t.Errorf("Expected 'example.org', got %s", err.Error())
	}
	if ListURLForwardingsResponse.URLForwarding[0].DomainName != "example.org" {
		t.Errorf("Expected 'example.org', got %s", ListURLForwardingsResponse.URLForwarding[0].DomainName)
	}
}

func TestGetUrlforwarding(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v4/domains/example.org/url/forwarding/www" {
			t.Errorf("Expected to request /v4/domains/example.org/url/forwarding/www', got: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		resp := &URLForwarding{DomainName: "example.org", Host: "www"}
		jData, err := json.Marshal(resp)
		if err != nil {
			t.Error("Could not marshal response")
		}
		w.Write(jData)
	}))
	defer server.Close()

	nc := Mock("username", "apitoken", server.URL)
	GetURLForwardingResponse, err := nc.GetURLForwarding(&GetURLForwardingRequest{DomainName: "example.org", Host: "www"})
	if err != nil {
		t.Errorf("Expected 'example.org', got %s", err.Error())
	}
	if GetURLForwardingResponse.DomainName != "example.org" {
		t.Errorf("Expected 'example.org', got %s", GetURLForwardingResponse.DomainName)
	}
}

func TestCreateUrlForwarding(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v4/domains/example.org/url/forwarding" {
			t.Errorf("Expected to request /v4/domains/example.org/url/forwarding', got: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		resp := &Record{DomainName: "example.org"}
		jData, err := json.Marshal(resp)
		if err != nil {
			t.Error("Could not marshal response")
		}
		w.Write(jData)
	}))
	defer server.Close()

	nc := Mock("username", "apitoken", server.URL)
	CreateURLForwardingResponse, err := nc.CreateURLForwarding(&URLForwarding{DomainName: "example.org"})
	if err != nil {
		t.Errorf("Expected 'example.org', got %s", err.Error())
	}
	if CreateURLForwardingResponse.DomainName != "example.org" {
		t.Errorf("Expected 'example.org', got %s", CreateURLForwardingResponse.DomainName)
	}

}

func TestUpdateUrlforwarding(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v4/domains/example.org/url/forwarding/www" {
			t.Errorf("Expected to request /v4/domains/example.org/url/forwarding/www', got: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		resp := &Record{DomainName: "example.org", Host: "www"}
		jData, err := json.Marshal(resp)
		if err != nil {
			t.Error("Could not marshal response")
		}
		w.Write(jData)
	}))
	defer server.Close()

	nc := Mock("username", "apitoken", server.URL)
	UpdateUrlforwardingResponse, err := nc.UpdateURLForwarding(&URLForwarding{DomainName: "example.org", Host: "www"})
	if err != nil {
		t.Errorf("Expected 'example.org', got %s", err.Error())
	}
	if UpdateUrlforwardingResponse.DomainName != "example.org" {
		t.Errorf("Expected 'example.org', got %s", UpdateUrlforwardingResponse.DomainName)
	}

	if UpdateUrlforwardingResponse.Host != "www" {
		t.Errorf("Expected 'www', got %s", UpdateUrlforwardingResponse.Host)
	}
}

func TestDeleteUrlforwarding(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v4/domains/example.org/url/forwarding/www" {
			t.Errorf("Expected to request /v4/domains/example.org/url/forwarding/www', got: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		resp := EmptyResponse{}
		jData, err := json.Marshal(resp)
		if err != nil {
			t.Error("Could not marshal response")
		}
		w.Write(jData)
	}))
	defer server.Close()

	nc := Mock("username", "apitoken", server.URL)
	_, err := nc.DeleteURLForwarding(&DeleteURLForwardingRequest{DomainName: "example.org", Host: "www"})
	if err != nil {
		t.Errorf("Expected 'example.org', got %s", err.Error())
	}

}
