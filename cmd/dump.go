package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/xanzy/go-gitlab"
)

func init() {
	rootCmd.AddCommand(dumpCmd)
}

var dumpCmd = &cobra.Command{
	Use:   "dump [project]",
	Short: "Dumps CI variables to stdout",
	Long:  "Dumps CI variables to stdout. Works with both project ID and NAMESPACE/PROJECTNAME.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		git := gitlab.NewClient(nil, viper.GetString("token"))

		pid := args[0]
		variables, _, err := git.ProjectVariables.ListVariables(pid, nil)
		if err != nil {
			fmt.Printf("couldn't get project variables: %v\n", err)
			os.Exit(1)
		}

		dump, err := json.MarshalIndent(variables, "", "  ")
		if err != nil {
			fmt.Printf("couldn't dump project variables: %v", err)
			os.Exit(1)
		}

		fmt.Println(string(dump))

		os.Exit(0)
	},
}
