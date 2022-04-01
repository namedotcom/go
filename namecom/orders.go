package namecom

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
)

var _ = bytes.MinRead

// ListOrders returns all orders in the account. It omits some information that can be retrieved from GetOrder.
func (n *NameCom) ListOrders(request *ListOrdersRequest) (*ListOrdersResponse, error) {
	endpoint := fmt.Sprintf("/v4/orders")

	values := url.Values{}
	if request.PerPage != 0 {
		values.Set("perPage", fmt.Sprintf("%d", request.PerPage))
	}
	if request.Page != 0 {
		values.Set("page", fmt.Sprintf("%d", request.Page))
	}

	body, err := n.get(endpoint, values)
	if err != nil {
		return nil, err
	}

	resp := &ListOrdersResponse{}

	err = json.NewDecoder(body).Decode(resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetOrder returns details about a specific order
func (n *NameCom) GetOrder(request *GetOrderRequest) (*Order, error) {
	endpoint := fmt.Sprintf("/v4/orders/%d", request.OrderID)

	values := url.Values{}

	body, err := n.get(endpoint, values)
	if err != nil {
		return nil, err
	}

	resp := &Order{}

	err = json.NewDecoder(body).Decode(resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
