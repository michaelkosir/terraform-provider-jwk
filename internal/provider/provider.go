// Copyright (c) Michael Kosir
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/action"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/ephemeral"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure JwkPemProvider satisfies various provider interfaces.
var _ provider.Provider = &JwkPemProvider{}
var _ provider.ProviderWithFunctions = &JwkPemProvider{}
var _ provider.ProviderWithEphemeralResources = &JwkPemProvider{}
var _ provider.ProviderWithActions = &JwkPemProvider{}

// JwkPemProvider defines the provider implementation.
type JwkPemProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// JwkPemProviderModel describes the provider data model.
type JwkPemProviderModel struct {
	Endpoint types.String `tfsdk:"endpoint"`
}

func (p *JwkPemProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "jwkpem"
	resp.Version = p.version
}

func (p *JwkPemProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Provider for converting JWK (JSON Web Key) to PEM format",
		Attributes:          map[string]schema.Attribute{},
	}
}

func (p *JwkPemProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data JwkPemProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// No configuration needed for this provider
}

func (p *JwkPemProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{}
}

func (p *JwkPemProvider) EphemeralResources(ctx context.Context) []func() ephemeral.EphemeralResource {
	return []func() ephemeral.EphemeralResource{}
}

func (p *JwkPemProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{}
}

func (p *JwkPemProvider) Functions(ctx context.Context) []func() function.Function {
	return []func() function.Function{
		NewJwkToPemFunction,
	}
}

func (p *JwkPemProvider) Actions(ctx context.Context) []func() action.Action {
	return []func() action.Action{}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &JwkPemProvider{
			version: version,
		}
	}
}
