package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

var memMap = make(map[string]string, 0)

func main() {
	go sigHandler()
	fmt.Println("get {key} will fetch the value stored at the provided key")
	fmt.Println("set {key}={value} will store the value stored at the provided key")
	r := bufio.NewReader(os.Stdin)
	f := openFile()
	for {
		fmt.Print("\nEnter your command:")

		cmd, err := r.ReadString('\n')
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

func sigHandler() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT)
	sig := <-sigs
	switch sig {
	case syscall.SIGINT:
		exit()
	}
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func exit() {
	fmt.Println("\nExiting Program")
	os.Exit(0)
}
