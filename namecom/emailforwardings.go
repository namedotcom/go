package namecom

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
)

var _ = bytes.MinRead

// ListEmailForwardings returns a pagenated list of email forwarding entries for a domain.
func (n *NameCom) ListEmailForwardings(request *ListEmailForwardingsRequest) (*ListEmailForwardingsResponse, error) {
	endpoint := fmt.Sprintf("/v4/domains/%s/email/forwarding", request.GetDomainName())

	values := url.Values{}
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

	resp := &ListEmailForwardingsResponse{}

	err = json.NewDecoder(body).Decode(resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetEmailForwarding returns an email forwarding entry.
func (n *NameCom) GetEmailForwarding(request *GetEmailForwardingRequest) (*EmailForwarding, error) {
	endpoint := fmt.Sprintf("/v4/domains/%s/email/forwarding/%s", request.GetDomainName(), request.GetEmailBox())

	values := url.Values{}

	body, err := n.Get(endpoint, values)
	if err != nil {
		return nil, err
	}

	resp := &EmailForwarding{}

	err = json.NewDecoder(body).Decode(resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// CreateEmailForwarding creates an email forwarding entry. If this is the first email forwarding entry, it may modify the MX records for the domain accordingly.
func (n *NameCom) CreateEmailForwarding(request *EmailForwarding) (*EmailForwarding, error) {
	endpoint := fmt.Sprintf("/v4/domains/%s/email/forwarding", request.GetDomainName())

	post := &bytes.Buffer{}
	json.NewEncoder(post).Encode(request)

	body, err := n.Post(endpoint, post)
	if err != nil {
		return nil, err
	}

	resp := &EmailForwarding{}

	err = json.NewDecoder(body).Decode(resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// UpdateEmailForwarding updates which email address the email is being forwarded to.
func (n *NameCom) UpdateEmailForwarding(request *EmailForwarding) (*EmailForwarding, error) {
	endpoint := fmt.Sprintf("/v4/domains/%s/email/forwarding/%s", request.GetDomainName(), request.GetEmailBox())

	post := &bytes.Buffer{}
	json.NewEncoder(post).Encode(request)

	body, err := n.Put(endpoint, post)
	if err != nil {
		return nil, err
	}

	resp := &EmailForwarding{}

	err = json.NewDecoder(body).Decode(resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// DeleteEmailForwarding deletes the email forwarding entry.
func (n *NameCom) DeleteEmailForwarding(request *DeleteEmailForwardingRequest) (*EmptyResponse, error) {
	endpoint := fmt.Sprintf("/v4/domains/%s/email/forwarding/%s", request.GetDomainName(), request.GetEmailBox())

	post := &bytes.Buffer{}
	json.NewEncoder(post).Encode(request)

	body, err := n.Delete(endpoint, post)
	if err != nil {
		return nil, err
	}

	resp := &EmptyResponse{}

	err = json.NewDecoder(body).Decode(resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
