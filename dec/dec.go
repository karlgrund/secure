package dec

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
)

// DecryptUsingPrivateKey decrypt using private key
func DecryptUsingPrivateKey(toDecrypt []byte, pKey []byte, pkPassword string) []byte {
	//var err error

	privateKeyPem, _ := pem.Decode([]byte(string(pKey)))
	var decPrivateKey []byte
	if pkPassword == "" {
		decPrivateKey = privateKeyPem.Bytes
	} else {
		decPrivateKey, _ = x509.DecryptPEMBlock(privateKeyPem, []byte(pkPassword))
	}

	privateKey, _ := x509.ParsePKCS1PrivateKey(decPrivateKey)

	unecryptedSecret, _ := rsa.DecryptPKCS1v15(rand.Reader, privateKey, toDecrypt)

	return unecryptedSecret
}

// DecryptUsingAsymmetricKey is decrypting the inbyte bytearray using the asymmetric key
func DecryptUsingAsymmetricKey(toDecrypt []byte, asymmetricKey []byte) []byte {
	secKey, _ := aes.NewCipher(asymmetricKey)
	aesgcm2, _ := cipher.NewGCM(secKey)
	clearText, _ := aesgcm2.Open(nil, make([]byte, 12), toDecrypt, nil)

	return clearText
}
