package encrypt

import (
	"fmt"
	"io/ioutil"

	"github.com/pypl-johan/secure/enc"
	"github.com/spf13/cobra"
)

//var publicKey = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQDFf4Ol/RCDrdX2iLkfUcJKqQ0qDtrlLhgszNzqnHK53mvZ+8hkN9gPey97MSo8okT9bF7hkpE4tNKuCA8t/+qoFz8MMIOo4J30r2frsQbmdGhLq4uVjHUDEZCCfwbcWjG/+QkMl43n+mWrbqeCAWD2p10a+mUud4gq5L5a3OG/k0jNkwKh7gbF7xRiw3v3k5WwnOmARPe70UbGo7Db6NXXsZFf54aeE05jVWQHNZPVAN5WXzqVbzEKxI2Eyy/yx3nzCCZTh03l/uFNCCmLrZnnT7YZ0sPABbPgkbWLfraBFvmoI9TTLZSA56gOx35qRdgtA9fE0Kqn1gZ6uVevMo/Gu3ACSkdrHszaNjbxtbSiDfirAoYL7rzTuWgsXi1hHE4yAU5xdtU63mF0Eeus3VCbdp/JDTbiCOJBoL0dIvGRel1Oq3NsgoUEBQ+85PKjElahwE4Ll0vnE83Z+lI8zJjF+27x3vEZDBmtzIbTq4HniaLgS+6Fi6CpQMsAukGp8CD6xMXg7HxXkVna1+Kuy0SzF7w8/AsvFYPQ/ZK3/IuXGyDQug/qh3Vc38xr2XHQek0KEwTxBC61/080/SuSlbjMBBR15DpjszU5jP+Ukx+ddwgwuVFn2TgsvEO50exmAiHIsnXSM1zGi/LKp+9yUuqlc+ERFvHU7X/fULAo6ClcFw== johan.karlgrund@izettle.com"

type encryption struct {
	fileToEncrypt string
	publicKey     string
}

// Encrypt allow encryption of secret using assymetric key and
// public key. Outputs two files, encrypted key and encrypted file
func Encrypt() *cobra.Command {
	encrypt := encryption{}

	cmd := &cobra.Command{
		Use:   "encrypt file using public key",
		Short: "encrypt file using public key",
		Long:  "encrypt file using public key",
		Run: func(cmd *cobra.Command, args []string) {
			encrypt.run()
		},
	}

	cmd.Flags().
		StringVar(
			&encrypt.fileToEncrypt,
			"file",
			"",
			"file to encrypt",
		)
	cmd.Flags().
		StringVar(
			&encrypt.publicKey,
			"publicKey",
			"",
			"public key used to encrypt",
		)

	cmd.MarkFlagRequired("publicKey")
	cmd.MarkFlagRequired("file")

	return cmd
}

func (e *encryption) run() {
	publicKey, _ := ioutil.ReadFile(e.publicKey)
	fileToEncrypt, _ := ioutil.ReadFile(e.fileToEncrypt)

	secretKey, cipherblock := enc.GenerateKeyAndCipherBlock(32)

	ciphertext := enc.EncryptUsingAsymmetricKey([]byte(fileToEncrypt), cipherblock)

	encSecretKey := enc.EncryptUsingPublicKey(secretKey, publicKey)

	writeToFile(ciphertext, "secret.txt.enc")
	writeToFile(encSecretKey, "secret.key.enc")
}

func writeToFile(data []byte, fileName string) {
	err := ioutil.WriteFile(fileName, data, 0644)
	if err != nil {
		fmt.Printf("%s", err)
	}
}
