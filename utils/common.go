package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/rocky/go-gnureadline"
)

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
