// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "coindrop-ap": Application Media Types
//
// Command:
// $ goagen
// --design=github.com/waymobetta/go-coindrop-ap-server/design
// --out=$(GOPATH)/src/github.com/waymobetta/go-coindrop-ap-server
// --version=v1.3.1

package client

import (
	"github.com/goadesign/goa"
	"net/http"
)

// A standard error response (default view)
//
// Identifier: application/standard_error+json; view=default
type StandardError struct {
	// A code that describes the error
	Code int `form:"code" json:"code" yaml:"code" xml:"code"`
	// A message that describes the error
	Message string `form:"message" json:"message" yaml:"message" xml:"message"`
}

// Validate validates the StandardError media type instance.
func (mt *StandardError) Validate() (err error) {

	if mt.Message == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "message"))
	}
	return
}

// DecodeStandardError decodes the StandardError instance encoded in resp body.
func (c *Client) DecodeStandardError(resp *http.Response) (*StandardError, error) {
	var decoded StandardError
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// Health check (default view)
//
// Identifier: application/vnd.healthcheck+json; view=default
type Healthcheck struct {
	// Status
	Status string `form:"status" json:"status" yaml:"status" xml:"status"`
}

// Validate validates the Healthcheck media type instance.
func (mt *Healthcheck) Validate() (err error) {
	if mt.Status == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "status"))
	}
	return
}

// DecodeHealthcheck decodes the Healthcheck instance encoded in resp body.
func (c *Client) DecodeHealthcheck(resp *http.Response) (*Healthcheck, error) {
	var decoded Healthcheck
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}
