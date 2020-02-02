package enc

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"fmt"

	"github.com/ssh-vault/ssh2pem"
)

// GenerateKeyAndCipherBlock returns random key with inputed size
// and an aes cipher.block
func GenerateKeyAndCipherBlock(size int) ([]byte, cipher.Block) {
	key := make([]byte, size)
	_, err := rand.Read(key)
	if err != nil {
		fmt.Printf("%s", err)
	}

	generatedKey, _ := aes.NewCipher(key)

	return key, generatedKey
}

// EncryptUsingAsymmetricKey encrypt bytes using asymmetric key
func EncryptUsingAsymmetricKey(toEncrypt []byte, asymmetricKey cipher.Block) []byte {
	aesgcm, err := cipher.NewGCM(asymmetricKey)
	if err != nil {
		panic(err.Error())
	}

	return aesgcm.Seal(nil, make([]byte, 12), toEncrypt, nil)
}

// EncryptUsingPublicKey encrypt using public key
func EncryptUsingPublicKey(toEncrypt, publicKey []byte) []byte {
	decodedPubKey, _ := ssh2pem.DecodePublicKey(string(publicKey))

	encrypted, _ := rsa.EncryptPKCS1v15(rand.Reader, decodedPubKey.(*rsa.PublicKey), toEncrypt)

	return encrypted
}
