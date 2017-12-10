package namecom

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
)

var _ = bytes.MinRead

// ListVanityNameservers lists all nameservers registered with the registry.
func (n *NameCom) ListVanityNameservers(request *ListVanityNameserversRequest) (*ListVanityNameserversResponse, error) {
	endpoint := fmt.Sprintf("/v4/domains/%s/vanity_nameservers", request.GetDomainName())

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

	resp := &ListVanityNameserversResponse{}

	err = json.NewDecoder(body).Decode(resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetVanityNameserver gets the details for a vanity nameserver registered with the registry.
func (n *NameCom) GetVanityNameserver(request *GetVanityNameserverRequest) (*VanityNameserver, error) {
	endpoint := fmt.Sprintf("/v4/domains/%s/vanity_nameservers/%s", request.GetDomainName(), request.GetHostname())

	values := url.Values{}

	body, err := n.Get(endpoint, values)
	if err != nil {
		return nil, err
	}

	resp := &VanityNameserver{}

	err = json.NewDecoder(body).Decode(resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// CreateVanityNameserver registers a nameserver with the registry.
func (n *NameCom) CreateVanityNameserver(request *CreateVanityNameserverRequest) (*VanityNameserver, error) {
	endpoint := fmt.Sprintf("/v4/domains/%s/vanity_nameservers", request.GetDomainName())

	post := &bytes.Buffer{}
	json.NewEncoder(post).Encode(request)

	body, err := n.Post(endpoint, post)
	if err != nil {
		return nil, err
	}

	resp := &VanityNameserver{}

	err = json.NewDecoder(body).Decode(resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// UpdateVanityNameserver allows you to update the glue record IP addresses at the registry.
func (n *NameCom) UpdateVanityNameserver(request *VanityNameserver) (*VanityNameserver, error) {
	endpoint := fmt.Sprintf("/v4/domains/%s/vanity_nameservers/%s", request.GetDomainName(), request.GetHostname())

	post := &bytes.Buffer{}
	json.NewEncoder(post).Encode(request)

	body, err := n.Put(endpoint, post)
	if err != nil {
		return nil, err
	}

	resp := &VanityNameserver{}

	err = json.NewDecoder(body).Decode(resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// DeleteVanityNameserver unregisteres the nameserver at the registry. This might fail if the registry believes the nameserver is in use.
func (n *NameCom) DeleteVanityNameserver(request *DeleteVanityNameserverRequest) (*EmptyResponse, error) {
	endpoint := fmt.Sprintf("/v4/domains/%s/vanity_nameservers/%s", request.GetDomainName(), request.GetHostname())

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
