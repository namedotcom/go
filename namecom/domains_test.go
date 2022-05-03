package namecom

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListDomain(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v4/domains" {
			t.Errorf("Expected to request '/v4/domains', got: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		resp := &ListDomainsResponse{
			Domains: []*Domain{
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
	ListDomainsResponse, err := nc.ListDomains(&ListDomainsRequest{})
	if err != nil {
		t.Errorf("Expected 'example.org', got %s", err.Error())
	}
	if ListDomainsResponse.Domains[0].DomainName != "example.org" {
		t.Errorf("Expected 'example.org', got %s", ListDomainsResponse.Domains[0].DomainName)
	}
}

func TestGetDomain(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v4/domains/example.org" {
			t.Errorf("Expected to request '/v4/domains/example.org', got: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		resp := &Domain{DomainName: "example.org"}
		jData, err := json.Marshal(resp)
		if err != nil {
			t.Error("Could not marshal response")
		}
		w.Write(jData)
	}))
	defer server.Close()

	nc := Mock("username", "apitoken", server.URL)
	Domain, err := nc.GetDomain(&GetDomainRequest{DomainName: "example.org"})
	if err != nil {
		t.Errorf("Expected 'example.org', got %s", err.Error())
	}
	if Domain.DomainName != "example.org" {
		t.Errorf("Expected 'example.org', got %s", Domain.DomainName)
	}
}

func TestCreateForDomain(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v4/domains" {
			t.Errorf("Expected to request '/v4/domains', got: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		//not sure why this has to be done like this
		resp := &CreateDomainResponse{Domain: &Domain{DomainName: "example.org"}}
		//	resp := &Domain{DomainName: "example.org"}
		jData, err := json.Marshal(resp)
		if err != nil {
			t.Error("Could not marshal response")
		}
		w.Write(jData)
	}))
	defer server.Close()
	//&RenewDomainResponse{Domain: &Domain{DomainName: "example.org"}}
	nc := Mock("username", "apitoken", server.URL)

	DomainCreateResponse, err := nc.CreateDomain(&CreateDomainRequest{Domain: &Domain{DomainName: "example.org"}})
	if err != nil {
		t.Errorf("Expected 'example.org', got %s", err.Error())
	}
	if DomainCreateResponse.Domain.DomainName != "example.org" {
		t.Errorf("Expected 'example.org', got %s", DomainCreateResponse.Domain.DomainName)
	}
}

func TestEnableAutorenew(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v4/domains/example.org:enableAutorenew" {
			t.Errorf("Expected to request '/v4/domains/example.org:enableAutorenew', got: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		resp := &Domain{DomainName: "example.org"}
		jData, err := json.Marshal(resp)
		if err != nil {
			t.Error("Could not marshal response")
		}
		w.Write(jData)
	}))
	defer server.Close()

	nc := Mock("username", "apitoken", server.URL)
	Domain, err := nc.EnableAutorenew(&EnableAutorenewForDomainRequest{DomainName: "example.org"})
	if err != nil {
		t.Errorf("Expected 'example.org', got %s", err.Error())
	}
	if Domain.DomainName != "example.org" {
		t.Errorf("Expected 'example.org', got %s", Domain.DomainName)
	}
}

func TestDisableAutorenew(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v4/domains/example.org:disableAutorenew" {
			t.Errorf("Expected to request '/v4/domains/example.org:disableAutorenew', got: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		resp := &Domain{DomainName: "example.org"}
		jData, err := json.Marshal(resp)
		if err != nil {
			t.Error("Could not marshal response")
		}
		w.Write(jData)
	}))
	defer server.Close()

	nc := Mock("username", "apitoken", server.URL)
	Domain, err := nc.DisableAutorenew(&DisableAutorenewForDomainRequest{DomainName: "example.org"})
	if err != nil {
		t.Errorf("Expected 'example.org', got %s", err.Error())
	}
	if Domain.DomainName != "example.org" {
		t.Errorf("Expected 'example.org', got %s", Domain.DomainName)
	}
}

func TestEnableWhoisPrivacy(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v4/domains/example.org:enableWhoisPrivacy" {
			t.Errorf("Expected to request '/v4/domains/example.org:enableWhoisPrivacy', got: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		resp := &Domain{DomainName: "example.org"}
		jData, err := json.Marshal(resp)
		if err != nil {
			t.Error("Could not marshal response")
		}
		w.Write(jData)
	}))
	defer server.Close()

	nc := Mock("username", "apitoken", server.URL)
	Domain, err := nc.EnableWhoisPrivacy(&EnableWhoisPrivacyForDomainRequest{DomainName: "example.org"})
	if err != nil {
		t.Errorf("Expected 'example.org', got %s", err.Error())
	}
	if Domain.DomainName != "example.org" {
		t.Errorf("Expected 'example.org', got %s", Domain.DomainName)
	}
}

func TestDisableWhoisPrivacy(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v4/domains/example.org:disableWhoisPrivacy" {
			t.Errorf("Expected to request '/v4/domains/example.org:disableWhoisPrivacy', got: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		resp := &Domain{DomainName: "example.org"}
		jData, err := json.Marshal(resp)
		if err != nil {
			t.Error("Could not marshal response")
		}
		w.Write(jData)
	}))
	defer server.Close()

	nc := Mock("username", "apitoken", server.URL)
	Domain, err := nc.DisableWhoisPrivacy(&DisableWhoisPrivacyForDomainRequest{DomainName: "example.org"})
	if err != nil {
		t.Errorf("Expected 'example.org', got %s", err.Error())
	}
	if Domain.DomainName != "example.org" {
		t.Errorf("Expected 'example.org', got %s", Domain.DomainName)
	}
}

func TestRenewDomain(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v4/domains/example.org:renew" {
			t.Errorf("Expected to request '/v4/domains/example.org:renew', got: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		//not sure why this has to be done like this
		resp := &RenewDomainResponse{Domain: &Domain{DomainName: "example.org"}}
		jData, err := json.Marshal(resp)
		if err != nil {
			t.Error("Could not marshal response")
		}
		w.Write(jData)
	}))
	defer server.Close()

	nc := Mock("username", "apitoken", server.URL)
	DomainRenewResonse, err := nc.RenewDomain(&RenewDomainRequest{DomainName: "example.org"})
	if err != nil {
		t.Errorf("Expected 'example.org', got %s", err.Error())
	}
	if DomainRenewResonse.Domain.DomainName != "example.org" {
		t.Errorf("Expected 'example.org', got %s", DomainRenewResonse.Domain.DomainName)
	}
}

func TestGetAuthCodeForDomain(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v4/domains/example.org:getAuthCode" {
			t.Errorf("Expected to request '/v4/domains/example.org:getAuthCode', got: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		//not sure why this has to be done like this
		resp := &AuthCodeResponse{AuthCode: "testauthcode"}
		jData, err := json.Marshal(resp)
		if err != nil {
			t.Error("Could not marshal response")
		}
		w.Write(jData)
	}))
	defer server.Close()

	nc := Mock("username", "apitoken", server.URL)
	AuthCodeResponse, err := nc.GetAuthCodeForDomain(&AuthCodeRequest{DomainName: "example.org"})
	if err != nil {
		t.Errorf("Expected 'example.org', got %s", err.Error())
	}
	if AuthCodeResponse.AuthCode != "testauthcode" {
		t.Errorf("Expected 'testauthcode', got %s", AuthCodeResponse.AuthCode)
	}
}

func TestPurchasePrivacyForDomain(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v4/domains/example.org:purchasePrivacy" {
			t.Errorf("Expected to request '/v4/domains/example.org:purchasePrivacy', got: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		//not sure why this has to be done like this
		resp := &RenewDomainResponse{Domain: &Domain{DomainName: "example.org"}}
		jData, err := json.Marshal(resp)
		if err != nil {
			t.Error("Could not marshal response")
		}
		w.Write(jData)
	}))
	defer server.Close()

	nc := Mock("username", "apitoken", server.URL)
	PurchasePrivacy, err := nc.PurchasePrivacy(&PrivacyRequest{DomainName: "example.org"})
	if err != nil {
		t.Errorf("Expected 'example.org', got %s", err.Error())
	}
	if PurchasePrivacy.Domain.DomainName != "example.org" {
		t.Errorf("Expected 'example.org', got %s", PurchasePrivacy.Domain.DomainName)
	}
}

func TestSetNameserversForDomain(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v4/domains/example.org:setNameservers" {
			t.Errorf("Expected to request '/v4/domains/example.org:setNameservers', got: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		//not sure why this has to be done like this
		resp := &Domain{DomainName: "example.org"}
		jData, err := json.Marshal(resp)
		if err != nil {
			t.Error("Could not marshal response")
		}
		w.Write(jData)
	}))
	defer server.Close()

	nc := Mock("username", "apitoken", server.URL)
	Nameservers, err := nc.SetNameservers(&SetNameserversRequest{DomainName: "example.org"})
	if err != nil {
		t.Errorf("Expected 'example.org', got %s", err.Error())
	}
	if Nameservers.DomainName != "example.org" {
		t.Errorf("Expected 'example.org', got %s", Nameservers.DomainName)
	}
}

func TestSetContactsForDomain(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v4/domains/example.org:setContacts" {
			t.Errorf("Expected to request '/v4/domains/example.org:setContacts', got: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		//not sure why this has to be done like this
		resp := &Domain{DomainName: "example.org"}
		jData, err := json.Marshal(resp)
		if err != nil {
			t.Error("Could not marshal response")
		}
		w.Write(jData)
	}))
	defer server.Close()

	nc := Mock("username", "apitoken", server.URL)
	Nameservers, err := nc.SetContacts(&SetContactsRequest{DomainName: "example.org"})
	if err != nil {
		t.Errorf("Expected 'example.org', got %s", err.Error())
	}
	if Nameservers.DomainName != "example.org" {
		t.Errorf("Expected 'example.org', got %s", Nameservers.DomainName)
	}
}

func TestLockForDomain(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v4/domains/example.org:lock" {
			t.Errorf("Expected to request '/v4/domains/example.org:lock', got: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		//not sure why this has to be done like this
		resp := &Domain{DomainName: "example.org"}
		jData, err := json.Marshal(resp)
		if err != nil {
			t.Error("Could not marshal response")
		}
		w.Write(jData)
	}))
	defer server.Close()

	nc := Mock("username", "apitoken", server.URL)
	Lock, err := nc.LockDomain(&LockDomainRequest{DomainName: "example.org"})
	if err != nil {
		t.Errorf("Expected 'example.org', got %s", err.Error())
	}
	if Lock.DomainName != "example.org" {
		t.Errorf("Expected 'example.org', got %s", Lock.DomainName)
	}
}

func TestUnlockForDomain(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v4/domains/example.org:unlock" {
			t.Errorf("Expected to request '/v4/domains/example.org:unlock', got: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		//not sure why this has to be done like this
		resp := &Domain{DomainName: "example.org"}
		jData, err := json.Marshal(resp)
		if err != nil {
			t.Error("Could not marshal response")
		}
		w.Write(jData)
	}))
	defer server.Close()

	nc := Mock("username", "apitoken", server.URL)
	Lock, err := nc.UnlockDomain(&UnlockDomainRequest{DomainName: "example.org"})
	if err != nil {
		t.Errorf("Expected 'example.org', got %s", err.Error())
	}
	if Lock.DomainName != "example.org" {
		t.Errorf("Expected 'example.org', got %s", Lock.DomainName)
	}
}

func TestCheckavailabilityForDomain(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v4/domains:checkAvailability" {
			t.Errorf("Expected to request '/v4/domains:checkAvailability', got: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		//not sure why this has to be done like this
		resp := &SearchResponse{
			Results: []*SearchResult{
				{DomainName: "example.org"},
			},
		}
		//	resp := &Domain{DomainName: "example.org"}
		jData, err := json.Marshal(resp)
		if err != nil {
			t.Error("Could not marshal response")
		}
		w.Write(jData)
	}))
	defer server.Close()

	nc := Mock("username", "apitoken", server.URL)
	CheckAvailabilityResponse, err := nc.CheckAvailability(&AvailabilityRequest{DomainNames: []string{"example.org"}})
	if err != nil {
		t.Errorf("Expected 'example.org', got %s", err.Error())
	}
	if CheckAvailabilityResponse.Results[0].DomainName != "example.org" {
		t.Errorf("Expected 'example.org', got %s", CheckAvailabilityResponse.Results[0].DomainName)
	}
}

func TestSearchForDomain(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v4/domains:search" {
			t.Errorf("Expected to request '/v4/domains:search', got: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		//not sure why this has to be done like this
		resp := &SearchResponse{
			Results: []*SearchResult{
				{DomainName: "example.org"},
			},
		}
		//	resp := &Domain{DomainName: "example.org"}
		jData, err := json.Marshal(resp)
		if err != nil {
			t.Error("Could not marshal response")
		}
		w.Write(jData)
	}))
	defer server.Close()

	nc := Mock("username", "apitoken", server.URL)
	SearchResponse, err := nc.Search(&SearchRequest{Keyword: "example.org"})
	if err != nil {
		t.Errorf("Expected 'example.org', got %s", err.Error())
	}
	if SearchResponse.Results[0].DomainName != "example.org" {
		t.Errorf("Expected 'example.org', got %s", SearchResponse.Results[0].DomainName)
	}
}

func TestSearchStreamForDomain(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v4/domains:searchStream" {
			t.Errorf("Expected to request '/v4/domains:searchStream', got: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		//not sure why this has to be done like this
		resp := &SearchResult{DomainName: "example.org"}
		jData, err := json.Marshal(resp)
		if err != nil {
			t.Error("Could not marshal response")
		}
		w.Write(jData)
	}))
	defer server.Close()

	nc := Mock("username", "apitoken", server.URL)
	SearchResponse, err := nc.SearchStream(&SearchRequest{Keyword: "example.org"})
	if err != nil {
		t.Errorf("Expected 'example.org', got %s", err.Error())
	}
	if SearchResponse.DomainName != "example.org" {
		t.Errorf("Expected 'example.org', got %s", SearchResponse.DomainName)
	}
}

func TestGetPricingForDomain(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v4/domains/example.org:getPricing" {
			t.Errorf("Expected to request '/v4/domains/example.com:getPricing', got: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		//not sure why this has to be done like this
		resp := &PricingResponse{PurchasePrice: 4.99, RenewalPrice: 4.99, TransferPrice: 4.99}
		jData, err := json.Marshal(resp)
		if err != nil {
			t.Error("Could not marshal response")
		}
		w.Write(jData)
	}))
	defer server.Close()

	nc := Mock("username", "apitoken", server.URL)
	PricingResponse, err := nc.GetPricingForDomain(&PricingRequest{DomainName: "example.org"})
	if err != nil {
		t.Errorf("Expected 'example.org', got %s", err.Error())
	}
	if PricingResponse.PurchasePrice != 4.99 {
		t.Errorf("Expected '4.99', got %f", PricingResponse.PurchasePrice)
	}

	if PricingResponse.RenewalPrice != 4.99 {
		t.Errorf("Expected '4.99', got %f", PricingResponse.RenewalPrice)
	}

	if PricingResponse.TransferPrice != 4.99 {
		t.Errorf("Expected '4.99', got %f", PricingResponse.TransferPrice)
	}
}
