package service

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"io"
	"math/big"
	"net/http"
	"rustdesk-api-server-pro/internal/errcode"
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
		return errcode.New(errcode.ERR3013.Code, errcode.ERR3013.Message)
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
		return errcode.New(errcode.ERR3025.Code, errcode.ERR3025.Message)
	}
	if header.Alg != "RS256" {
		return errcode.Errorf(errcode.ERR3026.Code, errcode.ERR3026.Message, header.Alg)
	}
	if metadata == nil || strings.TrimSpace(metadata.JWKSURI) == "" {
		return errcode.New(errcode.ERR3027.Code, errcode.ERR3027.Message)
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
		return nil, errcode.Errorf(errcode.ERR3028.Code, errcode.ERR3028.Message, resp.StatusCode)
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
	return nil, errcode.New(errcode.ERR3029.Code, errcode.ERR3029.Message)
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
		return nil, errcode.New(errcode.ERR3030.Code, errcode.ERR3030.Message)
	}
	e := 0
	for _, b := range eBytes {
		e = e<<8 + int(b)
	}
	if e == 0 {
		return nil, errcode.New(errcode.ERR3031.Code, errcode.ERR3031.Message)
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
