package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/kiley-poole/dist-systems/utils"
)

var memMap = make(map[string]string, 0)

func main() {

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

func getValue(k string) {
	if v, ok := memMap[k]; ok {
		printKV(k, v)
	} else {
		fmt.Println("Key Not Found")
	}
}

func setValue(s string) {
	memMap[kv[0]] = kv[1]
}
