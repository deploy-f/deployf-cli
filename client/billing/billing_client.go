// Code generated by go-swagger; DO NOT EDIT.

package billing

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new billing API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for billing API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetAllPlans(params *GetAllPlansParams, authInfo runtime.ClientAuthInfoWriter) (*GetAllPlansOK, error)

	GetPage(params *GetPageParams, authInfo runtime.ClientAuthInfoWriter) (*GetPageOK, error)

	ModifyUser(params *ModifyUserParams, authInfo runtime.ClientAuthInfoWriter) (*ModifyUserOK, error)

	RobokassaCallback(params *RobokassaCallbackParams, authInfo runtime.ClientAuthInfoWriter) (*RobokassaCallbackOK, error)

	RobokassaSignature(params *RobokassaSignatureParams, authInfo runtime.ClientAuthInfoWriter) (*RobokassaSignatureOK, error)

	SetPlan(params *SetPlanParams, authInfo runtime.ClientAuthInfoWriter) (*SetPlanOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetAllPlans get all plans API
*/
func (a *Client) GetAllPlans(params *GetAllPlansParams, authInfo runtime.ClientAuthInfoWriter) (*GetAllPlansOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAllPlansParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetAllPlans",
		Method:             "GET",
		PathPattern:        "/api/billing/plans",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAllPlansReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAllPlansOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAllPlans: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetPage get page API
*/
func (a *Client) GetPage(params *GetPageParams, authInfo runtime.ClientAuthInfoWriter) (*GetPageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPageParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPage",
		Method:             "GET",
		PathPattern:        "/api/billing/{userId}/audit",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetPageReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetPageOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetPage: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ModifyUser modify user API
*/
func (a *Client) ModifyUser(params *ModifyUserParams, authInfo runtime.ClientAuthInfoWriter) (*ModifyUserOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewModifyUserParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ModifyUser",
		Method:             "PUT",
		PathPattern:        "/api/billing/{userId}/modify-balance/{amount}",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ModifyUserReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ModifyUserOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ModifyUser: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  RobokassaCallback robokassa callback API
*/
func (a *Client) RobokassaCallback(params *RobokassaCallbackParams, authInfo runtime.ClientAuthInfoWriter) (*RobokassaCallbackOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRobokassaCallbackParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "RobokassaCallback",
		Method:             "GET",
		PathPattern:        "/api/billing/robokassa-callback",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &RobokassaCallbackReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RobokassaCallbackOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for RobokassaCallback: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  RobokassaSignature robokassa signature API
*/
func (a *Client) RobokassaSignature(params *RobokassaSignatureParams, authInfo runtime.ClientAuthInfoWriter) (*RobokassaSignatureOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRobokassaSignatureParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "RobokassaSignature",
		Method:             "GET",
		PathPattern:        "/api/billing/robokassa-signature/{outSum}",
		ProducesMediaTypes: []string{"application/json", "text/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &RobokassaSignatureReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RobokassaSignatureOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for RobokassaSignature: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SetPlan set plan API
*/
func (a *Client) SetPlan(params *SetPlanParams, authInfo runtime.ClientAuthInfoWriter) (*SetPlanOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSetPlanParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "SetPlan",
		Method:             "PUT",
		PathPattern:        "/api/billing/{userId}/set-plan/{planName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SetPlanReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SetPlanOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for SetPlan: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
