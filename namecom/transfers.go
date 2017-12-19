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
	if request.PerPage != 0 {
		values.Set("perPage", fmt.Sprintf("%d", request.PerPage))
	}
	if request.Page != 0 {
		values.Set("page", fmt.Sprintf("%d", request.Page))
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
	endpoint := fmt.Sprintf("/v4/transfers/%s", request.DomainName)

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
