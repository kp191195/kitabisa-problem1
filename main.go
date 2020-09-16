package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var input string
	var result int64
	valid := false
	result = 0

	for !valid {
		fmt.Printf("Please input equation : ")
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			input = scanner.Text()
		}
		fmt.Println(fmt.Sprintf("doCalc(\"%s\")", input))
		res, err := doCalc(input)

		if err != nil {
			fmt.Println("Input not valid!")
			continue
		}

		valid = true
		result = res
	}

	fmt.Println(fmt.Sprintf("Result : %d", result))
}
