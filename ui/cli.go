package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	addAmount := addCmd.Float64("amount", 0, "Amount of the transaction")
	addDescription := addCmd.String("desc", "", "Description of the transaction")

	switch os.Args[1] {
	case "add":
		addCmd.Parse(os.Args[2:])
		fmt.Println("Adding a transaction:", *addAmount, *addDescription)
		// Add transaction logic here
	default:
		fmt.Println("Expected 'add' command")
		os.Exit(1)
	}
}
