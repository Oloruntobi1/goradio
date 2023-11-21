package main

import (
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/neptunsk1y/goradio/cmd"
	"github.com/neptunsk1y/goradio/version"
	"os"
)

func handlePanic() {
	if err := recover(); err != nil {
		log.Error("crashed", "err", err)
		os.Exit(1)
	}
}

func main() {
	_, err := version.Latest()
	if err != nil {
		fmt.Errorf("Error version check")
	}
	cmd.Execute()
	defer handlePanic()
}
