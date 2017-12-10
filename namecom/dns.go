package namecom

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
)

var _ = bytes.MinRead

// ListRecords returns all records for a zone.
func (n *NameCom) ListRecords(request *ListRecordsRequest) (*ListRecordsResponse, error) {
	endpoint := fmt.Sprintf("/v4/domains/%s/records", request.GetDomainName())

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

	resp := &ListRecordsResponse{}

	err = json.NewDecoder(body).Decode(resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetRecord returns details about an individual record.
func (n *NameCom) GetRecord(request *GetRecordRequest) (*Record, error) {
	endpoint := fmt.Sprintf("/v4/domains/%s/records/%d", request.GetDomainName(), request.GetId())

	values := url.Values{}

	body, err := n.Get(endpoint, values)
	if err != nil {
		return nil, err
	}

	resp := &Record{}

	err = json.NewDecoder(body).Decode(resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// CreateRecord creates a new record in the zone.
func (n *NameCom) CreateRecord(request *Record) (*Record, error) {
	endpoint := fmt.Sprintf("/v4/domains/%s/records", request.GetDomainName())

	post := &bytes.Buffer{}
	json.NewEncoder(post).Encode(request)

	body, err := n.Post(endpoint, post)
	if err != nil {
		return nil, err
	}

	resp := &Record{}

	err = json.NewDecoder(body).Decode(resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// UpdateRecord replaces the record with the new record that is passed.
func (n *NameCom) UpdateRecord(request *Record) (*Record, error) {
	endpoint := fmt.Sprintf("/v4/domains/%s/records/%d", request.GetDomainName(), request.GetId())

	post := &bytes.Buffer{}
	json.NewEncoder(post).Encode(request)

	body, err := n.Put(endpoint, post)
	if err != nil {
		return nil, err
	}

	resp := &Record{}

	err = json.NewDecoder(body).Decode(resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// DeleteRecord deletes a record from the zone.
func (n *NameCom) DeleteRecord(request *DeleteRecordRequest) (*EmptyResponse, error) {
	endpoint := fmt.Sprintf("/v4/domains/%s/records/%d", request.GetDomainName(), request.GetId())

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
