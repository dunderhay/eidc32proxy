package eidc32proxy

import (
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	mathrand "math/rand"
	"os"
	"testing"
	"time"
)

func init() {
	mathrand.Seed(time.Now().UnixNano())
}

func TestGetCertFromFile(t *testing.T) {
	cdata := `-----BEGIN CERTIFICATE-----
MIICiDCCAfECEGHeMygE/7ugQumjhkFc+EowDQYJKoZIhvcNAQEFBQAwgYQxFDAS
BgNVBAMMCzN4TE9HSUMgSW5jMREwDwYDVQQKDAhpbmZpbmlhczEVMBMGA1UEBwwM
SW5kaWFuYXBvbGlzMRAwDgYDVQQIDAdJbmRpYW5hMQswCQYDVQQGEwJVUzEjMCEG
CSqGSIb3DQEJARYUc3VwcG9ydEBpbmZpbmlhcy5jb20wHhcNMTkxMDMxMTgxMDE5
WhcNMjQxMDI5MTgxMDE5WjCBhDEUMBIGA1UEAwwLM3hMT0dJQyBJbmMxETAPBgNV
BAoMCGluZmluaWFzMRUwEwYDVQQHDAxJbmRpYW5hcG9saXMxEDAOBgNVBAgMB0lu
ZGlhbmExCzAJBgNVBAYTAlVTMSMwIQYJKoZIhvcNAQkBFhRzdXBwb3J0QGluZmlu
aWFzLmNvbTCBnzANBgkqhkiG9w0BAQEFAAOBjQAwgYkCgYEA6x3fchmUKoTsWj29
5K5LRsoI1jSyUMReaiSdqhrZLQW4JTBAp4QYvJ5Z2uraZ6nuNt9hkcdkibU4NGGb
773+xtHAA0/ljttSZMyYKviEUO2qqVg6GGjElLTiWiGAo1S6rgwKafGyZZvNrz8Y
gi8GRAJCnwaOIlXoGde8+dUPKjUCAwEAATANBgkqhkiG9w0BAQUFAAOBgQApC/K2
w4kgjf2xeIdilv66l7nxvVfYEQvZu+e+JbfRtRPZObcrB9m3FngJEPG5aTUBQO34
9JiriIK4PPQ9y6kY9Pz7sZaJXU/0dyeSZDomKoY3zDH8ttJ7FC3eidhuTBPJ+Ncb
SGgpKjVziunyjPDValZKl9mQXdlnDRO7lPo5lQ==
-----END CERTIFICATE-----`
	certfile, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(certfile.Name())
	err = ioutil.WriteFile(certfile.Name(), []byte(cdata), 0600)
	if err != nil {
		t.Fatal(err)
	}

	_, err = getCertFromFile(certfile.Name())
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetClearKeyFromFile(t *testing.T) {
	kdata := `-----BEGIN RSA PRIVATE KEY-----
MIICXQIBAAKBgQDrHd9yGZQqhOxaPb3krktGygjWNLJQxF5qJJ2qGtktBbglMECn
hBi8nlna6tpnqe4232GRx2SJtTg0YZvvvf7G0cADT+WO21JkzJgq+IRQ7aqpWDoY
aMSUtOJaIYCjVLquDApp8bJlm82vPxiCLwZEAkKfBo4iVegZ17z51Q8qNQIDAQAB
AoGAbvHYt5GkXe/9S5Po4FjygoPhaZrSLdSLrNB8aYFjy5/wRfQf/iwSNCcQxYGe
792637/G3bBWG7kcvXL1z0o7RxS2FsW5UEUSeOJ3ohHltbV5SBd7Non2QURle5EE
yGzgXxuZO/K9snorfW32PizHEUc5wlwVe5AGvelM29TCOKECQQD246dhLhipHyDl
rxFbgYeL67JP4W855wnVo3Brj3KX4utapfyk1dZTxmsxR5/VM6uAUzz5BHUA1X/x
ePLOjrIdAkEA88sBFPxNaCA52lcR72VyveXTcFMuPELU0JqX1hMm1wL0qwPsuaFe
mVgnhh/KMfpps2eXtnhanQgWqINYcY/c+QJAV1Os563TYTa2fyeOXyyQ0kgbOTAH
FJcJHn0CDbmekeTc1KJzm6ZbeiRr0/F+sn3lQq2umnIeJJ5f8/yQ/cjxbQJBANqx
0uimZDHyJrOsw9QDJ2keUAxFMgaw1QPEikxppb/fUOhQfv0OuzPIFryEq/clcciU
N05irLaNWPYVzTMiINECQQDD6QMb0uKlQtLtujeaGsiQ/XCFs1E8NRaA6wU95qK4
HonFEwRcf5Q1WfMyeWHYdnaIRRRMqZIiW+7lv0jcewv+
-----END RSA PRIVATE KEY-----`
	keyfile, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(keyfile.Name())
	err = ioutil.WriteFile(keyfile.Name(), []byte(kdata), 0600)
	if err != nil {
		t.Fatal(err)
	}

	_, err = getKeyFromFile(keyfile.Name(), "")
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetCryptKeyFromFile(t *testing.T) {
	kdata := `-----BEGIN RSA PRIVATE KEY-----
Proc-Type: 4,ENCRYPTED
DEK-Info: AES-256-CBC,7209D131DA3B94ECAE82C5AC4D0D105E

a4cYp3FSQnvh/wCwz7bMBZsu5xtPi6dlc0dsaozwXEg7/g+x/OubJI0bp4bgQwxH
BDUzquYGvXZHaTl03gmpZEoCW5K217tgx1TaEB4hBb1l4IlvJw5XLEWvtknahYHT
tbkLELUMavmCDkjTyMWg7J1aCRoh1OGRXtm1LKCnwqx3DIrzFtnErBvO61Jai2kU
+MqNZt9mwUz1duT5CaWxawZ9DG/758xAqwVYfFUVz967PSKEXks0Lg44Z2v6ntTZ
s+MwdzhJwGjBZJVPJIp5DoroAQWPVRs0iO65mr95MlGCsSxrLuA9cQRW6rtsrVi/
043XxmJ5RPLreJ7wPuWCZJnpSA14DvamO1Ih4MN9M4H6EBv0nA58cuHV+wWvBJ5b
38SmqZvq5x8JnuORQET5O1suzO637DtJRkAnqG61ChvEaF5fM9Nwov5mK/4JWvuB
XWHN8wRtiyOVMvLPFkAMER/mvZ8Vo1npg0lxemy3GSYYkizI4INMbQrOIIWlM7Yb
alt6gJ740x3l2IP9ufR/Ln3oTF8/gY6JHYrgeLWUnGT+aAGwNKFVwqPpYAg1x+V+
2bABhW5jlZm+CneKj6mtwYXy3+UCfjzOLrAvbniwIiS6K/Y0sP/rs5ErxyhOo62Z
opI9l8pjz0FzKSdfFG2fKpmIjqcgF0eclddCwGpDEwiCvtJIpgRTcmD/nbB/Gtd7
0na0JEKmQY40lK/soMlX9dj7eJFR1gMx55YSESIDBOoAD/rQHjFVQVvFBJU7DwjZ
yjDxAmczVu1mZIbgpyJSovRG+P1/qK6C2zBP49u0m7LAlpK5Ah8V+9eYu9cHYnTS
-----END RSA PRIVATE KEY-----`

	keyfile, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(keyfile.Name())
	err = ioutil.WriteFile(keyfile.Name(), []byte(kdata), 0600)
	if err != nil {
		t.Fatal(err)
	}

	_, err = getKeyFromFile(keyfile.Name(), "secret")
	if err != nil {
		t.Fatal(err)
	}
}

func TestCertAndKey(t *testing.T) {
	csGenAll := InfiniasCertSetup()

	cert, key, err := CertAndKey(csGenAll)
	if err != nil {
		t.Fatal(err)
	}

	certFile, err := ioutil.TempFile("", "cert")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(certFile.Name())
	err = pem.Encode(certFile, &pem.Block{Type: "CERTIFICATE", Bytes: cert.Raw})
	if err != nil {
		t.Fatalf("Failed to write data to %s: %s", certFile.Name(), err)
	}
	err = certFile.Close()
	if err != nil {
		t.Fatalf("Error closing %s: %s", certFile.Name(), err)
	}

	clearKeyFile, err := ioutil.TempFile("", "clearkey.pem")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(certFile.Name())
	block := pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(key)}
	err = pem.Encode(clearKeyFile, &block)
	if err != nil {
		t.Fatalf("Failed to write data to %s: %s", clearKeyFile.Name(), err)
	}
	err = clearKeyFile.Close()
	if err != nil {
		t.Fatalf("Error closing %s: %s", clearKeyFile.Name(), err)
	}

	passphrase := randomString(10)
	cryptKeyFile, err := ioutil.TempFile("", "cryptkey.pem")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(cryptKeyFile.Name())
	cipherBlock, err := x509.EncryptPEMBlock(rand.Reader, block.Type, block.Bytes, []byte(passphrase), x509.PEMCipherAES256)
	if err != nil {
		t.Fatal(err)
	}
	err = pem.Encode(cryptKeyFile, &pem.Block{Type: cipherBlock.Type, Headers: cipherBlock.Headers, Bytes: cipherBlock.Bytes})
	if err != nil {
		t.Fatal(err)
	}
	if err := cryptKeyFile.Close(); err != nil {
		t.Fatalf("Error closing %s: %s", cryptKeyFile.Name(), err)
	}

	csClearKey := &CertSetup{
		keyFile:  clearKeyFile.Name(),
		template: csGenAll.template,
	}
	cert, key, err = CertAndKey(csClearKey)
	if err != nil {
		t.Fatal(err)
	}

	csCryptKey := &CertSetup{
		keyFile:    cryptKeyFile.Name(),
		passphrase: passphrase,
		template:   csGenAll.template,
	}
	cert, key, err = CertAndKey(csCryptKey)
	if err != nil {
		t.Fatal(err)
	}

	csClearKeyCert := &CertSetup{
		certFile: certFile.Name(),
		keyFile:  clearKeyFile.Name(),
	}
	cert, key, err = CertAndKey(csClearKeyCert)
	if err != nil {
		t.Fatal(err)
	}

	csCryptKeyCert := &CertSetup{
		certFile:   certFile.Name(),
		keyFile:    cryptKeyFile.Name(),
		passphrase: passphrase,
	}
	cert, key, err = CertAndKey(csCryptKeyCert)
	if err != nil {
		t.Fatal(err)
	}
}

func randomString(n int) string {
	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	b := make([]rune, n)
	for i := range b {
		b[i] = letter[mathrand.Intn(len(letter))]
	}
	return string(b)
}
