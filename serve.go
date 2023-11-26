package main

import (
	"log"

	"github.com/theartefak/artefak/bootstrap"
)

func main() {
    // Initialize the application kernel
	kernel, err := bootstrap.StartKernel()
	if err != nil {
		log.Fatalf("Error starting kernel: %s", err)
	}

	// Start the Echo server
	err = kernel.Echo.Start(":8080")
	if err != nil {
		log.Fatalf("Error starting Echo server: %s", err)
	}
}
