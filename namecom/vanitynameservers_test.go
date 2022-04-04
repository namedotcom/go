package namecom

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListVanitynameservers(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v4/domains/example.org/vanity_nameservers" {
			t.Errorf("Expected to request '/v4/domains/example.org/vanity_nameservers', got: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		resp := &ListVanityNameserversResponse{
			VanityNameservers: []*VanityNameserver{
				{DomainName: "example.org",
					Hostname: "ns1.example.org"},
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
	ListVanityNameserversResponse, err := nc.ListVanityNameservers(&ListVanityNameserversRequest{DomainName: "example.org"})
	if err != nil {
		t.Errorf("Expected 'example.org', got %s", err.Error())
	}
	if ListVanityNameserversResponse.VanityNameservers[0].DomainName != "example.org" {
		t.Errorf("Expected 'example.org', got %s", ListVanityNameserversResponse.VanityNameservers[0].DomainName)
	}

	if ListVanityNameserversResponse.VanityNameservers[0].Hostname != "ns1.example.org" {
		t.Errorf("Expected 'ns1.example.org', got %s", ListVanityNameserversResponse.VanityNameservers[0].DomainName)
	}
}

func TestGetVanitynameservers(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v4/domains/example.org/vanity_nameservers/ns1.example.org" {
			t.Errorf("Expected to request /v4/domains/example.org/vanity_nameservers/ns1.example.org', got: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		resp := &VanityNameserver{DomainName: "example.org", Hostname: "ns1.example.org"}
		jData, err := json.Marshal(resp)
		if err != nil {
			t.Error("Could not marshal response")
		}
		w.Write(jData)
	}))
	defer server.Close()

	nc := Mock("username", "apitoken", server.URL)
	GetVanityNameserverResponse, err := nc.GetVanityNameserver(&GetVanityNameserverRequest{DomainName: "example.org", Hostname: "ns1.example.org"})
	if err != nil {
		t.Errorf("Expected 'example.org', got %s", err.Error())
	}
	if GetVanityNameserverResponse.DomainName != "example.org" {
		t.Errorf("Expected 'example.org', got %s", GetVanityNameserverResponse.DomainName)
	}
}

func TestCreateVanitynameserver(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v4/domains/example.org/vanity_nameservers" {
			t.Errorf("Expected to request /v4/domains/example.org/vanity_nameservers', got: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		resp := &VanityNameserver{DomainName: "example.org", Hostname: "ns1.example.org"}
		jData, err := json.Marshal(resp)
		if err != nil {
			t.Error("Could not marshal response")
		}
		w.Write(jData)
	}))
	defer server.Close()

	nc := Mock("username", "apitoken", server.URL)
	CreateVanityNameserver, err := nc.CreateVanityNameserver(&VanityNameserver{DomainName: "example.org", Hostname: "ns1.example.org"})
	if err != nil {
		t.Errorf("Expected 'example.org', got %s", err.Error())
	}
	if CreateVanityNameserver.DomainName != "example.org" {
		t.Errorf("Expected 'example.org', got %s", CreateVanityNameserver.DomainName)
	}

}

func TestUpdateVanityNameserver(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v4/domains/example.org/vanity_nameservers/ns1.example.com" {
			t.Errorf("Expected to request /v4/domains/example.org/url/vanity_nameservers/ns1.example.com', got: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		resp := &VanityNameserver{DomainName: "example.org", Hostname: "ns1.example.com"}
		jData, err := json.Marshal(resp)
		if err != nil {
			t.Error("Could not marshal response")
		}
		w.Write(jData)
	}))
	defer server.Close()

	nc := Mock("username", "apitoken", server.URL)
	UpdateVanityNameserverResponse, err := nc.UpdateVanityNameserver(&VanityNameserver{DomainName: "example.org", Hostname: "ns1.example.com"})
	if err != nil {
		t.Errorf("Expected 'example.org', got %s", err.Error())
	}
	if UpdateVanityNameserverResponse.DomainName != "example.org" {
		t.Errorf("Expected 'example.org', got %s", UpdateVanityNameserverResponse.DomainName)
	}

	if UpdateVanityNameserverResponse.Hostname != "ns1.example.com" {
		t.Errorf("Expected 'www', got %s", UpdateVanityNameserverResponse.Hostname)
	}
}

func TestDeleteVanityNameserver(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v4/domains/example.org/vanity_nameservers/ns1.example.com" {
			t.Errorf("Expected to request /v4/domains/example.org/vanity_nameservers/ns1.example.com', got: %s", r.URL.Path)
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
	_, err := nc.DeleteVanityNameserver(&DeleteVanityNameserverRequest{DomainName: "example.org", Hostname: "ns1.example.com"})
	if err != nil {
		t.Errorf("Expected 'example.org', got %s", err.Error())
	}

}
