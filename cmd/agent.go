/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"

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

func getAgentInfo() {
	//get the hostname
	hostname, err := exec.Command("hostname", "-f").Output() // Agent's hostname
	if err != nil {
		log.Fatal(err)
	}

	hostnameResult := bytes.NewBuffer(hostname).String()
	fmt.Println("Hostname: ", hostnameResult)
	//get the user
	user, err := exec.Command("whoami").Output() // Agent's Username
	userResult := bytes.NewBuffer(user).String()
	fmt.Println("Username: ", userResult)

	//generate UUID
	//Code to randomly generate UUID
	id := uuid.New().String()
	print("UUID: ")
	fmt.Print(id)
	print("\n")

	//get the IP
	ipcmd := "hostname -I"
	ipaddress, err := exec.Command(ipcmd).Output()
	ipResult := bytes.NewBuffer(ipaddress).String()
	fmt.Println("IP Address: ", ipResult)

	generateAgent(id, ipResult, hostnameResult)
}

func generateAgent(UUID, RemoteIP string, Hostname string) {
	data.InsertAgent(UUID, RemoteIP, Hostname)
}

//agent in the background
func agentForeground() {

}

func init() {
	rootCmd.AddCommand(agentCmd)
}
