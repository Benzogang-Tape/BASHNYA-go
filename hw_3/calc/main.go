package main

import (
	"fmt"
	"github.com/Benzogang-Tape/BASHNYA-Go/hw_3/calc/calc"
)

func main() {
	var expression string
	for {
		fmt.Printf("Enter math expression(or \"exit\" to exit):\n")
		fmt.Scanln(&expression)

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
