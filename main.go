package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

const VERSION = "0.1.0"

var Config string

func handleSignals() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGHUP) // Add syscalls as needed

	// Handle any cleanup
	select {
	case <-ch:
		log.Printf("[INFO] Got a SIGHUP, cleaning up.")
		// Do cleanup here
	}
}

func init() {
	config := flag.String("config", "config.ini", "path to config file")
	version := flag.Bool("v", false, "prints current project version")

	flag.Usage = func() {
		fmt.Printf("Usage %s [OPTIONS] [name ...]\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	if *version {
		fmt.Println(VERSION)
		os.Exit(0)
	}

	// Instead parse and store as global config object
	Config = *config
}

func main() {
	go handleSignals()

	log.Println("[INFO] Starting project")

	log.Printf("[INFO] Using config file: %s", Config)
}
