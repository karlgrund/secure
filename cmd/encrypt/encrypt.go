package encrypt

import (
	"fmt"
	"io/ioutil"

	"github.com/pypl-johan/secure/enc"
	"github.com/spf13/cobra"
)

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

	err := cmd.MarkFlagRequired("publicKey")
	if err != nil {
		fmt.Printf("%s", err)
	}

	err = cmd.MarkFlagRequired("file")
	if err != nil {
		fmt.Printf("%s", err)
	}

	return cmd
}

func (e *encryption) run() {
	publicKey, _ := ioutil.ReadFile(e.publicKey)
	fileToEncrypt, _ := ioutil.ReadFile(e.fileToEncrypt)

	secretKey, cipherblock := enc.GenerateKeyAndCipherBlock(32)

	ciphertext := enc.EncryptUsingAsymmetricKey(fileToEncrypt, cipherblock)

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
