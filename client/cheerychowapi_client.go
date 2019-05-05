// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/cheerychow/chow-go-client/client/account"
	"github.com/cheerychow/chow-go-client/client/chow"
	"github.com/cheerychow/chow-go-client/client/fav_recipe"
	"github.com/cheerychow/chow-go-client/client/food"
	"github.com/cheerychow/chow-go-client/client/labels"
	"github.com/cheerychow/chow-go-client/client/nutrition"
	"github.com/cheerychow/chow-go-client/client/nutrition_tips"
	"github.com/cheerychow/chow-go-client/client/parse"
	"github.com/cheerychow/chow-go-client/client/rate_recipe"
	"github.com/cheerychow/chow-go-client/client/rda"
	"github.com/cheerychow/chow-go-client/client/recipe"
	"github.com/cheerychow/chow-go-client/client/recipe_group"
	"github.com/cheerychow/chow-go-client/client/scratch_recipe"
	"github.com/cheerychow/chow-go-client/client/user"
)

// Default cheerychowapi HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "api.cheerychow.com"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"https"}

// NewHTTPClient creates a new cheerychowapi HTTP client.
func NewHTTPClient(formats strfmt.Registry) *Cheerychowapi {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new cheerychowapi HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *Cheerychowapi {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new cheerychowapi client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Cheerychowapi {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(Cheerychowapi)
	cli.Transport = transport

	cli.Account = account.New(transport, formats)

	cli.Chow = chow.New(transport, formats)

	cli.FavRecipe = fav_recipe.New(transport, formats)

	cli.Food = food.New(transport, formats)

	cli.Labels = labels.New(transport, formats)

	cli.Nutrition = nutrition.New(transport, formats)

	cli.NutritionTips = nutrition_tips.New(transport, formats)

	cli.Parse = parse.New(transport, formats)

	cli.RateRecipe = rate_recipe.New(transport, formats)

	cli.Rda = rda.New(transport, formats)

	cli.Recipe = recipe.New(transport, formats)

	cli.RecipeGroup = recipe_group.New(transport, formats)

	cli.ScratchRecipe = scratch_recipe.New(transport, formats)

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

// Cheerychowapi is a client for cheerychowapi
type Cheerychowapi struct {
	Account *account.Client

	Chow *chow.Client

	FavRecipe *fav_recipe.Client

	Food *food.Client

	Labels *labels.Client

	Nutrition *nutrition.Client

	NutritionTips *nutrition_tips.Client

	Parse *parse.Client

	RateRecipe *rate_recipe.Client

	Rda *rda.Client

	Recipe *recipe.Client

	RecipeGroup *recipe_group.Client

	ScratchRecipe *scratch_recipe.Client

	User *user.Client

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *Cheerychowapi) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport

	c.Account.SetTransport(transport)

	c.Chow.SetTransport(transport)

	c.FavRecipe.SetTransport(transport)

	c.Food.SetTransport(transport)

	c.Labels.SetTransport(transport)

	c.Nutrition.SetTransport(transport)

	c.NutritionTips.SetTransport(transport)

	c.Parse.SetTransport(transport)

	c.RateRecipe.SetTransport(transport)

	c.Rda.SetTransport(transport)

	c.Recipe.SetTransport(transport)

	c.RecipeGroup.SetTransport(transport)

	c.ScratchRecipe.SetTransport(transport)

	c.User.SetTransport(transport)

}
