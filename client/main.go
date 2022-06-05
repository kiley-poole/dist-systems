package main

import (
	"fmt"
	"strings"

	"github.com/kiley-poole/dist-systems/utils"
	"github.com/rocky/go-gnureadline"
)

const HISTORY_FILE string = "~/my.history"

func main() {
	gnureadline.ReadHistory(HISTORY_FILE)
	gnureadline.StifleHistory(20)
	gnureadline.ReadInitFile("/etc/inputrc")

	fmt.Println("get {key} will fetch the value stored at the provided key")
	fmt.Println("set {key}={value} will store the value stored at the provided key")
	for {
		cmd, err := gnureadline.Readline(fmt.Sprintln("\nEnter your selection: "), true)
		utils.Check(err)

		cmd = strings.TrimSuffix(cmd, "\n")
		s := utils.HandleInput(cmd, " ")
		cmd = strings.ToLower(s[0])

		if cmd == "exit" {
			utils.Exit()
		}

		if cmd != "get" && cmd != "set" {
			fmt.Println("Invalid Command")
			continue
		}
		fmt.Printf("%s\n", cmd)
		gnureadline.WriteHistory(HISTORY_FILE)
	}
}

func printKV(k string, v string) {
	fmt.Printf("\nKey: %s\nValue: %s\n", k, v)
}
