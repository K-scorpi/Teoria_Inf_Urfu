package main

import (
	"fmt"
)

// Функция для деления с остатком
func divide(dividend, divisor []int) []int {
	n := len(divisor)
	temp := make([]int, len(dividend))
	copy(temp, dividend)

	for i := 0; i <= len(dividend)-n; i++ {
		if temp[i] == 1 {
			for j := 0; j < n; j++ {
				temp[i+j] ^= divisor[j]
			}
		}
	}
	return temp[len(dividend)-n+1:]
}

// Функция для генерации кода
func generateCode(data, generator []int) []int {
	n := len(generator) - 1
	extendedData := append(data, make([]int, n)...)
	remainder := divide(extendedData, generator)
	return append(data, remainder...)
}

// Функция для проверки корректности кода
func checkCode(code, generator []int) []int {
	return divide(code, generator)
}

// Функция для поиска однократных и двукратных ошибок
func findErrors(code, generator []int) {
	remainder := checkCode(code, generator)
	isZero := true
	for _, r := range remainder {
		if r != 0 {
			isZero = false
			break
		}
	}

	if isZero {
		fmt.Println("Код корректен.")
		return
	}

	fmt.Println("Код ошибочный. Остаток:", remainder)

	// Поиск однократных ошибок
	fmt.Println("Проверка на однократные ошибки...")
	for i := 0; i < len(code); i++ {
		corrected := make([]int, len(code))
		copy(corrected, code)
		corrected[i] ^= 1
		if len(checkCode(corrected, generator)) == len(remainder) {
			fmt.Printf("Ошибка в разряде %d может быть причиной.\n", i)
			break
		}
	}

	// Пример двукратных ошибок
	fmt.Println("Примеры двукратных ошибок:")
	count := 0
	for i := 0; i < len(code); i++ {
		for j := i + 1; j < len(code); j++ {
			corrected := make([]int, len(code))
			copy(corrected, code)
			corrected[i] ^= 1
			corrected[j] ^= 1
			if len(checkCode(corrected, generator)) == len(remainder) {
				fmt.Printf("Ошибочные разряды: %d и %d\n", i, j)
				count++
				if count >= 3 {
					return
				}
			}
		}
	}
}

func main() {
	// Входные данные
	data := []int{1, 0, 0, 1, 1, 1, 0, 1, 0}                // 1 0 0 1 1 1 0 1 0
	generator := []int{1, 0, 0, 1, 1, 1}                    // 1 0 0 1 1 1
	code := []int{1, 0, 0, 1, 0, 1, 1, 0, 1, 1, 1, 1, 0, 1} // 1 0 0 1 0 1 1 0 1 1 1 1 0 1

	// Генерация избыточного кода
	generatedCode := generateCode(data, generator)
	fmt.Println("Сгенерированный избыточный код:", generatedCode)

	// Проверка кода
	findErrors(code, generator)
}
