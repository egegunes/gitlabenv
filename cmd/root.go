package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gitlabenv [command]",
	Short: "gitlabenv lets you manage Gitlab CI variables",
	Long:  "gitlabenv lets you manage Gitlab CI variables",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
