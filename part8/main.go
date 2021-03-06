// part8 project main.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"part8/interpreter"
	"strings"
)

func main() {
	fmt.Println("Let's Build A Simple Interpreter - Part 8")

	interdivter := interpreter.New()
	reader := bufio.NewReader(os.Stdin)
	for {
		if s, err := reader.ReadString('\n'); err == nil {
			fmt.Println(interdivter.Interpret(strings.TrimSpace(s)))
			continue
		}
		break
	}
}
