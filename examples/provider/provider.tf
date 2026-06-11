terraform {
  required_providers {
    jwkpem = {
      source = "michaelkosir/jwkpem"
    }
  }
}

# No configuration is required for the jwkpem provider
provider "jwkpem" {}
