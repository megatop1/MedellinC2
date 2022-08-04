/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/megatop1/MedellinC2/data"
	"github.com/spf13/cobra"
)

// interactCmd represents the interact command
var interactCmd = &cobra.Command{
	Use:   "interact",
	Short: "Interact with a particular agent",
	Long:  `Interact and runs commands on a specific agent`,
	Run: func(cmd *cobra.Command, args []string) {
		remoteCommand()
		/* Code to Get User Input and send over channel  */
	},
}

func remoteCommand() {
	/* Step 1: Ask user to enter in agent UUID (name in the future) */

	/* Step 2: Checkf if UUID is in DB */
	data.GetAgentUUID()
}

func sendRemoteCommand() {
	//Step 1: Ask user to enter in agent UUID (name in the future)
	/*
		fmt.Print("Enter Agent UUID: ")
		var uuid string
		fmt.Scanln(&uuid)
		println("UUID: " + uuid)
	*/
	//If UID is in Database

	//Step 2: Check if the agent is alive

	//Step 3: Get Information from Agent based on UUID... Remote Port from Agent
	/*
		var remotePort string
		var remoteIP string
		var UUID string
		var hostname string */

	//Step 4: Open up interactive CLI menu with the agent
	/*
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
	*/
	//Map UID to a connection

}

func init() {
	agentCmd.AddCommand(interactCmd)
}
