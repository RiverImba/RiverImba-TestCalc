package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var romanToInt = map[string]int{
	"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5,
	"VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
}

var intToRoman = []string{
	"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X",
}


func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Input (or type 'q' to quit): ")

		// Read user input
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		// q to exit
		if input == "q" {
			fmt.Println("Exiting...")
			break
		}

		// Regular fo input
		parts := splitByOperators(input)

		if len(parts) != 3 {
			fmt.Println("Выдача паники, так как формат математической операции не удовлетворяет заданию.")
			continue
		}

		operand1, operator, operand2 := parts[0], parts[1], parts[2]

		isRomanOperands := isRoman(operand1) && isRoman(operand2)
		isArabicOperands := isArabic(operand1) && isArabic(operand2)

		if !isRomanOperands && !isArabicOperands {
    		fmt.Println("Выдача паники, так как используются одновременно разные системы счисления.")
    		continue
		}

		var num1, num2 int
		var err1, err2 error

		if isArabicOperands {
			num1, err1 = strconv.Atoi(operand1)
			num2, err2 = strconv.Atoi(operand2)
		} else if isRomanOperands {
			num1, err1 = romanToIntValue(operand1)
			num2, err2 = romanToIntValue(operand2)
		} else {
			fmt.Println("Выдача паники, так как одно из чисел некорректно.")
			continue
		}

		if err1 != nil || err2 != nil {
			fmt.Println("Выдача паники, так как одно из чисел некорректно.")
			continue
		}

		result := performOperation(num1, num2, operator)

		if isRomanOperands {
			if result <= 0 {
				fmt.Println("Выдача паники, так как в римской системе нет отрицательных чисел или нуля.")
				continue
			}
			if result > 3999 {
				fmt.Println("Число в римском исчисление не может привышать 3999, текущее %s", result)
			}
			fmt.Println("Результат:", intToRomanValue(result))
		} else {
			fmt.Println("Результат:", result)
		}
	}
}

// splitByOperators separate input by regex
func splitByOperators(s string) []string {
	var parts []string
	re := regexp.MustCompile(`\s*(\d+|[IVX]+)\s*([\+\-\*/])\s*(\d+|[IVX]+)\s*`)
	matches := re.FindStringSubmatch(s)

	if len(matches) == 4 {
		parts = append(parts, matches[1])
		parts = append(parts, matches[2])
		parts = append(parts, matches[3])
	}

	return parts
}

// romanToIntValue convert rome value to arabic
func romanToIntValue(s string) (int, error) {
	value, ok := romanToInt[s]
	if !ok {
		return 0, fmt.Errorf("некорректная римская цифра")
	}
	return value, nil
}

// intToRomanValue arabic to rome
func intToRomanValue(num int) string {
	if num < 1 || num > 10 {
		return ""
	}
	return intToRoman[num]
}

// performOperation Evaluating the exoression
func performOperation(a, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		if b == 0 {
			fmt.Println("Ошибка: деление на ноль.")
			os.Exit(1)
		}
		return a / b
	default:
		fmt.Println("Ошибка: неизвестный оператор.")
		os.Exit(1)
		return 0
	}
}

func isRoman(s string) bool {
	_, exists := romanToInt[s]
	return exists
}

func isArabic(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func integerToRoman(number int) string {
	maxRomanNumber := 3999
	if number > maxRomanNumber {
		//fmt.Println("Число %s превышает значениe %s для римского исчисления", number, maxRomanNumber)
		return nil
	}

	// if number < 1 {
	// 	fmt.Println("Число %s меньше единици для римского исчисления", number)
	// 	return nil
	// }

	conversions := []struct {
		value int
		digit string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
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

	var roman strings.Builder
	for _, conversion := range conversions {
		for number >= conversion.value {
			roman.WriteString(conversion.digit)
			number -= conversion.value
		}
	}

	return roman.String()
}
