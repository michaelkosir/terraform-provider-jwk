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

// Ensure JwkProvider satisfies various provider interfaces.
var _ provider.Provider = &JwkProvider{}
var _ provider.ProviderWithFunctions = &JwkProvider{}
var _ provider.ProviderWithEphemeralResources = &JwkProvider{}
var _ provider.ProviderWithActions = &JwkProvider{}

// JwkProvider defines the provider implementation.
type JwkProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// JwkProviderModel describes the provider data model.
type JwkProviderModel struct {
	Endpoint types.String `tfsdk:"endpoint"`
}

func (p *JwkProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "jwk"
	resp.Version = p.version
}

func (p *JwkProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Provider for converting JWK (JSON Web Key) to PEM format",
		Attributes:          map[string]schema.Attribute{},
	}
}

func (p *JwkProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data JwkProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// No configuration needed for this provider
}

func (p *JwkProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{}
}

func (p *JwkProvider) EphemeralResources(ctx context.Context) []func() ephemeral.EphemeralResource {
	return []func() ephemeral.EphemeralResource{}
}

func (p *JwkProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{}
}

func (p *JwkProvider) Functions(ctx context.Context) []func() function.Function {
	return []func() function.Function{
		NewToPemFunction,
	}
}

func (p *JwkProvider) Actions(ctx context.Context) []func() action.Action {
	return []func() action.Action{}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &JwkProvider{
			version: version,
		}
	}
}
