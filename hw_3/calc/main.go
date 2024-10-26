package main

import (
	"bufio"
	"fmt"
	"github.com/Benzogang-Tape/BASHNYA-Go/hw_3/calc/calc"
	"os"
	"strings"
)

func main() {
	var expression string
	in := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("Enter math expression(or \"exit\" to exit):\n")
		expression, _ = in.ReadString('\n')
		//expression = strings.TrimSuffix(expression, "\r\n")
		expression = strings.TrimSpace(expression)

		if expression == "exit" {
			break
		}

		result, err := calc.Calc(expression)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(result)
	}
}
