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

// linuxCmd represents the linux command
var linuxCmd = &cobra.Command{
	Use:   "linux",
	Short: "Displays a list of different linux payloads that can be generated",
	Long: `Please select a linux payload that you wish to generate and deploy on
	the target machine. Please choose wisely and do not get caught`,
	Run: func(cmd *cobra.Command, args []string) {
		createLinuxPayload()
	},
}

type promptLinuxContent struct {
	remoteIP    string
	listener    string
	listenerIP  string
	remotePort  string
	Jitter      string
	payloadType string
}

func promptGetLinuxInput(pc promptContent) string {
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

func createLinuxPayload() {
	payloadNamePromptContent := promptContent{
		"Please enter in the remote IP: ",
		"Remote IP: ",
	}
	remoteIP := promptGetWindowsInput(payloadNamePromptContent)

	//payload type
	linuxPayloadPromptContent := promptContent{
		"Please choose an existing listener type",
		"Listener: ",
	}

	listener := promptGetLinuxInput(linuxPayloadPromptContent)

	definitionListenerIP := promptContent{
		"Please choose a local IP",
		"LHOST: ",
	}

	listenerIP := promptGetLinuxInput(definitionListenerIP)

	definitionRemotePort := promptContent{
		"Please choose a remote port",
		"RPORT: ",
	}

	remotePort := promptGetLinuxInput(definitionRemotePort)

	definitionJitter := promptContent{
		"Please enter a jitter percentage",
		"Please enter desired jitter percentage: %",
	}

	jitter := promptGetLinuxInput(definitionJitter)

	definitionPayloadType := promptContent{
		"Please enter a payload type",
		"Please enter desired payloadtype: ",
	}

	payloadType := promptLinuxSelect(definitionPayloadType)

	data.InsertLauncher(remoteIP, listener, listenerIP, remotePort, remoteIP, jitter)

	if payloadType != "" {
		if payloadType == "bash" {
			print("generating bash launcher...\n")
		} else if payloadType == "executab;e" {
			print("generating executable launcher...\n")
			//powershellLauncher(remoteIP, remotePort)
		} else if payloadType == "python" {
			print("generating python launcher...\n")
		} else if payloadType == "perl" {
			print("generating reverse_ssh launcher...\n")
		} else if payloadType == "socat" {
			print("generating socat launcher...\n")
		}
	}

	print("Launcher created\n")
}

func promptLinuxSelect(pc promptContent) string {
	payloads := []string{"bash", "executable", "python", "perl", "reverse_ssh", "socat"}
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

func bashLauncher() {

}

func executableLauncher() {

}

func pythonLauncher() {

}

func perlLauncher() {

}

func sshLauncher() {

}

func socatLauncher() {

}

func init() {
	launcherCmd.AddCommand(linuxCmd)
}
