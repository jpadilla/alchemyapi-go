package alchemy

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

const (
	AlchemyAPI         = "http://access.alchemyapi.com"
	TextExtractionAPI  = AlchemyAPI + "/calls/url/URLGetText"
	TitleExtractionAPI = AlchemyAPI + "/calls/url/URLGetTitle"
)

type Alchemy struct {
	apikey string
}

func New(apikey string) *Alchemy {
	return &Alchemy{apikey}
}

func (client *Alchemy) GetTitle(requestedURL string, options GetTitleOptions) (*GetTitleResponse, error) {
	addr := TitleExtractionAPI + "?"

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

func (client *Alchemy) GetText(requestedURL string, options GetTextOptions) (*GetTextResponse, error) {
	addr := TextExtractionAPI + "?"

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
