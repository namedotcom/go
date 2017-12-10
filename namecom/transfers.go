package namecom

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
)

var _ = bytes.MinRead

// ListTransfers lists all pending transfer in requests.
func (n *NameCom) ListTransfers(request *ListTransfersRequest) (*ListTransfersResponse, error) {
	endpoint := fmt.Sprintf("/v4/transfers")

	values := url.Values{}
	if v := request.GetDomainName(); v != "" {
		values.Set("domain_name", v)
	}
	if v := request.GetPerPage(); v != 0 {
		values.Set("per_page", fmt.Sprintf("%d", v))
	}
	if v := request.GetPage(); v != 0 {
		values.Set("page", fmt.Sprintf("%d", v))
	}

	body, err := n.Get(endpoint, values)
	if err != nil {
		return nil, err
	}

	resp := &ListTransfersResponse{}

	err = json.NewDecoder(body).Decode(resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetTransfer gets details for a transfer request.
func (n *NameCom) GetTransfer(request *GetTransferRequest) (*Transfer, error) {
	endpoint := fmt.Sprintf("/v4/transfers/%s", request.GetDomainName())

	values := url.Values{}

	body, err := n.Get(endpoint, values)
	if err != nil {
		return nil, err
	}

	resp := &Transfer{}

	err = json.NewDecoder(body).Decode(resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// CreateTransfer purchases a new domain transfer request.
func (n *NameCom) CreateTransfer(request *CreateTransferRequest) (*CreateTransferResponse, error) {
	endpoint := fmt.Sprintf("/v4/transfers")

	post := &bytes.Buffer{}
	json.NewEncoder(post).Encode(request)

	body, err := n.Post(endpoint, post)
	if err != nil {
		return nil, err
	}

	resp := &CreateTransferResponse{}

	err = json.NewDecoder(body).Decode(resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// CancelTransfer cancels a pending transfer request and refunds the amount to account credit.
func (n *NameCom) CancelTransfer(request *CancelTransferRequest) (*Transfer, error) {
	endpoint := fmt.Sprintf("/v4/transfers")

	post := &bytes.Buffer{}
	json.NewEncoder(post).Encode(request)

	body, err := n.Post(endpoint, post)
	if err != nil {
		return nil, err
	}

	resp := &Transfer{}

	err = json.NewDecoder(body).Decode(resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
