// Code generated by go-swagger; DO NOT EDIT.

package file

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new file API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for file API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	AddDockerFile(params *AddDockerFileParams, authInfo runtime.ClientAuthInfoWriter) (*AddDockerFileOK, error)

	CreateArchive(params *CreateArchiveParams, authInfo runtime.ClientAuthInfoWriter) (*CreateArchiveOK, error)

	GetFile(params *GetFileParams, authInfo runtime.ClientAuthInfoWriter) (*GetFileOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  AddDockerFile add docker file API
*/
func (a *Client) AddDockerFile(params *AddDockerFileParams, authInfo runtime.ClientAuthInfoWriter) (*AddDockerFileOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddDockerFileParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "AddDockerFile",
		Method:             "PUT",
		PathPattern:        "/api/file/{fid}/add-dockerfile-text",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AddDockerFileReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AddDockerFileOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for AddDockerFile: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CreateArchive create archive API
*/
func (a *Client) CreateArchive(params *CreateArchiveParams, authInfo runtime.ClientAuthInfoWriter) (*CreateArchiveOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateArchiveParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateArchive",
		Method:             "POST",
		PathPattern:        "/api/file/archive",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"multipart/form-data"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateArchiveReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateArchiveOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreateArchive: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetFile get file API
*/
func (a *Client) GetFile(params *GetFileParams, authInfo runtime.ClientAuthInfoWriter) (*GetFileOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFileParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetFile",
		Method:             "GET",
		PathPattern:        "/api/file/{fid}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetFileReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetFileOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetFile: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
