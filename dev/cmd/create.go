/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a listener",
	Long: `A listener listens for connections from our agents.

Listeners are responsible for listeneing for the callbacks from our agents. `,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create called")
	},
}

type promptContent struct {
	errorMsg string
	label    string
}

func init() {
	listenersCmd.AddCommand(createCmd)
}

func promptGetInput(pc promptContent) string {
	validate := func(input string) error { //validate function to ensure the user enters input
		if len(input) <= 0 {
			return errors.New(pc.errorMsg)
		}
		return nil
	}

	templates := &promptui.PromptTemplates{ //template will style different parts of the prompt
		Prompt:  "{{ . }}",
		Valid:   "{{ . | green }}",
		Invalid: "{{ . | red }}",
		Success: "{{ . | bold }}",
	}

	prompt := promptui.Prompt{ //combine the templates and validate function to determine behavior of input prompt
		Label:     pc.label,
		Templates: templates,
		Validate:  validate,
	}

	result, err := prompt.Run() //now that config is done we can run our prompt. Run method returns result from the user and an error
	if err != nil {             // quick error handling
		fmt.Printf("Prompt Failed %v\n", err) //let user know the prompt has failed
		os.Exit(1)
	}

	fmt.Printf("Input: %s\n", result)
	return result
}
