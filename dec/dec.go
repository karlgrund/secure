package dec

import (
	"fmt"

	"encoding/pem"
	"crypto/x509"
	"crypto/cipher"
	"crypto/rsa"
	"crypto/rand"
	"crypto/aes"
)

// DecryptUsingPrivateKey decrypt using private key
func DecryptUsingPrivateKey(toDecrypt []byte, pKey []byte, pkPassword string) ([]byte) {
	var err error

	privateKeyPem, _ := pem.Decode([]byte(string(pKey)))
	var decPrivateKey []byte
	if pkPassword == "" {
		fmt.Println("a")
		decPrivateKey = privateKeyPem.Bytes
	} else {

		fmt.Printf("Type: %T", pkPassword)
		decPrivateKey, _ = x509.DecryptPEMBlock(privateKeyPem, []byte(pkPassword))
	}
		
	var parsedKey interface{}
	if parsedKey, _ = x509.ParsePKCS1PrivateKey(decPrivateKey); err != nil {
		if parsedKey, err = x509.ParsePKCS8PrivateKey(decPrivateKey); err != nil { // note this returns type `interface{}`
			fmt.Printf("Unable to parse RSA private key, generating a temp one: %s", err)
		}
	}

	var privateKey *rsa.PrivateKey
	var ok bool
	privateKey, ok = parsedKey.(*rsa.PrivateKey)
	if !ok {
		fmt.Printf("Unable to parse RSA private key, generating a temp one: %s", err)
	}

	unecryptedSecret, _ := rsa.DecryptPKCS1v15(rand.Reader, privateKey, toDecrypt)

	return unecryptedSecret
}

//
func DecryptUsingAsymmetricKey(toDecrypt []byte, asymmetricKey []byte) ([]byte) {
	secKey, _ := aes.NewCipher(asymmetricKey)
	aesgcm2, _ := cipher.NewGCM(secKey)
	clearText, _ := aesgcm2.Open(nil, make([]byte, 12), toDecrypt, nil)
	
	return clearText
}