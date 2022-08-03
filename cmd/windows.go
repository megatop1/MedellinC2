/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/megatop1/MedellinC2/data"
	"github.com/spf13/cobra"
)

// windowsCmd represents the windows command
var windowsCmd = &cobra.Command{
	Use:   "windows",
	Short: "Displays a list of different windows payloads that can be generated",
	Long: `Please select a windows payload that you wish to generate and deploy on
	the target machine. Please choose wisely and do not get caught`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("windows called")
		createWindowsPayload()
	},
}

type promptWindowsContent struct {
	remoteIP   string
	listener   string
	listenerIP string
	Jitter     string
}

func promptGetWindowsInput(pc promptContent) string {
	validate := func(input string) error {
		if len(input) <= 0 {
			return errors.New(pc.errorMsg)
		}
		return nil
	}

	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }}",
		Valid:   "{{ . | green }}",
		Invalid: "{{ . | red }}",
		Success: "{{ . | bold }}",
	}

	prompt := promptui.Prompt{
		Label:     pc.label,
		Templates: templates,
		Validate:  validate,
	}

	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Input: %s\n", result)
	return result
}

func promptWindowsSelect(pc promptContent) string {
	payloads := []string{"powershell", "batch", "executable"}
	index := -1 //keeps prompt open until user chooses a choice

	var result string
	var err error

	for index < 0 {
		prompt := promptui.SelectWithAdd{
			Label:    pc.label,
			Items:    payloads,
			AddLabel: "Other",
		}
		index, result, err = prompt.Run()
		if index == -1 {
			payloads = append(payloads, result)
		}
	}

	//error handling
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}
	//if user input successfully received then print it
	fmt.Printf("Input: %s\n", result)
	return result
}

func createWindowsPayload() {
	payloadNamePromptContent := promptContent{
		"Please enter in a name for your payload: ",
		"Payload Name: ",
	}
	payloadName := promptGetWindowsInput(payloadNamePromptContent)

	//payload type
	windowsPayloadPromptContent := promptContent{
		"Please choose a payload type",
		fmt.Sprintf("Windows Payloads Types: %s", payloadName),
	}

	payload := promptGetWindowsInput(windowsPayloadPromptContent)

	definitionJitter := promptContent{
		"Please choose a jitter percentage",
		fmt.Sprintf("Enter jitter percent: , %s", payload),
	}

	jitter := promptGetWindowsInput(definitionJitter)

	definitionListenerIP := promptContent{
		"Please enter a listener IP",
		fmt.Sprintf("Please enter the listener IP: %s", jitter),
	}

	listenerIP := promptGetWindowsInput(definitionListenerIP)

	data.InsertLauncher(payloadName, payload, jitter, listenerIP)
}

func init() {
	launcherCmd.AddCommand(windowsCmd)
}
