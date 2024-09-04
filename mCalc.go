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

// Основная функция
func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Input (or type 'q' to quit): ")

		// Считываем строку ввода
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		// Проверяем, хочет ли пользователь выйти
		if input == "q" {
			fmt.Println("Exiting...")
			break
		}

		// Используем регулярное выражение для разбивки строки на операнды и оператор
		parts := splitByOperators(input)

		if len(parts) != 3 {
			fmt.Println("Выдача паники, так как формат математической операции не удовлетворяет заданию.")
			continue
		}

		operand1, operator, operand2 := parts[0], parts[1], parts[2]

		// Определяем тип чисел (арабские или римские)
		isRomanOperands := isRoman(operand1) && isRoman(operand2)
		isArabicOperands := !isRoman(operand1) && !isRoman(operand2)

		if isRomanOperands && !(isRoman(operand1) || !isRoman(operand2)) {
			fmt.Println("Выдача паники, так как используются одновременно разные системы счисления.")
			continue
		}

		if isArabicOperands && (!isArabic(operand1) || !isArabic(operand2)) {
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
			fmt.Println("Результат:", intToRomanValue(result))
		} else {
			fmt.Println("Результат:", result)
		}
	}
}

// splitByOperators разбивает строку на части по операторам
func splitByOperators(s string) []string {
	var parts []string
	re := regexp.MustCompile(`\s*(\d+|[IVX]+)\s*([\+\-\*/])\s*(\d+|[IVX]+)\s*`)
	matches := re.FindStringSubmatch(s)

	if len(matches) == 4 {
		parts = append(parts, matches[1]) // первый операнд
		parts = append(parts, matches[2]) // оператор
		parts = append(parts, matches[3]) // второй операнд
	}

	return parts
}

// romanToIntValue преобразует римскую цифру в арабское число
func romanToIntValue(s string) (int, error) {
	value, ok := romanToInt[s]
	if !ok {
		return 0, fmt.Errorf("некорректная римская цифра")
	}
	return value, nil
}

// intToRomanValue преобразует арабское число в римскую цифру
func intToRomanValue(num int) string {
	if num < 1 || num > 10 {
		return ""
	}
	return intToRoman[num]
}

// performOperation выполняет арифметическую операцию
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

// isArabic проверяет, является ли строка арабским числом
func isArabic(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}
