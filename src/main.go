package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	version string
)

type config struct {
	port      int
	directory string
	html      string
}

func main() {
	var cfg config
	flag.IntVar(&cfg.port, "port", 9000, "Port number for server to run")
	flag.StringVar(&cfg.directory, "directory", "./web", "Folder name where you exported your project")
	flag.StringVar(&cfg.html, "html", "index.html", "Your main HTML filename")

	showVersion := flag.Bool("version", false, "Display version and exit")
	flag.Parse()

	// This is for testing will be added automatically at build time
	version = "0.0.1"
	if *showVersion {
		fmt.Printf("Version:\t%s\n", version)
		os.Exit(0)
	}

	app := &application{
		config: cfg,
	}

	app.serve()
}
