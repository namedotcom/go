package namecom

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListRecords(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v4/domains/example.org/records" {
			t.Errorf("Expected to request '/v4/domains/example.org/records', got: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		resp := &ListRecordsResponse{
			Records: []*Record{
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
	ListRecordsResponse, err := nc.ListRecords(&ListRecordsRequest{DomainName: "example.org"})
	if err != nil {
		t.Errorf("Expected 'example.org', got %s", err.Error())
	}
	if ListRecordsResponse.Records[0].DomainName != "example.org" {
		t.Errorf("Expected 'example.org', got %s", ListRecordsResponse.Records[0].DomainName)
	}
}

func TestGetRecord(t *testing.T) {
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
	GetRecordResponse, err := nc.GetRecord(&GetRecordRequest{DomainName: "example.org", ID: 12345})
	if err != nil {
		t.Errorf("Expected 'example.org', got %s", err.Error())
	}
	if GetRecordResponse.DomainName != "example.org" {
		t.Errorf("Expected 'example.org', got %s", GetRecordResponse.DomainName)
	}

	if GetRecordResponse.ID != 12345 {
		t.Errorf("Expected '12345, got %d", GetRecordResponse.ID)
	}
}

func TestCreateRecord(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v4/domains/example.org/records" {
			t.Errorf("Expected to request /v4/domains/example.org/records', got: %s", r.URL.Path)
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
	CreateRecordResponse, err := nc.CreateRecord(&Record{DomainName: "example.org", ID: 12345})
	if err != nil {
		t.Errorf("Expected 'example.org', got %s", err.Error())
	}
	if CreateRecordResponse.DomainName != "example.org" {
		t.Errorf("Expected 'example.org', got %s", CreateRecordResponse.DomainName)
	}

	if CreateRecordResponse.ID != 12345 {
		t.Errorf("Expected '12345, got %d", CreateRecordResponse.ID)
	}
}

func TestUpdateRecord(t *testing.T) {
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
	UpdateRecordResponse, err := nc.UpdateRecord(&Record{DomainName: "example.org", ID: 12345})
	if err != nil {
		t.Errorf("Expected 'example.org', got %s", err.Error())
	}
	if UpdateRecordResponse.DomainName != "example.org" {
		t.Errorf("Expected 'example.org', got %s", UpdateRecordResponse.DomainName)
	}

	if UpdateRecordResponse.ID != 12345 {
		t.Errorf("Expected '12345, got %d", UpdateRecordResponse.ID)
	}
}

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
