// Copyright (c) Michael Kosir
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
)

func TestJwkToPemFunction_Valid(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				output "test" {
					value = provider::jwkpem::jwk_to_pem(jsonencode({
						use = "sig"
						kty = "RSA"
						kid = "bAfiahKKpQMLjvezK3KKPLZuJc48JwmVNHVh9KbHXjE"
						alg = "RS256"
						n   = "mRlYmG_vmIgKnw1dmi5XNfsQEGlOwDLf_sXp7HCfgQLiAzUJRF1zExzeK4oFPG11PhP4Iu56Xcb7QyWatn9QbWRHjpsRjjzDbCNdVSHznsIUTfUVlQO_vCEhFzmqN00JS0zGEm8QlAUm20GcL-ZstqMoHJvLWIrZydGYdYtgKqCl1ob_5WhswezU6s3wRsdapybA07qnREh_PX7KfCmLE3nJB81WGa5FAqDmbSdRdQ641MTP-A3SRYF_4DTCe2wKqIFQ5gwgZqJlF2qIF1_TcVWRm9unmlAWgnAcD12FSRfs7Fv2X-UWf6oJM45gMJBT9KTvxr2ghfD75KHSNXwLIQ"
						e   = "AQAB"
					}))
				}
				`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownOutputValue(
						"test",
						knownvalue.StringRegexp(regexp.MustCompile(`^-----BEGIN PUBLIC KEY-----\n`)),
					),
				},
			},
		},
	})
}

func TestJwkToPemFunction_ValidWithOptionalFields(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				output "test" {
					value = provider::jwkpem::jwk_to_pem(jsonencode({
						use = "sig"
						kty = "RSA"
						kid = "bAfiahKKpQMLjvezK3KKPLZuJc48JwmVNHVh9KbHXjE"
						alg = "RS256"
						n   = "mRlYmG_vmIgKnw1dmi5XNfsQEGlOwDLf_sXp7HCfgQLiAzUJRF1zExzeK4oFPG11PhP4Iu56Xcb7QyWatn9QbWRHjpsRjjzDbCNdVSHznsIUTfUVlQO_vCEhFzmqN00JS0zGEm8QlAUm20GcL-ZstqMoHJvLWIrZydGYdYtgKqCl1ob_5WhswezU6s3wRsdapybA07qnREh_PX7KfCmLE3nJB81WGa5FAqDmbSdRdQ641MTP-A3SRYF_4DTCe2wKqIFQ5gwgZqJlF2qIF1_TcVWRm9unmlAWgnAcD12FSRfs7Fv2X-UWf6oJM45gMJBT9KTvxr2ghfD75KHSNXwLIQ"
						e   = "AQAB"
					}))
				}
				`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownOutputValue(
						"test",
						knownvalue.StringRegexp(regexp.MustCompile(`^-----BEGIN PUBLIC KEY-----\n`)),
					),
				},
			},
		},
	})
}

func TestJwkToPemFunction_InvalidJSON(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				output "test" {
					value = provider::jwkpem::jwk_to_pem("{invalid json}")
				}
				`,
				ExpectError: regexp.MustCompile(`Invalid JWK JSON`),
			},
		},
	})
}

func TestJwkToPemFunction_MissingModulus(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				output "test" {
					value = provider::jwkpem::jwk_to_pem(jsonencode({
						kty = "RSA"
						e   = "AQAB"
					}))
				}
				`,
				ExpectError: regexp.MustCompile(`JWK must contain`),
			},
		},
	})
}

func TestJwkToPemFunction_MissingExponent(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				output "test" {
					value = provider::jwkpem::jwk_to_pem(jsonencode({
						kty = "RSA"
						n   = "0vx7agoebGcQSuuPiLJXZptN9nndrQmbXEps2aiAFbWhM78LhWx4cbbfAAtVT86zwu1RK7aPFFxuhDR1L6tSoc_BJECPebWKRXjBZCiFV4n3oknjhMstn64tZ_2W-5JsGY4Hc5n9yBXArwl93lqt7_RN5w6Cf0h4QyQ5v-65YGjQR0_FDW2QvzqY368QQMicAtaSqzs8KJZgnYb9c7d0zgdAZHzu6qMQvRL5hajrn1n91CbOpbISD08qNLyrdkt-bFTWhAI4vMQFh6WeZu0fM4lFd2NcRwr3XPksINHaQ-G_xBniIqbw0Ls1jF44-csFCur-kEgU8awapJzKnqDKgw"
					}))
				}
				`,
				ExpectError: regexp.MustCompile(`JWK must contain`),
			},
		},
	})
}

func TestJwkToPemFunction_UnsupportedKeyType(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				output "test" {
					value = provider::jwkpem::jwk_to_pem(jsonencode({
						kty = "EC"
						n   = "0vx7agoebGcQSuuPiLJXZptN9nndrQmbXEps2aiAFbWhM78LhWx4cbbfAAtVT86zwu1RK7aPFFxuhDR1L6tSoc_BJECPebWKRXjBZCiFV4n3oknjhMstn64tZ_2W-5JsGY4Hc5n9yBXArwl93lqt7_RN5w6Cf0h4QyQ5v-65YGjQR0_FDW2QvzqY368QQMicAtaSqzs8KJZgnYb9c7d0zgdAZHzu6qMQvRL5hajrn1n91CbOpbISD08qNLyrdkt-bFTWhAI4vMQFh6WeZu0fM4lFd2NcRwr3XPksINHaQ-G_xBniIqbw0Ls1jF44-csFCur-kEgU8awapJzKnqDKgw"
						e   = "AQAB"
					}))
				}
				`,
				ExpectError: regexp.MustCompile(`Unsupported key type: EC`),
			},
		},
	})
}

func TestJwkToPemFunction_InvalidBase64Modulus(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				output "test" {
					value = provider::jwkpem::jwk_to_pem(jsonencode({
						kty = "RSA"
						n   = "not-valid-base64!!!"
						e   = "AQAB"
					}))
				}
				`,
				ExpectError: regexp.MustCompile(`failed to decode modulus 'n'`),
			},
		},
	})
}

func TestJwkToPemFunction_InvalidBase64Exponent(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				output "test" {
					value = provider::jwkpem::jwk_to_pem(jsonencode({
						kty = "RSA"
						n   = "0vx7agoebGcQSuuPiLJXZptN9nndrQmbXEps2aiAFbWhM78LhWx4cbbfAAtVT86zwu1RK7aPFFxuhDR1L6tSoc_BJECPebWKRXjBZCiFV4n3oknjhMstn64tZ_2W-5JsGY4Hc5n9yBXArwl93lqt7_RN5w6Cf0h4QyQ5v-65YGjQR0_FDW2QvzqY368QQMicAtaSqzs8KJZgnYb9c7d0zgdAZHzu6qMQvRL5hajrn1n91CbOpbISD08qNLyrdkt-bFTWhAI4vMQFh6WeZu0fM4lFd2NcRwr3XPksINHaQ-G_xBniIqbw0Ls1jF44-csFCur-kEgU8awapJzKnqDKgw"
						e   = "not-valid-base64!!!"
					}))
				}
				`,
				ExpectError: regexp.MustCompile(`failed to decode exponent 'e'`),
			},
		},
	})
}

func TestJwkToPemFunction_Null(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				output "test" {
					value = provider::jwkpem::jwk_to_pem(null)
				}
				`,
				ExpectError: regexp.MustCompile(`argument must not be null`),
			},
		},
	})
}

func TestJwkToPemFunction_Unknown(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
				resource "terraform_data" "test" {
					input = jsonencode({
						use = "sig"
						kty = "RSA"
						kid = "bAfiahKKpQMLjvezK3KKPLZuJc48JwmVNHVh9KbHXjE"
						alg = "RS256"
						n   = "mRlYmG_vmIgKnw1dmi5XNfsQEGlOwDLf_sXp7HCfgQLiAzUJRF1zExzeK4oFPG11PhP4Iu56Xcb7QyWatn9QbWRHjpsRjjzDbCNdVSHznsIUTfUVlQO_vCEhFzmqN00JS0zGEm8QlAUm20GcL-ZstqMoHJvLWIrZydGYdYtgKqCl1ob_5WhswezU6s3wRsdapybA07qnREh_PX7KfCmLE3nJB81WGa5FAqDmbSdRdQ641MTP-A3SRYF_4DTCe2wKqIFQ5gwgZqJlF2qIF1_TcVWRm9unmlAWgnAcD12FSRfs7Fv2X-UWf6oJM45gMJBT9KTvxr2ghfD75KHSNXwLIQ"
						e   = "AQAB"
					})
				}
				
				output "test" {
					value = provider::jwkpem::jwk_to_pem(terraform_data.test.output)
				}
				`,
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownOutputValue(
						"test",
						knownvalue.StringRegexp(regexp.MustCompile(`^-----BEGIN PUBLIC KEY-----\n`)),
					),
				},
			},
		},
	})
}
