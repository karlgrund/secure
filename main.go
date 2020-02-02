package main

import (
	"os"

	"github.com/fatih/color"
	"github.com/pypl-johan/secure/cmd/decrypt"
	"github.com/pypl-johan/secure/cmd/encrypt"
	"github.com/pypl-johan/secure/cmd/version"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := NewRootCmd()
	if err := rootCmd.Execute(); err != nil {
		color.Red("%s", err)
		os.Exit(1)
	}
}

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "secure",
		Short: "tool for sharing secrets using public and private keys",
	}

	cmd.AddCommand(version.NewCmd())
	cmd.AddCommand(encrypt.Encrypt())
	cmd.AddCommand(decrypt.Decrypt())

	return cmd
}
