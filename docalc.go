package main

import (
	"strconv"
	"strings"
)

func doCalc(str string) (int64, error) {
	var arrNumber []int64
	var arrOperand []string
	var temp string
	var result int64
	result = 0
	str = strings.Replace(str, " ", "", -1)
	sliceStr := strings.Split(str, "")

	for i := 0; i < len(sliceStr); i++ {
		if sliceStr[i] != "+" && sliceStr[i] != "-" {
			temp += sliceStr[i]
		} else if sliceStr[i] == "+" {
			number, err := strconv.ParseInt(temp, 10, 64)
			if err != nil {
				return 0, err
			}
			arrNumber = append(arrNumber, number)
			arrOperand = append(arrOperand, sliceStr[i])
			temp = ""
		} else if sliceStr[i] == "-" {
			number, err := strconv.ParseInt(temp, 10, 64)
			if err != nil {
				return 0, err
			}
			arrNumber = append(arrNumber, number)
			arrOperand = append(arrOperand, sliceStr[i])
			temp = ""
		}

		if i == len(sliceStr)-1 && temp != "" {
			number, err := strconv.ParseInt(temp, 10, 64)
			if err != nil {
				return 0, err
			}
			arrNumber = append(arrNumber, number)
			temp = ""
		}

	}

	for j := 0; j < len(arrOperand); j++ {
		if arrOperand[j] == "+" {
			if j == 0 {
				result = arrNumber[j] + arrNumber[j+1]
			} else {
				result = result + arrNumber[j+1]
			}
		} else if arrOperand[j] == "-" {
			if j == 0 {
				result = arrNumber[j] - arrNumber[j+1]
			} else {
				result = result - arrNumber[j+1]
			}
		}
	}
	return result, nil
}
