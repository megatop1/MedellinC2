/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

// interactCmd represents the interact command
var interactCmd = &cobra.Command{
	Use:   "interact",
	Short: "Interact with a particular agent",
	Long:  `Interact and runs commands on a specific agent`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("interact called")
		//sendRemoteCommand(con net.Conn)
	},
}

func sendRemoteCommand(con net.Conn) {
	//Step 1: Ask user to enter in agent UUID (name in the future)
	fmt.Print("Enter Agent UUID: ")
	var uuid string
	fmt.Scanln(&uuid)
	println("UUID: " + uuid)

	//If UID is in Database

	//Step 2: Check if the agent is alive

	//Step 3: Get Information from Agent based on UUID... Remote Port from Agent
	/*
		var remotePort string
		var remoteIP string
		var UUID string
		var hostname string */

	//Step 4: Open up interactive CLI menu with the agent
	l, err := net.Listen("tcp", "0.0.0.0:4444")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	//con, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		netData, err := bufio.NewReader(con).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		if strings.TrimSpace(string(netData)) == "STOP" {
			fmt.Println("Exiting TCP server!")
			return
		}

		fmt.Print("-> ", string(netData))
		t := time.Now()
		myTime := t.Format(time.RFC3339) + "\n"
		con.Write([]byte(myTime))
	}
}

func init() {
	agentCmd.AddCommand(interactCmd)
}
