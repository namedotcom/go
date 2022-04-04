package namecom

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListEmailForwardings(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v4/domains/example.org/email/forwarding" {
			t.Errorf("Expected to request '/v4/domains/example.org/email/forwarding', got: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		resp := &ListEmailForwardingsResponse{
			EmailForwarding: []*EmailForwarding{
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
	ListEmailForwardingsResponse, err := nc.ListEmailForwardings(&ListEmailForwardingsRequest{DomainName: "example.org"})
	if err != nil {
		t.Errorf("Expected 'example.org', got %s", err.Error())
	}
	if ListEmailForwardingsResponse.EmailForwarding[0].DomainName != "example.org" {
		t.Errorf("Expected 'example.org', got %s", ListEmailForwardingsResponse.EmailForwarding[0].DomainName)
	}
}

func TestGetEmailForwarding(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v4/domains/example.org/email/forwarding/admin" {
			t.Errorf("Expected to request '/v4/domains/example.org/email/forwarding/admin', got: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		resp := &EmailForwarding{DomainName: "example.org", EmailBox: "admin"}
		jData, err := json.Marshal(resp)
		if err != nil {
			t.Error("Could not marshal response")
		}
		w.Write(jData)
	}))
	defer server.Close()

	nc := Mock("username", "apitoken", server.URL)
	GetEmailForwardingRequest, err := nc.GetEmailForwarding(&GetEmailForwardingRequest{DomainName: "example.org", EmailBox: "admin"})
	if err != nil {
		t.Errorf("Expected 'example.org', got %s", err.Error())
	}
	if GetEmailForwardingRequest.DomainName != "example.org" {
		t.Errorf("Expected 'example.org', got %s", GetEmailForwardingRequest.DomainName)
	}
	if GetEmailForwardingRequest.EmailBox != "admin" {
		t.Errorf("Expected 'admin', got %s", GetEmailForwardingRequest.EmailBox)
	}
}

func TestCreateEmailForwarding(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v4/domains/example.org/email/forwarding" {
			t.Errorf("Expected to request '/v4/domains/example.org/email/forwarding', got: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		resp := &EmailForwarding{DomainName: "example.org", EmailBox: "admin", EmailTo: "webmaster@example.net"}
		jData, err := json.Marshal(resp)
		if err != nil {
			t.Error("Could not marshal response")
		}
		w.Write(jData)
	}))
	defer server.Close()

	nc := Mock("username", "apitoken", server.URL)
	CreateEmailForwardingResponse, err := nc.CreateEmailForwarding(&EmailForwarding{DomainName: "example.org", EmailBox: "admin", EmailTo: "webmaster@example.net"})
	if err != nil {
		t.Errorf("Expected 'example.org', got %s", err.Error())
	}
	if CreateEmailForwardingResponse.DomainName != "example.org" {
		t.Errorf("Expected 'example.org', got %s", CreateEmailForwardingResponse.DomainName)
	}

	if CreateEmailForwardingResponse.EmailBox != "admin" {
		t.Errorf("Expected 'admin', got %s", CreateEmailForwardingResponse.EmailBox)
	}

	if CreateEmailForwardingResponse.EmailTo != "webmaster@example.net" {
		t.Errorf("Expected 'webmaster@example.net', got %s", CreateEmailForwardingResponse.EmailTo)
	}
}

func TestDeleteEmailForwarding(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v4/domains/example.org/email/forwarding/admin" {
			t.Errorf("Expected to request '/v4/domains/example.org/email/forwarding/admin', got: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		resp := &EmailForwarding{DomainName: "example.org", EmailBox: "admin"}
		jData, err := json.Marshal(resp)
		if err != nil {
			t.Error("Could not marshal response")
		}
		w.Write(jData)
	}))
	defer server.Close()

	nc := Mock("username", "apitoken", server.URL)
	_, err := nc.DeleteEmailForwarding(&DeleteEmailForwardingRequest{DomainName: "example.org", EmailBox: "admin"})
	if err != nil {
		t.Errorf("Expected 'example.org', got %s", err.Error())
	}
}
