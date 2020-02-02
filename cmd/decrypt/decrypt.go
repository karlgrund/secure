package decrypt

import (
	"io/ioutil"
	"os"
	"fmt"

	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
	"github.com/pypl-johan/secure/dec"
)
type decryption struct {
	fileToDecrypt string
	privateKey    string
	secretKey	  string
}

func NewCmd() *cobra.Command {
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
			"",
			"file to decrypt",
		)
	cmd.Flags().
		StringVar(
			&decrypt.privateKey,
			"privateKey",
			os.Getenv("HOME") + "/.ssh/id_rsa",
			"private key used to decrypt",
		)
	cmd.Flags().
		StringVar(
			&decrypt.secretKey,
			"secretKey",
			"",
			"secret key to decrypt",
		)

	cmd.MarkFlagRequired("secretKey")
	cmd.MarkFlagRequired("file")

	return cmd
}

func (e *decryption) run() {
	privateKey, _ := ioutil.ReadFile(e.privateKey)
	secretKey, _ := ioutil.ReadFile(e.secretKey)
	fileToDecrypt, _ := ioutil.ReadFile(e.fileToDecrypt)

	fmt.Println("Enter password: ")
	pkPassword, _ := terminal.ReadPassword(0)

	unecryptedSecret := dec.DecryptUsingPrivateKey(secretKey , privateKey, string(pkPassword))

	clearText := dec.DecryptUsingAsymmetricKey(fileToDecrypt, unecryptedSecret)
	
	writeToFile(clearText, "secret.txt")
}

func writeToFile(data []byte, fileName string) {
	err := ioutil.WriteFile(fileName, data, 0644)
	if err != nil {
		fmt.Errorf("%s", err)
	}
}