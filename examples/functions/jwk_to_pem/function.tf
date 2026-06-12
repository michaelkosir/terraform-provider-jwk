terraform {
  required_providers {
    jwk = {
      source = "michaelkosir/jwk"
    }
    http = {
      source = "hashicorp/http"
    }
  }
}

provider "jwk" {}

# Example 1: Fetch JWKS from an endpoint and convert all keys to PEM
data "http" "jwks" {
  url = "https://example.com/.well-known/jwks.json"
}

locals {
  jwks_data = jsondecode(data.http.jwks.response_body)

  # Convert each JWK in the keys array to PEM format
  pem_keys = [
    for jwk in local.jwks_data.keys : provider::jwk::to_pem(jsonencode(jwk))
  ]
}

output "all_pem_keys" {
  description = "All public keys from JWKS converted to PEM format"
  value       = local.pem_keys
}

# Example 2: Use with Vault JWT auth backend
# This replaces the external data source pattern with native Terraform
resource "vault_jwt_auth_backend" "k8s" {
  path        = "k8s"
  type        = "jwt"
  description = "JWT auth backend for Kubernetes workloads"

  # Convert JWKS keys to PEM format for Vault
  jwt_validation_pubkeys = [
    for jwk in local.jwks_data.keys : provider::jwk::to_pem(jsonencode(jwk))
  ]
}

# Example 3: Convert a single JWK to PEM
output "single_pem_key" {
  value = provider::jwk::to_pem(jsonencode({
    kty = "RSA"
    kid = "my-key-id"
    n   = "0vx7agoebGcQSuuPiLJXZptN9nndrQmbXEps2aiAFbWhM78LhWx4cbbfAAtVT86zwu1RK7aPFFxuhDR1L6tSoc_BJECPebWKRXjBZCiFV4n3oknjhMstn64tZ_2W-5JsGY4Hc5n9yBXArwl93lqt7_RN5w6Cf0h4QyQ5v-65YGjQR0_FDW2QvzqY368QQMicAtaSqzs8KJZgnYb9c7d0zgdAZHzu6qMQvRL5hajrn1n91CbOpbISD08qNLyrdkt-bFTWhAI4vMQFh6WeZu0fM4lFd2NcRwr3XPksINHaQ-G_xBniIqbw0Ls1jF44-csFCur-kEgU8awapJzKnqDKgw"
    e   = "AQAB"
  }))
}
