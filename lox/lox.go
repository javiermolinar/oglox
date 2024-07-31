package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 2 {
		fmt.Println("Usage: oglox [script]")
		os.Exit(64)
	} else if len(os.Args) == 2 {
		err := runFile(os.Args[1])
		if err != nil {
			fmt.Printf("Error running file: %s", err)
		}
	} else {
		err := runPrompt()
		if err != nil {
			fmt.Printf("Error running prompt: %s", err)
		}
	}
}

func runFile(path string) error {
	scriptContent, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("Error reading script file: %s", err)
	}
	return run(string(scriptContent))
}

func runPrompt() error {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		text, err := reader.ReadString('\n')
		if err != nil {
			return fmt.Errorf("Error reading from stdin %s", err)
		}
		if len(text) == 0 {
			break
		}
		err = run(text)
		if err != nil {
			return err
		}
	}
	return nil
}

func run(source string) error {
	fmt.Println(source)
	return nil
}
