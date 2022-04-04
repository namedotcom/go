package namecom

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListDnssecs(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v4/domains/example.org/dnssec" {
			t.Errorf("Expected to request '/v4/domains/example.org/dnssec', got: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		resp := &ListDNSSECsResponse{
			Dnssec: []*DNSSEC{
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
	ListDnssecsResponse, err := nc.ListDNSSECs(&ListDNSSECsRequest{DomainName: "example.org"})
	if err != nil {
		t.Errorf("Expected 'example.org', got %s", err.Error())
	}
	if ListDnssecsResponse.Dnssec[0].DomainName != "example.org" {
		t.Errorf("Expected 'example.org', got %s", ListDnssecsResponse.Dnssec[0].DomainName)
	}
}

func TestGetDnssec(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v4/domains/example.org/dnssec/E2D3C916F6DEEAC73294E8268FB5885044A833FC5459588F4A9184CFC41A5766" {
			t.Errorf("Expected to request '/v4/domains/example.org/dnssec/E2D3C916F6DEEAC73294E8268FB5885044A833FC5459588F4A9184CFC41A5766', got: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		resp := &DNSSEC{DomainName: "example.org", Digest: "E2D3C916F6DEEAC73294E8268FB5885044A833FC5459588F4A9184CFC41A5766"}
		jData, err := json.Marshal(resp)
		if err != nil {
			t.Error("Could not marshal response")
		}
		w.Write(jData)
	}))
	defer server.Close()

	nc := Mock("username", "apitoken", server.URL)
	GetDNSSECResponse, err := nc.GetDNSSEC(&GetDNSSECRequest{DomainName: "example.org", Digest: "E2D3C916F6DEEAC73294E8268FB5885044A833FC5459588F4A9184CFC41A5766"})
	if err != nil {
		t.Errorf("Expected 'example.org', got %s", err.Error())
	}
	if GetDNSSECResponse.DomainName != "example.org" {
		t.Errorf("Expected 'example.org', got %s", GetDNSSECResponse.DomainName)
	}
	if GetDNSSECResponse.Digest != "E2D3C916F6DEEAC73294E8268FB5885044A833FC5459588F4A9184CFC41A5766" {
		t.Errorf("Expected 'E2D3C916F6DEEAC73294E8268FB5885044A833FC5459588F4A9184CFC41A5766', got %s", GetDNSSECResponse.Digest)
	}
}

func TestCreateDnssec(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v4/domains/example.org/dnssec" {
			t.Errorf("Expected to request '/v4/domains/example.org/dnssec', got: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		resp := &DNSSEC{DomainName: "example.org"}
		jData, err := json.Marshal(resp)
		if err != nil {
			t.Error("Could not marshal response")
		}
		w.Write(jData)
	}))
	defer server.Close()

	nc := Mock("username", "apitoken", server.URL)
	CreateDNSSECResponse, err := nc.CreateDNSSEC(&DNSSEC{DomainName: "example.org"})
	if err != nil {
		t.Errorf("Expected 'example.org', got %s", err.Error())
	}
	if CreateDNSSECResponse.DomainName != "example.org" {
		t.Errorf("Expected 'example.org', got %s", CreateDNSSECResponse.DomainName)
	}
}

func TestDeleteDnssec(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v4/domains/example.org/dnssec/E2D3C916F6DEEAC73294E8268FB5885044A833FC5459588F4A9184CFC41A5766" {
			t.Errorf("Expected to request '/v4/domains/example.org/dnssec/E2D3C916F6DEEAC73294E8268FB5885044A833FC5459588F4A9184CFC41A5766', got: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		resp := &DNSSEC{DomainName: "example.org"}
		jData, err := json.Marshal(resp)
		if err != nil {
			t.Error("Could not marshal response")
		}
		w.Write(jData)
	}))
	defer server.Close()

	nc := Mock("username", "apitoken", server.URL)
	_, err := nc.DeleteDNSSEC(&DeleteDNSSECRequest{DomainName: "example.org", Digest: "E2D3C916F6DEEAC73294E8268FB5885044A833FC5459588F4A9184CFC41A5766"})
	if err != nil {
		t.Errorf("Expected 'example.org', got %s", err.Error())
	}
}
