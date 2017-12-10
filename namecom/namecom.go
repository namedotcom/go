package namecom

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/pkg/errors"
)

type NameCom struct {
	Server string
	User   string
	Token  string
	Client *http.Client
}

func New(user, token string) *NameCom {
	return &NameCom{
		Server: "api.name.com",
		User:   user,
		Token:  token,
		Client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func Test(user, token string) *NameCom {
	return &NameCom{
		Server: "api.dev.name.com",
		User:   user,
		Token:  token,
		Client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (er ErrorResponse) Error() string {
	return er.Message + ": " + er.Details
}

func (n *NameCom) ErrorResponse(resp *http.Response) error {
	er := &ErrorResponse{}
	err := json.NewDecoder(resp.Body).Decode(er)
	if err != nil {
		return errors.Wrap(err, "api returned unexpected response")
	}

	return errors.WithStack(er)
}

func (n *NameCom) Get(endpoint string, values url.Values) (io.Reader, error) {
	if len(values) == 0 {
		endpoint = endpoint + "?" + values.Encode()
	}
	return n.doRequest("GET", endpoint, nil)
}

func (n *NameCom) Post(endpoint string, post io.Reader) (io.Reader, error) {
	return n.doRequest("POST", endpoint, post)
}

func (n *NameCom) Put(endpoint string, post io.Reader) (io.Reader, error) {
	return n.doRequest("PUT", endpoint, post)
}

func (n *NameCom) Delete(endpoint string, post io.Reader) (io.Reader, error) {
	return n.doRequest("DELETE", endpoint, post)
}

func (n *NameCom) doRequest(method, endpoint string, post io.Reader) (io.Reader, error) {
	url := "https://" + n.Server + endpoint

	req, err := http.NewRequest(method, url, post)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(n.User, n.Token)

	resp, err := n.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, n.ErrorResponse(resp)
	}

	return resp.Body, nil
}
