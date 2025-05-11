package main

import (
	"fmt"
	"os"

	"go-test-tutorial/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
