package enc
 
import (
	"crypto/aes"
	"crypto/rand"
	"crypto/cipher"
	"crypto/rsa"

	"github.com/ssh-vault/ssh2pem"
)

// GenerateSecret retuurns random key with inputed size
func GenerateSecret(size int) ([]byte, cipher.Block) {
	key := make([]byte, size)
	rand.Read(key)

	generatedKey, _ := aes.NewCipher(key)

	return key, generatedKey
}

// EncryptUsingAsymmetricKey encrypt bytes using asymmetric key
func EncryptUsingAsymmetricKey(toEncrypt []byte, asymmetricKey cipher.Block) ([]byte) {
	aesgcm, err := cipher.NewGCM(asymmetricKey)
	if err != nil {
		panic(err.Error())
	}

	return aesgcm.Seal(nil, make([]byte, 12), toEncrypt, nil)
}

// EncryptUsingPublicKey encrypt using public key
func EncryptUsingPublicKey(toEncrypt []byte, publicKey []byte) ([]byte) {
	decodedPubKey, _ := ssh2pem.DecodePublicKey(string(publicKey))

	encrypted, _ := rsa.EncryptPKCS1v15(rand.Reader, decodedPubKey.(*rsa.PublicKey), toEncrypt)

	return encrypted
}
