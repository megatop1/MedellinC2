/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"

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

	l, err := net.Listen("tcp4", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleConnection(c)
		count++
	}
}

//Function to Get List of all ports listeners are listenening over
func checkListenerPorts() {

}
