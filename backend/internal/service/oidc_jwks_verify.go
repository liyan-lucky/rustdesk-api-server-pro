package service

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"strings"
	"time"
)

type oidcIDTokenHeader struct {
	Alg string `json:"alg"`
	Kid string `json:"kid"`
	Typ string `json:"typ"`
}

type oidcJWKSet struct {
	Keys []oidcJWK `json:"keys"`
}

type oidcJWK struct {
	Kty string `json:"kty"`
	Use string `json:"use"`
	Kid string `json:"kid"`
	Alg string `json:"alg"`
	N   string `json:"n"`
	E   string `json:"e"`
}

func (s *OIDCAuthService) verifyIDTokenSignature(idToken string, metadata *oidcMetadata) error {
	parts := strings.Split(idToken, ".")
	if len(parts) != 3 {
		return errors.New("invalid id token")
	}

	headerBytes, err := base64.RawURLEncoding.DecodeString(parts[0])
	if err != nil {
		return err
	}
	var header oidcIDTokenHeader
	if err = json.Unmarshal(headerBytes, &header); err != nil {
		return err
	}
	if strings.TrimSpace(header.Kid) == "" {
		return errors.New("id token kid missing")
	}
	if header.Alg != "RS256" {
		return fmt.Errorf("unsupported id token alg %s", header.Alg)
	}
	if metadata == nil || strings.TrimSpace(metadata.JWKSURI) == "" {
		return errors.New("oidc jwks_uri missing")
	}

	key, err := s.fetchRS256JWK(metadata.JWKSURI, header.Kid)
	if err != nil {
		return err
	}

	signature, err := base64.RawURLEncoding.DecodeString(parts[2])
	if err != nil {
		return err
	}
	signed := []byte(parts[0] + "." + parts[1])
	digest := sha256.Sum256(signed)
	return rsa.VerifyPKCS1v15(key, crypto.SHA256, digest[:], signature)
}

func (s *OIDCAuthService) fetchRS256JWK(jwksURL, kid string) (*rsa.PublicKey, error) {
	req, _ := http.NewRequest(http.MethodGet, jwksURL, nil)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "rustdesk-api-server-pro")

	client := http.DefaultClient
	if s != nil && s.httpClient != nil {
		client = s.httpClient
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("oidc jwks fetch failed with status %d", resp.StatusCode)
	}

	body, err := io.ReadAll(io.LimitReader(resp.Body, 1024*1024))
	if err != nil {
		return nil, err
	}
	var set oidcJWKSet
	if err = json.Unmarshal(body, &set); err != nil {
		return nil, err
	}
	for _, key := range set.Keys {
		if key.Kid != kid || key.Kty != "RSA" {
			continue
		}
		if key.Use != "" && key.Use != "sig" {
			continue
		}
		if key.Alg != "" && key.Alg != "RS256" {
			continue
		}
		publicKey, err := rsaPublicKeyFromJWK(key)
		if err != nil {
			return nil, err
		}
		return publicKey, nil
	}
	return nil, errors.New("matching oidc jwk not found")
}

func rsaPublicKeyFromJWK(key oidcJWK) (*rsa.PublicKey, error) {
	nBytes, err := base64.RawURLEncoding.DecodeString(key.N)
	if err != nil {
		return nil, err
	}
	eBytes, err := base64.RawURLEncoding.DecodeString(key.E)
	if err != nil {
		return nil, err
	}
	if len(nBytes) == 0 || len(eBytes) == 0 {
		return nil, errors.New("invalid rsa jwk")
	}
	e := 0
	for _, b := range eBytes {
		e = e<<8 + int(b)
	}
	if e == 0 {
		return nil, errors.New("invalid rsa exponent")
	}
	return &rsa.PublicKey{N: new(big.Int).SetBytes(nBytes), E: e}, nil
}

func idTokenIssuedTooFarInFuture(claims map[string]interface{}, maxSkew time.Duration) bool {
	iat, ok := claimUnixTime(claims["iat"])
	if !ok {
		return false
	}
	return time.Unix(iat, 0).After(time.Now().Add(maxSkew))
}
