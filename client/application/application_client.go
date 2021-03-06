// Code generated by go-swagger; DO NOT EDIT.

package application

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new application API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for application API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateApp creates an application
*/
func (a *Client) CreateApp(params *CreateAppParams, authInfo runtime.ClientAuthInfoWriter) (*CreateAppOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateAppParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createApp",
		Method:             "POST",
		PathPattern:        "/application",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateAppReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateAppOK), nil

}

/*
DeleteApp deletes an application
*/
func (a *Client) DeleteApp(params *DeleteAppParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteAppOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteAppParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteApp",
		Method:             "DELETE",
		PathPattern:        "/application/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteAppReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteAppOK), nil

}

/*
GetApps returns all applications
*/
func (a *Client) GetApps(params *GetAppsParams, authInfo runtime.ClientAuthInfoWriter) (*GetAppsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAppsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getApps",
		Method:             "GET",
		PathPattern:        "/application",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetAppsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetAppsOK), nil

}

/*
UpdateApplication updates an application
*/
func (a *Client) UpdateApplication(params *UpdateApplicationParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateApplicationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateApplicationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateApplication",
		Method:             "PUT",
		PathPattern:        "/application/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateApplicationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateApplicationOK), nil

}

/*
UploadAppImage uploads an image for an application
*/
func (a *Client) UploadAppImage(params *UploadAppImageParams, authInfo runtime.ClientAuthInfoWriter) (*UploadAppImageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUploadAppImageParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "uploadAppImage",
		Method:             "POST",
		PathPattern:        "/application/{id}/image",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"multipart/form-data"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UploadAppImageReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UploadAppImageOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
