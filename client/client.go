package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"

	"github.com/kiley-poole/dist-systems/utils"
	"github.com/rocky/go-gnureadline"
)

const HISTORY_FILE string = ""

func main() {
	gnureadline.ReadHistory(HISTORY_FILE)
	gnureadline.StifleHistory(20)
	gnureadline.ReadInitFile("/etc/inputrc")

	fmt.Println("get {key} will fetch the value stored at the provided key")
	fmt.Println("set {key}={value} will store the value stored at the provided key")
	for {
		cmd, err := gnureadline.Readline(fmt.Sprintln("\nEnter your selection: "), true)
		utils.Check(err)

		s := strings.SplitN(cmd, " ", 2)
		cmdChk := strings.ToLower(s[0])

		if cmdChk == "exit" {
			gnureadline.WriteHistory(HISTORY_FILE)
			utils.Exit()
		}

		if cmdChk != "get" && cmdChk != "set" {
			fmt.Println("Invalid Command")
			continue
		}

		sendCommand(cmd)
	}
}

func sendCommand(s string) {
	conn, err := net.Dial("tcp", "localhost:9740")
	utils.Check(err)

	fmt.Fprintln(conn, s)

	line, err := bufio.NewReader(conn).ReadString('\n')
	utils.Check(err)

	fmt.Printf("\n%s", string(line))
	conn.Close()
}
