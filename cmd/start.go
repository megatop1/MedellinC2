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
	fmt.Print(".")
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
	/*	arguments := os.Args
		if len(arguments) == 1 {
			fmt.Println("Please provide a port number!")
			return
		}

		PORT := ":" + arguments[1] */
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
