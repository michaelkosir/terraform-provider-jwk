terraform {
  required_providers {
    jwk = {
      source = "michaelkosir/jwk"
    }
  }
}

# No configuration is required for the jwk provider
provider "jwk" {}
