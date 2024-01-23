package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Input:")

	reader := bufio.NewReader(os.Stdin)

	text, _ := reader.ReadString('\n')
	str := strings.Split(text, " ")

	if len(str) > 4 {
		panic("Выдача паники, так как формат математической операции не удовлетворяет заданию - два операнда и один оператор (+, -, /, *)")
	} else if len(str) < 3 {
		panic("Выдача паники, так как строка не является математической операцией")
	}

	userNum1 := str[0]
	symbol := str[1]
	userNum2 := strings.TrimSpace(str[2])

	num1, _ := strconv.Atoi(userNum1)
	num2, _ := strconv.Atoi(userNum2)

	if num1 != 0 && num2 != 0 {
		switch symbol {
		case "+":
			fmt.Println(num1 + num2)
		case "-":
			fmt.Println(num1 - num2)
		case "*":
			fmt.Println(num1 * num2)
		case "/":
			c := num1 / num2
			fmt.Println(c)
		}
	} else if num1 == 0 && num2 == 1 || num1 == 1 && num2 == 0 {
		panic("Выдача паники, так как используются одновременно разные системы счисления")
	} else if num1 == 0 && num2 == 0 {

		var Num1, Num2 int
		var totalNum int
		romanDigits := map[string]int{
			"I":    1,
			"II":   2,
			"III":  3,
			"IV":   4,
			"V":    5,
			"VI":   6,
			"VII":  7,
			"VIII": 8,
			"IX":   9,
			"X":    10,
		}

		for key, value := range romanDigits {
			if key == userNum1 {
				Num1 = value
			}
			if key == userNum2 {
				Num2 = value
			}
		}

		switch symbol {
		case "+":
			totalNum = Num1 + Num2
			fmt.Print(totalRomanNum(totalNum))
		case "-":
			totalNum = Num1 - Num2
			if totalNum > 0 {
				fmt.Println(totalRomanNum(totalNum))
			} else {
				panic("Выдача паники, так как в римской системе нет отрицательных чисел")
			}
		case "*":
			totalNum = Num1 * Num2
			fmt.Println(totalRomanNum(totalNum))
		case "/":
			totalNum = Num1 / Num2
			fmt.Println(totalRomanNum(totalNum))
		}

	}

}

func totalRomanNum(totalNum int) string {
	type romanDigits []struct {
		number int
		romD   string
	}
	var romArr = romanDigits{
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}
	var concatenation strings.Builder
	for _, romDig := range romArr {
		for totalNum >= romDig.number {
			concatenation.WriteString(romDig.romD)
			totalNum -= romDig.number
		}
	}

	return concatenation.String()
}
