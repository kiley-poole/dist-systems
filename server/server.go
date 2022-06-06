package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	"github.com/kiley-poole/dist-systems/utils"
)

var memMap = make(map[string]string, 0)

func main() {
	f := openFile()
	l, err := net.Listen("tcp", ":9740")
	utils.Check(err)

	for {
		buildMap(f)
		conn, err := l.Accept()
		if err != nil {
			log.Println(err)
		}
		go handleCommand(conn, f)
	}
}

func handleCommand(conn net.Conn, f *os.File) {
	input, err := bufio.NewReader(conn).ReadString('\n')
	utils.Check(err)

	input = strings.TrimSuffix(input, "\n")
	s := strings.SplitN(input, " ", 2)

	cmd := strings.ToLower(s[0])
	kv := s[1]

	var res string
	if cmd == "get" {
		res = getValue(kv)
	} else {
		res = setValue(kv)
		flush(f)
	}
	fmt.Fprintf(conn, "%s\n", res)
	conn.Close()

}

func buildMap(f *os.File) {
	data, err := os.ReadFile("db")
	utils.Check(err)
	if len(data) > 0 {
		err = json.Unmarshal(data, &memMap)
		utils.Check(err)
	}
}

func flush(f *os.File) {
	data, err := json.Marshal(memMap)
	utils.Check(err)
	os.WriteFile("db", data, 0644)
}

func openFile() *os.File {
	file, err := os.OpenFile("db", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	return file
}

func getValue(k string) string {
	if v, ok := memMap[k]; ok {
		return v
	}
	return "Value Not Found"

}

func setValue(s string) string {
	kv := strings.SplitN(s, "=", 2)
	memMap[strings.TrimSpace(kv[0])] = strings.TrimSpace(kv[1])
	return "Value Set"
}
