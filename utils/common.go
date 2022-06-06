package utils

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/rocky/go-gnureadline"
)

func HandleInput(cmd string, sep string) (parsedCmd []string) {
	return strings.SplitN(cmd, sep, 2)
}

func Check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func Exit() {
	gnureadline.Rl_reset_terminal("")
	fmt.Println("\nExiting Program")
	os.Exit(0)
}
