/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"errors"
	"fmt"
	"net"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/megatop1/MedellinC2/data"
	"github.com/spf13/cobra"
)

//global variables
var count = 0

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a listener",
	Long: `A listener listens for connections from our agents.
	Listeners are responsible for listeneing for the callbacks from our agents.`,
	Run: func(cmd *cobra.Command, args []string) {
		createNewListener()
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

func promptGetSelect(pc promptContent) string {
	protocols := []string{"HTTP", "RPC", "SMB"} // Since we are using the select input mode, we need to give the user some items (protocols) to select from

	//set the initial index to -1 b/c -1 does not exist as an index in the item slice. Keeps prompt open until user selects protocol with a valid index from the array
	index := -1

	var result string
	var err error
	//as long as index value < 0, keep the prompt open
	for index < 0 {
		//give user ability to add their own option
		prompt := promptui.SelectWithAdd{
			Label:    pc.label,
			Items:    protocols,
			AddLabel: "Other",
		}
		// When we run our prompt, it will return an index, a result, and an error
		index, result, err = prompt.Run()
		// If index value = -1, append option to the array
		if index == -1 {
			protocols = append(protocols, result)
		}
	}
	//quick error handling. If we encounter an error, let the user know the prompt has failed
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}
	//if prompt successfuly captures user input
	fmt.Printf("Input: %s\n", result)
	return result
}

func createNewListener() { //function to construct our listener
	namePromptContent := promptContent{ //prompt user to enter a name for the listener
		"Please provide a listener name",
		"Enter a name for your listener: ",
	}
	Name := promptGetInput(namePromptContent) //capture the name as an input from the user

	//promptContent struct for the port
	portPromptContent := promptContent{
		"Please provide a port number",
		fmt.Sprintf("Enter the port number would you like to use for listener %s: ", Name), //pass the name of the listener as an argument
	}
	Port := promptGetInput(portPromptContent) //capture the port number as input from the user

	//promptContent struct for the IP
	IPPromptContent := promptContent{
		"Please provide the IP address for the listener",
		fmt.Sprintf("Enter the IP Address would you like to use for listener %s: ", Name), //pass the name of the listener as an argument
	}
	IP := promptGetInput(IPPromptContent) //capture the port number as input from the user

	protocolPromptContent := promptContent{ //prompt for user to enter in the protocol
		"Please provide a protocol",
		fmt.Sprintf("Select the protocol do you want to use for your listener: ", Name), //pass the listener name as an argument
	}

	Protocol := promptGetSelect(protocolPromptContent) //capture the protocol as an input from the user

	data.InsertListener(Name, Port, IP, Protocol)
	// Generate the listener using data from the user
	// Create the TCP connection to your attackers Netcat
	//var ipAndPortString string = data.GetIP() + ":" + data.GetPort()
	generateListener()
}

func generateListener() {
	//CONCURRENT TCP SERVER (LISTENER)
	var ipAndPortString string = data.GetIP() + ":" + data.GetPort()
	listener, err := net.Listen("tcp4", ipAndPortString) //create listener and listen over tcp protocol over user defined Ip and port
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer listener.Close()
	println("listener created")
}
