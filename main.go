package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/rocky/go-gnureadline"
)

var memMap = make(map[string]string, 0)

const HISTORY_FILE string = "my.history"

func main() {
	gnureadline.ReadHistory(HISTORY_FILE)
	gnureadline.StifleHistory(20)
	gnureadline.ReadInitFile("/etc/inputrc")

	fmt.Println("get {key} will fetch the value stored at the provided key")
	fmt.Println("set {key}={value} will store the value stored at the provided key")
	f := openFile()
	for {
		cmd, err := gnureadline.Readline(fmt.Sprintln("\nEnter your selection: "), true)
		check(err)

		cmd = strings.TrimSuffix(cmd, "\n")
		s := handleInput(cmd, " ")
		cmd = strings.ToLower(s[0])

		if cmd == "exit" {
			exit()
		}

		if cmd != "get" && cmd != "set" {
			fmt.Println("Invalid Command")
			continue
		}

		buildMap(f)

		if cmd == "get" {
			getValue(s[1])
		} else {
			setValue(s[1])
			flush(f)
		}
		gnureadline.WriteHistory(HISTORY_FILE)
	}
}

func getValue(k string) {
	if v, ok := memMap[k]; ok {
		printKV(k, v)
	} else {
		fmt.Println("Key Not Found")
	}
}

func setValue(s string) {
	kv := handleInput(s, "=")
	memMap[kv[0]] = kv[1]
	printKV(kv[0], kv[1])
}

func buildMap(f *os.File) {
	data, err := os.ReadFile("db")
	check(err)
	if len(data) > 0 {
		err = json.Unmarshal(data, &memMap)
		check(err)
	}
}

func flush(f *os.File) {
	data, err := json.Marshal(memMap)
	check(err)
	os.WriteFile("db", data, 0644)
}

func openFile() *os.File {
	file, err := os.OpenFile("db", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	return file
}

func printKV(k string, v string) {
	fmt.Printf("\nKey: %s\nValue: %s\n", k, v)
}

func handleInput(cmd string, sep string) (parsedCmd []string) {
	return strings.Split(cmd, sep)
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func exit() {
	gnureadline.Rl_reset_terminal("")
	fmt.Println("\nExiting Program")
	os.Exit(0)
}
