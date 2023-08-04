
package main

import (
	"fmt"
	"os"
)

func main() {
	// Access the environment variable
	envValue := os.Getenv("DRA_ENV")

	// Use the environment variable in the application
	fmt.Println("Received environment variable:", envValue)
}
