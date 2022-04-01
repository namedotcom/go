package namecom

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListTransfers(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v4/transfers" {
			t.Errorf("Expected to request '/v4/transfers', got: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		resp := &ListTransfersResponse{
			Transfers: []*Transfer{
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
	ListTransfersRequest, err := nc.ListTransfers(&ListTransfersRequest{})
	if err != nil {
		t.Errorf("Expected 'example.org', got %s", err.Error())
	}
	if ListTransfersRequest.Transfers[0].DomainName != "example.org" {
		t.Errorf("Expected 'example.org', got %s", ListTransfersRequest.Transfers[0].DomainName)
	}
}

func TestGetTransfer(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v4/transfers/example.org" {
			t.Errorf("Expected to request /v4/transfers/example.org', got: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		resp := &Transfer{DomainName: "example.org"}
		jData, err := json.Marshal(resp)
		if err != nil {
			t.Error("Could not marshal response")
		}
		w.Write(jData)
	}))
	defer server.Close()

	nc := Mock("username", "apitoken", server.URL)
	GetTransfer, err := nc.GetTransfer(&GetTransferRequest{DomainName: "example.org"})
	if err != nil {
		t.Errorf("Expected 'example.org', got %s", err.Error())
	}
	if GetTransfer.DomainName != "example.org" {
		t.Errorf("Expected 'example.org', got %s", GetTransfer.DomainName)
	}
}

func TestCreateTransfer(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v4/transfers" {
			t.Errorf("Expected to request /v4/transfers', got: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		resp := &CreateTransferResponse{Transfer: &Transfer{DomainName: "example.org"}}
		jData, err := json.Marshal(resp)
		if err != nil {
			t.Error("Could not marshal response")
		}
		w.Write(jData)
	}))
	defer server.Close()

	nc := Mock("username", "apitoken", server.URL)
	CreateTransferResponse, err := nc.CreateTransfer(&CreateTransferRequest{DomainName: "example.org"})
	if err != nil {
		t.Errorf("Expected 'example.org', got %s", err.Error())
	}
	if CreateTransferResponse.Transfer.DomainName != "example.org" {
		t.Errorf("Expected 'example.org', got %s", CreateTransferResponse.Transfer.DomainName)
	}
}

func TestCancelTransfer(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v4/transfers/example.org:cancel" {
			t.Errorf("Expected to request '/v4/transfers/example.org:cancel', got: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		resp := &Transfer{DomainName: "example.org"}
		jData, err := json.Marshal(resp)
		if err != nil {
			t.Error("Could not marshal response")
		}
		w.Write(jData)
	}))
	defer server.Close()

	nc := Mock("username", "apitoken", server.URL)
	CancelTransferRequest, err := nc.CancelTransfer(&CancelTransferRequest{DomainName: "example.org"})
	if err != nil {
		t.Errorf("Expected 'example.org', got %s", err.Error())
	}
	if CancelTransferRequest.DomainName != "example.org" {
		t.Errorf("Expected 'example.org', got %s", CancelTransferRequest.DomainName)
	}
}

/*
func TestDeleteRecord(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v4/domains/example.org/records/12345" {
			t.Errorf("Expected to request /v4/domains/example.org/records/12345', got: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		resp := &Record{DomainName: "example.org", ID: 12345}
		jData, err := json.Marshal(resp)
		if err != nil {
			t.Error("Could not marshal response")
		}
		w.Write(jData)
	}))
	defer server.Close()

	nc := Mock("username", "apitoken", server.URL)
	_, err := nc.DeleteRecord(&DeleteRecordRequest{DomainName: "example.org", ID: 12345})
	if err != nil {
		t.Errorf("Expected 'example.org', got %s", err.Error())
	}

}
*/
