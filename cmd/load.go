package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/xanzy/go-gitlab"
)

func init() {
	rootCmd.AddCommand(loadCmd)
}

var loadCmd = &cobra.Command{
	Use:   "load [project] [file]",
	Short: "Loads CI variables to Gitlab",
	Long:  "Loads project level CI variables to Gitlab. Works with both project ID and NAMESPACE/PROJECTNAME.",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		git := gitlab.NewClient(nil, viper.GetString("token"))

		pid := args[0]
		inputFile := args[1]

		content, err := ioutil.ReadFile(inputFile)
		if err != nil {
			fmt.Printf("couldn't read input file: %v\n", err)
			os.Exit(1)
		}

		var variables []gitlab.ProjectVariable
		if err = json.Unmarshal(content, &variables); err != nil {
			fmt.Printf("couldn't load file contents: %v\n", err)
			os.Exit(1)
		}

		for _, variable := range variables {
			fmt.Printf("updating %s... ", variable.Key)
			_, _, err := git.ProjectVariables.UpdateVariable(
				pid,
				variable.Key,
				&gitlab.UpdateVariableOptions{
					&variable.Value,
					&variable.Protected,
					&variable.EnvironmentScope,
				},
				nil,
			)

			if err != nil {
				fmt.Printf("\n%s not found, creating... ", variable.Key)
				_, _, err := git.ProjectVariables.CreateVariable(
					pid,
					&gitlab.CreateVariableOptions{
						&variable.Key,
						&variable.Value,
						&variable.Protected,
						&variable.EnvironmentScope,
					},
					nil,
				)

				if err != nil {
					fmt.Printf("%s\n", RED("error"))
					fmt.Printf("couldn't create variable %s: %v\n", variable.Key, err)
					os.Exit(1)
				}
			}

			fmt.Printf("%s\n", GREEN("done"))
		}

		os.Exit(0)
	},
}
