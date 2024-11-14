package main

import (
	"fmt"
)

// Матрица образующей
var G = [][]int{
	{1, 0, 1, 0},
	{0, 0, 1, 1},
	{1, 1, 1, 0},
	{1, 1, 0, 1},
	{0, 1, 1, 1},
}

// Функция для вычисления синдрома
func syndrome(codeWord []int) []int {
	var synd []int
	for i := 0; i < len(G); i++ {
		sum := 0
		// Используем все 9 бит codeWord, но ограничены размером строки G[i]
		for j := 0; j < len(G[i]); j++ {
			sum ^= codeWord[j] * G[i][j]
		}
		synd = append(synd, sum)
	}
	return synd
}

// Функция для формирования избыточного кода
func encode(informationWord []int) []int {
	var codeWord []int
	// Сначала добавляем информационную часть
	for i := 0; i < len(informationWord); i++ {
		codeWord = append(codeWord, informationWord[i])
	}
	//  Добавляем избыточные биты, вычисленные с помощью матрицы G
	for i := 0; i < len(G); i++ {
		sum := 0
		for j := 0; j < len(informationWord); j++ {
			sum ^= informationWord[j] * G[i][j]
		}
		codeWord = append(codeWord, sum)
	}
	return codeWord
}

// Функция для проверки кодовой комбинации
func check(codeWord []int) (bool, int) {
	synd := syndrome(codeWord)

	// Проверка, является ли синдром нулевым (все элементы равны нулю)
	isValid := true
	for _, v := range synd {
		if v != 0 {
			isValid = false
			break
		}
	}

	// Если кодовая комбинация верна, возвращаем true
	if isValid {
		return true, -1
	}

	// Поиск позиции первой ошибки (если есть)
	errorPosition := -1
	for i := 0; i < len(synd); i++ {
		if synd[i] == 1 {
			errorPosition = i
			break
		}
	}
	return false, errorPosition
}

func main() {
	// Контрольные соотношения
	//fmt.Println("Контрольные соотношения:")
	/*for i := 0; i < len(G); i++ {
	  //fmt.Print("C" + strconv.Itoa(i+1) + " = ")
	  for j := 0; j < len(G[i]); j++ {
	   if G[i][j] == 1 {
	    //fmt.Print("x" + strconv.Itoa(j+1))
	    if j < len(G[i])-1 {
	     //fmt.Print(" ^ ")
	    }
	   }
	  }
	  //fmt.Println(" = 0")
	 }*/
	// Формирование избыточного кода
	informationWord1 := []int{1, 0, 1, 1}
	codeWord1 := encode(informationWord1)
	fmt.Println("\nИзбыточный код 1:", codeWord1)

	informationWord2 := []int{0, 1, 1, 0}
	codeWord2 := encode(informationWord2)
	fmt.Println("Избыточный код 2:", codeWord2)

	// Проверка кодовых комбинаций
	fmt.Println("\nПроверка кодовых комбинаций:")
	valid, errorPos := check(codeWord1)
	if valid {
		fmt.Println("Кодовая комбинация 1 верна.")
	} else {
		fmt.Println("Кодовая комбинация 1 неверна.")
		fmt.Println("Ошибочный разряд:", errorPos)
	}

	valid, errorPos = check(codeWord2)
	if valid {
		fmt.Println("Кодовая комбинация 2 верна.")
	} else {
		fmt.Println("Кодовая комбинация 2 неверна.")
		fmt.Println("Ошибочный разряд:", errorPos)
	}
}
