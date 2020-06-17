// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new user API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for user API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	ConsistencyCheckAndResolve(params *ConsistencyCheckAndResolveParams, authInfo runtime.ClientAuthInfoWriter) (*ConsistencyCheckAndResolveOK, error)

	Find(params *FindParams, authInfo runtime.ClientAuthInfoWriter) (*FindOK, error)

	GetMe(params *GetMeParams, authInfo runtime.ClientAuthInfoWriter) (*GetMeOK, error)

	PutAPIUser(params *PutAPIUserParams, authInfo runtime.ClientAuthInfoWriter) (*PutAPIUserOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  ConsistencyCheckAndResolve consistency check and resolve API
*/
func (a *Client) ConsistencyCheckAndResolve(params *ConsistencyCheckAndResolveParams, authInfo runtime.ClientAuthInfoWriter) (*ConsistencyCheckAndResolveOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewConsistencyCheckAndResolveParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ConsistencyCheckAndResolve",
		Method:             "GET",
		PathPattern:        "/api/user/consistency",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ConsistencyCheckAndResolveReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ConsistencyCheckAndResolveOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ConsistencyCheckAndResolve: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  Find find API
*/
func (a *Client) Find(params *FindParams, authInfo runtime.ClientAuthInfoWriter) (*FindOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Find",
		Method:             "GET",
		PathPattern:        "/api/user/find",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &FindReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*FindOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Find: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetMe get me API
*/
func (a *Client) GetMe(params *GetMeParams, authInfo runtime.ClientAuthInfoWriter) (*GetMeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetMeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetMe",
		Method:             "GET",
		PathPattern:        "/api/user/@me",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetMeReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetMeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetMe: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PutAPIUser обновлениеs данных пользователя данные применятся если имя авторизованого пользователя соответствует имени в указанных в запросе данных
*/
func (a *Client) PutAPIUser(params *PutAPIUserParams, authInfo runtime.ClientAuthInfoWriter) (*PutAPIUserOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutAPIUserParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutAPIUser",
		Method:             "PUT",
		PathPattern:        "/api/user",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/*+json", "application/json", "application/json-patch+json", "text/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PutAPIUserReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PutAPIUserOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PutAPIUser: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}