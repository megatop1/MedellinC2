/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"net"
	"os"
	"os/user"

	"github.com/google/uuid"
	"github.com/megatop1/MedellinC2/data"
	"github.com/spf13/cobra"
)

// agentCmd represents the agent command
var agentCmd = &cobra.Command{
	Use:   "agent",
	Short: "Interact with agents",
	Long: `Agents are compromised machines that have an established callback to the C2 server. Use the various agents commands to interact with your agents, and run commands. 
	You are in charge of your "houses" in your city of medellin!
	`,
}

func GetHostIP() (hostIP net.IP) {
	conn, err := net.Dial("udp", "172.23.252.10:80")
	if err != nil {
		os.Exit(1)
	}
	defer conn.Close()

	// We only want the IP, not "IP:port"
	hostIP = conn.LocalAddr().(*net.UDPAddr).IP

	return
}

/* Slice of socket linked to UID */
type ConnDetails struct {
	connection net.Conn
	UID        string
}

func getAgentInfoAndGenerateAgent() {
	//get the hostname
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Hostname: %s", hostname)
	print("\n")

	//get the user
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Username: %s", user.Username)
	print("\n")

	//generate UUID
	//Code to randomly generate UUID
	id := uuid.New().String()
	print("UUID: ")
	fmt.Print(id)
	print("\n")

	//get the IP
	ipResult := GetHostIP().String()
	print("IP Address: " + ipResult)
	print("\n")

	//If UUID Exists then DO NOT recreate it
	// Look for repeating value in the UUID column

	generateAgent(id, ipResult, hostname)
}

func generateAgent(UUID, RemoteIP string, Hostname string) {
	data.InsertAgent(UUID, RemoteIP, Hostname)
}

func awaitCommands() {
	/* for (DefaultDelayValue) { { */
	/* Loop through every row based off of UUID in the DB */
	/* Checks DefaultDelay value in Agent*/
	/* Check Command section in Agent table for that UUID */
	/* Send the command to the server */

	/* Step 1: Loop through every UUID in the database */

}

func listAliveAgents() {

}

//check if agent is alive or not
func checkAgentHealth() {
	/* For each UUID, send a ping. If response is seen or not, change value in DB to Y/N */

	/* For each callback, send a timestamp of the last command was ran and the callback time. Timestamp saved on server. If callback time the agent said is greater than timestamp last time the agent was found*/
}

//agent in the background SHAMELESSLY STOLEN FROM CHRISTIAN
func agentForeground() {

}

func init() {
	rootCmd.AddCommand(agentCmd)
}

/* TESTING */
