package namecom

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListOrders(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v4/orders" {
			t.Errorf("Expected to request ' /v4/orders', got: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		resp := &ListOrdersResponse{
			Orders: []*Order{
				{ID: 123456789},
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
	ListOrdersResponse, err := nc.ListOrders(&ListOrdersRequest{})
	if err != nil {
		t.Errorf("Expected '123456789', got %s", err.Error())
	}
	if ListOrdersResponse.Orders[0].ID != 123456789 {
		t.Errorf("Expected '123456789', got %d", ListOrdersResponse.Orders[0].ID)
	}
}

func TestGetOrder(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v4/orders/123456789" {
			t.Errorf("Expected to request '/v4/orders/123456789', got: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		resp := &Order{ID: 123456789}
		jData, err := json.Marshal(resp)
		if err != nil {
			t.Error("Could not marshal response")
		}
		w.Write(jData)
	}))
	defer server.Close()

	nc := Mock("username", "apitoken", server.URL)
	GetOrderRequest, err := nc.GetOrder(&GetOrderRequest{OrderID: 123456789})
	if err != nil {
		t.Errorf("Expected 'example.org', got %s", err.Error())
	}
	if GetOrderRequest.ID != 123456789 {
		t.Errorf("Expected '123456789', got %d", GetOrderRequest.ID)
	}

}
