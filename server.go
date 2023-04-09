package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// define and parse command line arguments
	host := flag.String("h", "localhost", "server host")
	port := flag.String("p", "8080", "server port")
	flag.Parse()

	// run server
	err := runServer(*host, *port)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
}

func runServer(host string, port string) error {
	cmd := exec.Command("go", "run", "app/main.go", fmt.Sprintf("-h=%s", host), fmt.Sprintf("-p=%s", port))
	cmd.Dir = "."
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to run server: %s", err)
	}
	return nil
}
