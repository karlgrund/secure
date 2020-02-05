package decrypt

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/pypl-johan/secure/dec"
	"github.com/spf13/cobra"
)

type decryption struct {
	fileToDecrypt string
	privateKey    string
	secretKey     string
}

// Decrypt allows decryption of symmetric key using private key
// and will decrypt the secret of file using the unencrypted
// symmetric key.
func Decrypt() *cobra.Command {
	decrypt := decryption{}

	cmd := &cobra.Command{
		Use:   "decrypt file using private key",
		Short: "decrypt file using private key",
		Long:  "decrypt file using private key",
		Run: func(cmd *cobra.Command, args []string) {
			decrypt.run()
		},
	}

	cmd.Flags().
		StringVar(
			&decrypt.fileToDecrypt,
			"file",
			"secret.txt.enc",
			"file to decrypt",
		)
	cmd.Flags().
		StringVar(
			&decrypt.privateKey,
			"privateKey",
			os.Getenv("HOME")+"/.ssh/id_rsa",
			"private key used to decrypt",
		)
	cmd.Flags().
		StringVar(
			&decrypt.secretKey,
			"secretKey",
			"secret.key.enc",
			"secret key to decrypt",
		)

	return cmd
}

// run executes the command to decrypt the files
func (e *decryption) run() {
	privateKey, _ := ioutil.ReadFile(e.privateKey)
	secretKey, _ := ioutil.ReadFile(e.secretKey)
	fileToDecrypt, _ := ioutil.ReadFile(e.fileToDecrypt)

	unecryptedSecret := dec.DecryptUsingPrivateKey(secretKey, privateKey)

	clearText := dec.DecryptUsingAsymmetricKey(fileToDecrypt, unecryptedSecret)

	writeToFile(clearText, "secret.txt")
}

// writeToFile writes the data to a file with name fileName
func writeToFile(data []byte, fileName string) {
	err := ioutil.WriteFile(fileName, data, 0644)
	if err != nil {
		fmt.Printf("%s", err)
	}
}
