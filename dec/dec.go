package dec

import (
	"fmt"
	"strings"

	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
)

// DecryptUsingPrivateKey decrypt using private key
func DecryptUsingPrivateKey(toDecrypt, pKey []byte) []byte {
	var privateKey *rsa.PrivateKey
	if strings.Contains(string(pKey), "OPENSSH") {
		pk, _ := ssh.ParseRawPrivateKey(pKey)
		privateKey = pk.(*rsa.PrivateKey)
	} else {
		pkPassword := getPkPassword()

		privateKeyPem, _ := pem.Decode(pKey)
		decPrivateKey, _ := x509.DecryptPEMBlock(privateKeyPem, []byte(pkPassword))

		privateKey, _ = x509.ParsePKCS1PrivateKey(decPrivateKey)
	}
	unecryptedSecret, _ := rsa.DecryptPKCS1v15(rand.Reader, privateKey, toDecrypt)

	return unecryptedSecret
}

// DecryptUsingAsymmetricKey is decrypting the inbyte bytearray using the asymmetric key
func DecryptUsingAsymmetricKey(toDecrypt, asymmetricKey []byte) []byte {
	secKey, _ := aes.NewCipher(asymmetricKey)
	aesgcm2, _ := cipher.NewGCM(secKey)
	clearText, _ := aesgcm2.Open(nil, make([]byte, 12), toDecrypt, nil)

	return clearText
}

// getPkPassword asks the user to enter the password for their private key.
func getPkPassword() string {
	fmt.Println("Enter password: ")
	pkPassword, _ := terminal.ReadPassword(0)
	return string(pkPassword)
}
