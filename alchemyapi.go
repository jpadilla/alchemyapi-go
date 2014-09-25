// Package alchemyapi provides the binding for AlchemyAPI.
package alchemyapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

// alchemyAPI is the public AlchemyAPI URL for APIs.
const alchemyAPI = "http://access.alchemyapi.com"

// textExtractionAPI is the public AlchemyAPI URL for URLGetText.
const textExtractionAPI = alchemyAPI + "/calls/url/URLGetText"

// titleExtractionAPI is the public AlchemyAPI URL for URLGetTitle.
const titleExtractionAPI = alchemyAPI + "/calls/url/URLGetTitle"

// AlchemyAPI is used to invoke API calls.
type AlchemyAPI struct {
	apikey string
}

// New returns a new AlchemyAPI client.
func New(apikey string) *Alchemy {
	return &AlchemyAPI{apikey}
}

// GetTitle returns the extracted title for a given URL.
func (client *AlchemyAPI) GetTitle(requestedURL string, options GetTitleOptions) (*GetTitleResponse, error) {
	addr := titleExtractionAPI + "?"

	params := url.Values{}
	params.Add("apikey", client.apikey)
	params.Add("url", requestedURL)
	params.Add("outputMode", "json")

	if options.UseMetadata > 0 {
		params.Add("useMetadata", strconv.Itoa(options.UseMetadata))
	}

	addr += params.Encode()

	// Make request
	resp, err := http.Get(addr)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 500 {
		body, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("Got non 200 status code: %s %q", resp.Status, body)
	}

	// Read the JSON message from the body.
	response := &GetTitleResponse{}
	decoder := json.NewDecoder(resp.Body)

	if err := decoder.Decode(&response); err != nil {
		return nil, err
	}

	return response, nil
}

// GetText returns the extracted text for a given URL.
func (client *AlchemyAPI) GetText(requestedURL string, options GetTextOptions) (*GetTextResponse, error) {
	addr := textExtractionAPI + "?"

	params := url.Values{}
	params.Add("apikey", client.apikey)
	params.Add("url", requestedURL)
	params.Add("outputMode", "json")

	if options.UseMetadata > 0 {
		params.Add("useMetadata", strconv.Itoa(options.UseMetadata))
	}

	if options.ExtractLinks > 0 {
		params.Add("extractLinks", strconv.Itoa(options.ExtractLinks))
	}

	addr += params.Encode()

	// Make request
	resp, err := http.Get(addr)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 500 {
		body, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("Got non 200 status code: %s %q", resp.Status, body)
	}

	// Read the JSON message from the body.
	response := &GetTextResponse{}
	decoder := json.NewDecoder(resp.Body)

	if err := decoder.Decode(&response); err != nil {
		return nil, err
	}

	return response, nil
}
