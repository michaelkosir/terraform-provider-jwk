// Copyright (c) Michael Kosir
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"math/big"

	"github.com/hashicorp/terraform-plugin-framework/function"
)

var (
	_ function.Function = ToPemFunction{}
)

func NewToPemFunction() function.Function {
	return ToPemFunction{}
}

type ToPemFunction struct{}

// JWK represents a JSON Web Key structure.
type JWK struct {
	Kty string `json:"kty"` // Key Type
	N   string `json:"n"`   // Modulus
	E   string `json:"e"`   // Exponent
	Use string `json:"use,omitempty"`
	Kid string `json:"kid,omitempty"`
	Alg string `json:"alg,omitempty"`
}

func (r ToPemFunction) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "to_pem"
}

func (r ToPemFunction) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:             "Convert JWK to PEM format",
		MarkdownDescription: "Converts a JSON Web Key (JWK) to PEM (Privacy Enhanced Mail) format. Accepts a JWK JSON string and returns the corresponding PEM-encoded public key.",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:                "jwk_json",
				MarkdownDescription: "JSON string containing the JWK (JSON Web Key) to convert. Must include 'n' (modulus) and 'e' (exponent) fields.",
			},
		},
		Return: function.StringReturn{},
	}
}

func (r ToPemFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var jwkJSON string

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &jwkJSON))
	if resp.Error != nil {
		return
	}

	// Parse the JWK JSON
	var jwk JWK
	if err := json.Unmarshal([]byte(jwkJSON), &jwk); err != nil {
		resp.Error = function.ConcatFuncErrors(
			function.NewArgumentFuncError(0, fmt.Sprintf("Invalid JWK JSON: %s", err.Error())),
		)
		return
	}

	// Validate key type
	if jwk.Kty != "RSA" && jwk.Kty != "" {
		resp.Error = function.ConcatFuncErrors(
			function.NewArgumentFuncError(0, fmt.Sprintf("Unsupported key type: %s. Only RSA keys are supported.", jwk.Kty)),
		)
		return
	}

	// Validate required fields
	if jwk.N == "" || jwk.E == "" {
		resp.Error = function.ConcatFuncErrors(
			function.NewArgumentFuncError(0, "JWK must contain 'n' (modulus) and 'e' (exponent) fields"),
		)
		return
	}

	// Convert the JWK to PEM
	pemString, err := convertJWKToPEM(jwk)
	if err != nil {
		resp.Error = function.ConcatFuncErrors(
			function.NewFuncError(fmt.Sprintf("Failed to convert JWK to PEM: %s", err.Error())),
		)
		return
	}

	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, pemString))
}

// convertJWKToPEM converts a JWK to PEM format.
func convertJWKToPEM(jwk JWK) (string, error) {
	// Decode base64url-encoded modulus (n)
	// Add padding if necessary (Python script adds "===" which is more than needed, but works)
	nBytes, err := base64.RawURLEncoding.DecodeString(jwk.N)
	if err != nil {
		return "", fmt.Errorf("failed to decode modulus 'n': %w", err)
	}

	// Decode base64url-encoded exponent (e)
	eBytes, err := base64.RawURLEncoding.DecodeString(jwk.E)
	if err != nil {
		return "", fmt.Errorf("failed to decode exponent 'e': %w", err)
	}

	// Convert bytes to big integers
	n := new(big.Int).SetBytes(nBytes)
	e := new(big.Int).SetBytes(eBytes)

	// Create RSA public key
	pubKey := &rsa.PublicKey{
		N: n,
		E: int(e.Int64()),
	}

	// Marshal the public key to PKIX format (SubjectPublicKeyInfo)
	pubKeyBytes, err := x509.MarshalPKIXPublicKey(pubKey)
	if err != nil {
		return "", fmt.Errorf("failed to marshal public key: %w", err)
	}

	// Encode to PEM format
	pemBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: pubKeyBytes,
	}

	pemBytes := pem.EncodeToMemory(pemBlock)
	return string(pemBytes), nil
}
