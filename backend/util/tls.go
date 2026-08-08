package util

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"net"
	"os"
	"path/filepath"
	"time"
)

// GenerateSelfSignedCert 生成自签名 ECC 证书和私钥，保存到指定路径
// hosts 参数指定证书覆盖的域名/IP列表，为空时使用 localhost
// 返回 certFile, keyFile, error
func GenerateSelfSignedCert(certFile, keyFile string, hosts []string) (string, string, error) {
	if len(hosts) == 0 {
		hosts = []string{"localhost", "127.0.0.1"}
	}

	priv, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return "", "", err
	}

	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		return "", "", err
	}

	dnsNames := make([]string, 0, len(hosts))
	ipAddrs := make([]net.IP, 0, len(hosts))
	for _, h := range hosts {
		if ip := net.ParseIP(h); ip != nil {
			ipAddrs = append(ipAddrs, ip)
		} else {
			dnsNames = append(dnsNames, h)
		}
	}
	if len(dnsNames) == 0 {
		dnsNames = []string{"localhost"}
	}

	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			Organization: []string{"RustDesk API Server Pro"},
			CommonName:   dnsNames[0],
		},
		NotBefore:             time.Now().Add(-time.Hour),
		NotAfter:              time.Now().AddDate(10, 0, 0),
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
		DNSNames:              dnsNames,
		IPAddresses:           ipAddrs,
	}

	derBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &priv.PublicKey, priv)
	if err != nil {
		return "", "", err
	}

	if err := os.MkdirAll(filepath.Dir(certFile), 0755); err != nil {
		return "", "", err
	}

	certOut, err := os.Create(certFile)
	if err != nil {
		return "", "", err
	}
	defer certOut.Close()
	if err := pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes}); err != nil {
		return "", "", err
	}

	keyOut, err := os.Create(keyFile)
	if err != nil {
		return "", "", err
	}
	defer keyOut.Close()
	privBytes, err := x509.MarshalPKCS8PrivateKey(priv)
	if err != nil {
		return "", "", err
	}
	if err := pem.Encode(keyOut, &pem.Block{Type: "PRIVATE KEY", Bytes: privBytes}); err != nil {
		return "", "", err
	}

	return certFile, keyFile, nil
}
