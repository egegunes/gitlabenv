package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/xanzy/go-gitlab"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list [project]",
	Short: "List CI variables",
	Long:  "List CI variables. Works with both project id and NAMESPACE/PROJECTNAME",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		git := gitlab.NewClient(nil, viper.GetString("token"))

		pid := args[0]
		variables, _, err := git.ProjectVariables.ListVariables(pid, nil)

		if err != nil {
			fmt.Printf("couldn't get project variables: %v\n", err)
			os.Exit(1)
		}

		for _, variable := range variables {
			fmt.Printf("%s=%s\n", variable.Key, variable.Value)
		}

		os.Exit(0)
	},
}
