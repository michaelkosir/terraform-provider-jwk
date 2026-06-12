# Terraform Provider: jwk

This Terraform provider enables conversion of JSON Web Keys (JWK) to PEM (Privacy Enhanced Mail) format within Terraform configurations.

## Requirements

- [Terraform](https://developer.hashicorp.com/terraform/downloads) >= 1.8
- [Go](https://golang.org/doc/install) >= 1.25 (for development)

## Using the Provider

```terraform
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

data "http" "jwks" {
  url = "https://kubernetes.example.com:6443/openid/v1/jwks"
}

locals {
  jwks_data = jsondecode(data.http.jwks.response_body)
}

resource "vault_jwt_auth_backend" "k8s" {
  path        = "k8s"
  type        = "jwt"
  description = "JWT auth backend for Kubernetes workloads"

  # Convert JWKS keys to PEM format for Vault
  jwt_validation_pubkeys = [
    for jwk in local.jwks_data.keys : provider::jwk::to_pem(jsonencode(jwk))
  ]
}
```

## Available Functions

### `to_pem`

Converts a JSON Web Key (JWK) to PEM format.

**Signature:** `to_pem(jwk_json string) string`

**Parameters:**

- `jwk_json` (String) - A JSON string containing the JWK with required fields `n` (modulus) and `e` (exponent)

**Returns:** A PEM-encoded public key string

See the [function documentation](docs/functions/to_pem.md) for more details and examples.

## Use Cases

- Converting JWKs from OIDC/OAuth2 discovery endpoints to PEM format
- Processing public keys from JWKS (JSON Web Key Set) endpoints
- Integrating with systems that require PEM-formatted keys
- Certificate and key management in cloud environments

## Development

### Building the Provider

```bash
go build -v
```

### Running Tests

```bash
go test -v ./internal/provider/
```

### Installing Locally for Testing

```bash
# Build the provider
go build -v

# Create local provider directory
mkdir -p ~/.terraform.d/plugins/registry.terraform.io/michaelkosir/jwk/0.1.0/darwin_arm64

# Copy the binary (adjust path for your OS/architecture)
cp terraform-provider-jwk ~/.terraform.d/plugins/registry.terraform.io/michaelkosir/jwk/0.1.0/darwin_arm64/terraform-provider-jwk
```

Then in your Terraform configuration:

```terraform
terraform {
  required_providers {
    jwk = {
      source  = "michaelkosir/jwk"
      version = "0.1.0"
    }
  }
}
```

## Documentation

- [Provider Documentation](docs/index.md)
- [Function: to_pem](docs/functions/to_pem.md)
- [Examples](examples/)

## License

Mozilla Public License v2.0 - see [LICENSE](LICENSE) for details.
