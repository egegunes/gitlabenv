package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/xanzy/go-gitlab"
)

func init() {
	rootCmd.AddCommand(updateCmd)
}

var updateCmd = &cobra.Command{
	Use:   "update [project] [key] [value]",
	Short: "update CI variable",
	Long:  "update CI variable. Works with both project id and NAMESPACE/PROJECTNAME",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		git := gitlab.NewClient(nil, viper.GetString("token"))

		pid := args[0]
		key := args[1]
		value := args[2]
		protected := false
		scope := "*"

		variableOptions := &gitlab.UpdateVariableOptions{
			Value:            &value,
			Protected:        &protected,
			EnvironmentScope: &scope,
		}

		variable, _, err := git.ProjectVariables.UpdateVariable(pid, key, variableOptions, nil)
		fmt.Fprintf(os.Stderr, "updating %s...", variable.Key)

		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", RED("error"))
			fmt.Fprintf(os.Stderr, "couldn't update variable: %v\n", err)
			os.Exit(1)
		}

		fmt.Fprintf(os.Stderr, "%s\n", GREEN("done"))

		os.Exit(0)
	},
}
