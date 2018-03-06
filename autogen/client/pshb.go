// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "feedpushr": pshb Resource Client
//
// Command:
// $ goagen
// --design=github.com/ncarlier/feedpushr/design
// --out=$(GOPATH)/src/github.com/ncarlier/feedpushr/autogen
// --version=v1.3.1

package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// PubPshbPath computes a request path to the pub action of pshb.
func PubPshbPath() string {

	return fmt.Sprintf("/v1/pshb")
}

// Publication endpoint for PSHB hubs
func (c *Client) PubPshb(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewPubPshbRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewPubPshbRequest create the request corresponding to the pub action endpoint of the pshb resource.
func (c *Client) NewPubPshbRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// SubPshbPath computes a request path to the sub action of pshb.
func SubPshbPath() string {

	return fmt.Sprintf("/v1/pshb")
}

// Callback to validate the (un)subscription to the topic of a Hub
func (c *Client) SubPshb(ctx context.Context, path string, hubChallenge string, hubMode string, hubTopic string, hubLeaseSeconds *int) (*http.Response, error) {
	req, err := c.NewSubPshbRequest(ctx, path, hubChallenge, hubMode, hubTopic, hubLeaseSeconds)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewSubPshbRequest create the request corresponding to the sub action endpoint of the pshb resource.
func (c *Client) NewSubPshbRequest(ctx context.Context, path string, hubChallenge string, hubMode string, hubTopic string, hubLeaseSeconds *int) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	values.Set("hub.challenge", hubChallenge)
	values.Set("hub.mode", hubMode)
	values.Set("hub.topic", hubTopic)
	if hubLeaseSeconds != nil {
		tmp15 := strconv.Itoa(*hubLeaseSeconds)
		values.Set("hub.lease_seconds", tmp15)
	}
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}