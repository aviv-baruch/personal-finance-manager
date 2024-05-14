package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/aviv-baruch/personal-finance-manager/pkg/finance"
	"github.com/aviv-baruch/personal-finance-manager/ui"
)

func main() {
	fm := finance.NewFinanceManagerImpl()
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to Personal Finance Manager. Type 'exit' to quit.")
	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			continue
		}

		input = strings.TrimSpace(input)
		if input == "exit" {
			break
		}

		if err := ui.HandleCommand(input, fm); err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
		}
	}
}
