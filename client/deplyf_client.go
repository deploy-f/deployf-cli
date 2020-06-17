// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"cli/client/admin_setting"
	"cli/client/application"
	"cli/client/archive_analysis"
	"cli/client/billing"
	"cli/client/docker"
	"cli/client/domains"
	"cli/client/file"
	"cli/client/home"
	"cli/client/notification"
	"cli/client/promocode"
	"cli/client/ready_application"
	"cli/client/resources"
	"cli/client/user"
)

// Default deplyf HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "localhost"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"http"}

// NewHTTPClient creates a new deplyf HTTP client.
func NewHTTPClient(formats strfmt.Registry) *Deplyf {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new deplyf HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *Deplyf {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new deplyf client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Deplyf {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(Deplyf)
	cli.Transport = transport
	cli.AdminSetting = admin_setting.New(transport, formats)
	cli.Application = application.New(transport, formats)
	cli.ArchiveAnalysis = archive_analysis.New(transport, formats)
	cli.Billing = billing.New(transport, formats)
	cli.Docker = docker.New(transport, formats)
	cli.Domains = domains.New(transport, formats)
	cli.File = file.New(transport, formats)
	cli.Home = home.New(transport, formats)
	cli.Notification = notification.New(transport, formats)
	cli.Promocode = promocode.New(transport, formats)
	cli.ReadyApplication = ready_application.New(transport, formats)
	cli.Resources = resources.New(transport, formats)
	cli.User = user.New(transport, formats)
	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// Deplyf is a client for deplyf
type Deplyf struct {
	AdminSetting admin_setting.ClientService

	Application application.ClientService

	ArchiveAnalysis archive_analysis.ClientService

	Billing billing.ClientService

	Docker docker.ClientService

	Domains domains.ClientService

	File file.ClientService

	Home home.ClientService

	Notification notification.ClientService

	Promocode promocode.ClientService

	ReadyApplication ready_application.ClientService

	Resources resources.ClientService

	User user.ClientService

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *Deplyf) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport
	c.AdminSetting.SetTransport(transport)
	c.Application.SetTransport(transport)
	c.ArchiveAnalysis.SetTransport(transport)
	c.Billing.SetTransport(transport)
	c.Docker.SetTransport(transport)
	c.Domains.SetTransport(transport)
	c.File.SetTransport(transport)
	c.Home.SetTransport(transport)
	c.Notification.SetTransport(transport)
	c.Promocode.SetTransport(transport)
	c.ReadyApplication.SetTransport(transport)
	c.Resources.SetTransport(transport)
	c.User.SetTransport(transport)
}