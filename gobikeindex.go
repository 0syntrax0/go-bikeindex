package gobikeindex

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"
)

const (
	baseURL       = "https://bikeindex.org/"
	apiVersion3   = "v3"
	apiURL        = baseURL + "api/" + apiVersion3
	clientTimeout = time.Second * 30
)

// The default http.Client to use if nothing is specified
var defaultClient = &http.Client{
	Timeout: clientTimeout,
}

// BikeIndex stores basic infor when connecting to their API
type BikeIndex struct {
	ID          string
	Secret      string
	Name        string
	BaseURL     string
	APIVersion  string
	APIURL      string
	CallbackURL *string
	HTTPClient  *http.Client
}

// NewBikeIndexClient creates a new bikeindex
func NewBikeIndexClient(id, secret string) *BikeIndex {
	return NewBikeindexCustomClient(id, secret, nil)
}

// NewBikeindexCustomClient creates a new bikeindex client with custom http.client
func NewBikeindexCustomClient(id, secret string, HTTPClient *http.Client) *BikeIndex {
	if HTTPClient == nil {
		HTTPClient = defaultClient
	}

	return &BikeIndex{
		ID:         id,
		Secret:     secret,
		BaseURL:    baseURL,
		APIURL:     apiURL,
		APIVersion: apiVersion3,
		HTTPClient: HTTPClient,
	}
}

// performs a get request
func (bi *BikeIndex) get(url string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, bi.APIURL, nil)
	if err != nil {
		return nil, err
	}
	return bi.do(req)
}

// perform a delete request
func (bi *BikeIndex) delete(url string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return nil, err
	}
	// req.SetBasicAuth(twilio.getBasicAuthCredentials())
	return bi.do(req)
}

// perform a post request
func (bi *BikeIndex) post(values url.Values, url string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodPost, bi.APIURL, strings.NewReader(values.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	return bi.do(req)
}

// perform a put request
func (bi *BikeIndex) put(values url.Values, url string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodPut, bi.APIURL, strings.NewReader(values.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	return bi.do(req)
}

// excecutes the request
func (bi *BikeIndex) do(req *http.Request) (*http.Response, error) {
	client := bi.HTTPClient
	if client == nil {
		client = defaultClient
	}
	return client.Do(req)
}

func (bi *BikeIndex) getJSON(url string, result interface{}) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %v", err)
	}
	// req.SetBasicAuth(bi.getBasicAuthCredentials())
	resp, err := bi.do(req)
	if err != nil {
		return fmt.Errorf("failed to submit HTTP request: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		re := Exception{}
		json.NewDecoder(resp.Body).Decode(&re)
		return re
	}

	return json.NewDecoder(resp.Body).Decode(&result)
}

// buildURL -
func (bi *BikeIndex) buildURL(paths ...string) string {
	return bi.APIURL + "/" + path.Join(paths...)
}
