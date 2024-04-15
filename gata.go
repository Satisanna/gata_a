package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var romanMap = []struct {
	decVal int
	symbol string
}{
	{100, "C"}, {90, "XC"}, {50, "L"}, {40, "XL"},
	{10, "X"}, {9, "IX"}, {5, "V"}, {4, "IV"}, {1, "I"},
	{3, "III"}, {2, "II"}, {6, "VI"}, {7, "VII"}, {8, "VIII"},
}

func decimalToRomanIterative(num int) string {
	result := ""
	for _, pair := range romanMap {
		for num >= pair.decVal {
			result += pair.symbol
			num -= pair.decVal
		}
	}
	return result
}

func CheckingForNumbers(n string) bool {
	for i := 0; i < len(n); i++ {
		if unicode.IsDigit([]rune(n)[i]) != true {
			return false
		}
	}
	return true
}

func Numbers(s string) int {
	number, _ := strconv.Atoi(s)
	return number
}

func String(st int) string {
	line := strconv.Itoa(st)
	return line
}

func CheakingRomanNumerals(crn string) int {
	number := 0
	for _, c := range crn {
		k2 := strconv.Itoa(number)
		for _, r := range romanMap {
			if r.symbol == string(c) {
				for i := len(k2) - 1; i < len(k2); i++ {
					if string(k2[i]) == "1" && r.symbol != "I" {
						number -= 2
					}
				}
				number += r.decVal
			}
		}
	}
	return number
}

func MathematicalCalculations(m string) bool {
	switch m {
	case "+":
		return true
	case "-":
		return true
	case "*":
		return true
	case "/":
		return true
	default:
		return false
	}
}

func StringsCount(sc string) bool {
	count := 0
	for i := 0; i < len(sc); i++ {
		if string(sc[i]) == "+" || string(sc[i]) == "-" || string(sc[i]) == "*" || string(sc[i]) == "/" {
			count++
		}
	}
	if count == 1 {
		return false
	}
	return true
}

func CheckingForARangeArNum(cr string) bool {
	if Numbers(cr) < 1 || Numbers(cr) > 10 {
		return false
	}
	return true
}

func CheckingForARangeRomNum(cr string) bool {
	if CheakingRomanNumerals(cr) < 0 || CheakingRomanNumerals(cr) > 10 {
		return false
	}
	return true
}

func AnswerArabicNumerals(n1, s, n2 string) string {
	var answer string
	switch s {
	case "+":
		answer1 := Numbers(n1) + Numbers(n2)
		answer = String(answer1)
	case "-":
		answer2 := Numbers(n1) - Numbers(n2)
		answer = String(answer2)
	case "*":
		answer3 := Numbers(n1) * Numbers(n2)
		answer = String(answer3)
	case "/":
		answer4 := Numbers(n1) / Numbers(n2)
		if Numbers(n2) != 0 {
			answer = String(answer4)
		}
	}
	return answer
}
func CheakingForErrors(code int) error {
	return errors.New(CheakingForCorrect[code])
}

var CheakingForCorrect = map[int]string{
	1: "Выдача паники, так как используются неизвестные системы счисления!",
	2: "Выдача паники, так как используются одновременно разные системы счисления!",
	3: "Выдача паники, так как используются неизвестные математические действия!",
	4: "Выдача паники, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)!",
	5: "Выдача паники, так как в римской системе нет отрицательных чисел!",
	6: "Выдача паники, число не удовлетворяет диапозону от 1 до 10!",
	7: "Выдача паники, так как строка не является математической операцией!",
}

func AnswerRomanNumbers(n1 string, s string, n2 string) string {
	var answer string
	switch s {
	case "+":
		answer1 := CheakingRomanNumerals(n1) + CheakingRomanNumerals(n2)
		answer = decimalToRomanIterative(answer1)
	case "-":
		answer2 := CheakingRomanNumerals(n1) - CheakingRomanNumerals(n2)
		if answer2 >= 1 {
			answer = decimalToRomanIterative(answer2)
		} else if answer2 == 0 {
			answer = "Выдача паники, так как в римской системе нет числа  '0'!"
		} else {
			answer = "Выдача паники, так как в римской системе нет отрицательных чисел!"
		}
	case "*":
		answer3 := CheakingRomanNumerals(n1) * CheakingRomanNumerals(n2)
		answer = decimalToRomanIterative(answer3)
	case "/":
		answer4 := CheakingRomanNumerals(n1) / CheakingRomanNumerals(n2)
		if float32(CheakingRomanNumerals(n1))/float32(CheakingRomanNumerals(n2)) >= 1 {
			answer = decimalToRomanIterative(int(answer4))
		} else {
			answer = "Выдача паники, так как значение меньше 1!"
		}
	}
	return answer
}

func DataInput() (string, string, string, error) {
	data, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	check := strings.Split(strings.Trim(data, "\n"), " ")
	if len(check) < 3 {
		return "", "", "", CheakingForErrors(7)
	}
	check[2] = check[2][:len(check[2])-1]
	if len(check) > 3 && StringsCount(check[1]) == false {
		return "", "", "", CheakingForErrors(4)
	}
	if MathematicalCalculations(check[1]) == false {
		return "", "", "", CheakingForErrors(3)
	}
	if (CheckingForNumbers(check[0]) == true && CheakingRomanNumerals(check[2]) != 0) || (CheckingForNumbers(check[2]) == true && CheakingRomanNumerals(check[0]) != 0) {
		return "", "", "", CheakingForErrors(2)
	}
	if (CheckingForARangeArNum(check[0]) == false || CheckingForARangeArNum(check[2]) == false) && CheckingForNumbers(check[0]) == true && CheckingForNumbers(check[2]) == true {
		return "", "", "", CheakingForErrors(6)
	}
	if (CheckingForARangeRomNum(check[0]) == false || CheckingForARangeRomNum(check[2]) == false) && CheakingRomanNumerals(check[0]) != 0 && CheakingRomanNumerals(check[2]) != 0 {
		return "", "", "", CheakingForErrors(6)
	}
	if (CheckingForNumbers(check[0]) != true || CheckingForNumbers(check[2]) != true) && (CheakingRomanNumerals(check[0]) == 0 || CheakingRomanNumerals(check[2]) == 0) {
		return "", "", "", CheakingForErrors(1)
	}
	return check[0], check[1], check[2], nil
}

func main() {
	number1, sign, number2, err := DataInput()

	if err != nil {
		fmt.Println(err)
		return
	} else if CheckingForNumbers(number1) == true && MathematicalCalculations(sign) == true && CheckingForNumbers(number2) == true {
		fmt.Println(AnswerArabicNumerals(number1, sign, number2))
		return
	} else if CheakingRomanNumerals(number1) != 0 && MathematicalCalculations(sign) == true && CheakingRomanNumerals(number2) != 0 {
		fmt.Println(AnswerRomanNumbers(number1, sign, number2))
		return
	}

}
