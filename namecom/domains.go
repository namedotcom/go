package namecom

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
)

var _ = bytes.MinRead

// List returns all domains
func (n *NameCom) ListDomains(request *ListDomainsRequest) (*ListDomainsResponse, error) {
	endpoint := fmt.Sprintf("/v4/domains")

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

	resp := &ListDomainsResponse{}

	err = json.NewDecoder(body).Decode(resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Get returns details about a specific domain
func (n *NameCom) GetDomain(request *GetDomainRequest) (*Domain, error) {
	endpoint := fmt.Sprintf("/v4/domains/%s", request.GetDomainName())

	values := url.Values{}

	body, err := n.Get(endpoint, values)
	if err != nil {
		return nil, err
	}

	resp := &Domain{}

	err = json.NewDecoder(body).Decode(resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Create purchases a new domain. Domains that are not regularly priced require the purchase_price field to be specified.
func (n *NameCom) CreateDomain(request *CreateDomainRequest) (*CreateDomainResponse, error) {
	endpoint := fmt.Sprintf("/v4/domains")

	post := &bytes.Buffer{}
	json.NewEncoder(post).Encode(request)

	body, err := n.Post(endpoint, post)
	if err != nil {
		return nil, err
	}

	resp := &CreateDomainResponse{}

	err = json.NewDecoder(body).Decode(resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Update can change only the autorenew_enabled field, since all other modifications require an update at the registry, which would make this command non-idempotent.
func (n *NameCom) UpdateDomain(request *Domain) (*Domain, error) {
	endpoint := fmt.Sprintf("/v4/domains/%s", request.GetDomainName())

	post := &bytes.Buffer{}
	json.NewEncoder(post).Encode(request)

	body, err := n.Put(endpoint, post)
	if err != nil {
		return nil, err
	}

	resp := &Domain{}

	err = json.NewDecoder(body).Decode(resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Renew will renew a domain. Purchase_price is required if the renewal is not regularly priced.
func (n *NameCom) RenewDomain(request *RenewDomainRequest) (*RenewDomainResponse, error) {
	endpoint := fmt.Sprintf("/v4/domains")

	post := &bytes.Buffer{}
	json.NewEncoder(post).Encode(request)

	body, err := n.Post(endpoint, post)
	if err != nil {
		return nil, err
	}

	resp := &RenewDomainResponse{}

	err = json.NewDecoder(body).Decode(resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetAuthCode returns the Transfer Authorization Code for the domain.
func (n *NameCom) GetAuthCodeForDomain(request *AuthCodeRequest) (*AuthCodeResponse, error) {
	endpoint := fmt.Sprintf("/v4/domains")

	values := url.Values{}
	if v := request.GetDomainName(); v != "" {
		values.Set("domain_name", v)
	}

	body, err := n.Get(endpoint, values)
	if err != nil {
		return nil, err
	}

	resp := &AuthCodeResponse{}

	err = json.NewDecoder(body).Decode(resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// PurchasePrivacy will add Whois Privacy protection to a domain or will an renew existing subscription.
func (n *NameCom) PurchasePrivacy(request *PrivacyRequest) (*PrivacyResponse, error) {
	endpoint := fmt.Sprintf("/v4/domains")

	post := &bytes.Buffer{}
	json.NewEncoder(post).Encode(request)

	body, err := n.Post(endpoint, post)
	if err != nil {
		return nil, err
	}

	resp := &PrivacyResponse{}

	err = json.NewDecoder(body).Decode(resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// Search will perform a search for specified keywords.
func (n *NameCom) Search(request *SearchRequest) (*SearchResponse, error) {
	endpoint := fmt.Sprintf("/v4/domains:search")

	post := &bytes.Buffer{}
	json.NewEncoder(post).Encode(request)

	body, err := n.Post(endpoint, post)
	if err != nil {
		return nil, err
	}

	resp := &SearchResponse{}

	err = json.NewDecoder(body).Decode(resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// SearchStream will return JSON encoded SearchResults as they are recieved from the registry. This can allow clients to react to results before the search is fully completed.
func (n *NameCom) SearchStream(request *SearchRequest) (*SearchResult, error) {
	endpoint := fmt.Sprintf("/v4/domains:searchStream")

	post := &bytes.Buffer{}
	json.NewEncoder(post).Encode(request)

	body, err := n.Post(endpoint, post)
	if err != nil {
		return nil, err
	}

	resp := &SearchResult{}

	err = json.NewDecoder(body).Decode(resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// SetNameservers will set the nameservers for the Domain.
func (n *NameCom) SetNameservers(request *SetNameserversRequest) (*Domain, error) {
	endpoint := fmt.Sprintf("/v4/domains")

	post := &bytes.Buffer{}
	json.NewEncoder(post).Encode(request)

	body, err := n.Post(endpoint, post)
	if err != nil {
		return nil, err
	}

	resp := &Domain{}

	err = json.NewDecoder(body).Decode(resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// SetContacts will set the contacts for the Domain.
func (n *NameCom) SetContacts(request *SetContactsRequest) (*Domain, error) {
	endpoint := fmt.Sprintf("/v4/domains")

	post := &bytes.Buffer{}
	json.NewEncoder(post).Encode(request)

	body, err := n.Post(endpoint, post)
	if err != nil {
		return nil, err
	}

	resp := &Domain{}

	err = json.NewDecoder(body).Decode(resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// LockDomain will lock a domain so that it cannot be transfered to another registrar.
func (n *NameCom) LockDomain(request *LockDomainRequest) (*Domain, error) {
	endpoint := fmt.Sprintf("/v4/domains")

	post := &bytes.Buffer{}
	json.NewEncoder(post).Encode(request)

	body, err := n.Post(endpoint, post)
	if err != nil {
		return nil, err
	}

	resp := &Domain{}

	err = json.NewDecoder(body).Decode(resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// UnlockDomain will unlock a domain so that it can be transfered to another registrar.
func (n *NameCom) UnlockDomain(request *UnlockDomainRequest) (*Domain, error) {
	endpoint := fmt.Sprintf("/v4/domains")

	post := &bytes.Buffer{}
	json.NewEncoder(post).Encode(request)

	body, err := n.Post(endpoint, post)
	if err != nil {
		return nil, err
	}

	resp := &Domain{}

	err = json.NewDecoder(body).Decode(resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
