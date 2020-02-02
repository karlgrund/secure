package version

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	// BuildTime is a time label from when the binary was built.
	BuildTime = "unset"
	// Commit is the git hash from when the binary was built.
	Commit = "unset"
	// Version is the semantic version of the current build.
	Version = "unset"
)

func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "output version information",
		Run: func(cmd *cobra.Command, args []string) {
			printVersion()
		},
	}

	return cmd
}

func printVersion() {
	output := "version:\t%s\nhash:\t\t%s\nbuild_time:\t%s\n"
	fmt.Printf(output, Version, Commit, BuildTime)
}
