package main

import (
	"fmt"
)

// Генерация порождающей матрицы (симуляция)
var G = [][]int{
	{1, 0, 0, 0, 0, 0, 0, 1, 0, 1, 1, 0, 1, 1, 1},
	{0, 1, 0, 0, 0, 0, 0, 1, 1, 0, 1, 1, 0, 1, 1},
	{0, 0, 1, 0, 0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 1},
	{0, 0, 0, 1, 0, 0, 0, 1, 1, 1, 1, 0, 1, 0, 0},
	{0, 0, 0, 0, 1, 0, 0, 1, 0, 1, 1, 1, 0, 1, 0},
	{0, 0, 0, 0, 0, 1, 0, 0, 1, 0, 1, 1, 1, 0, 1},
	{0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 1, 1, 1, 1, 0},
}

// Функция кодирования
func encode(informationWord []int) []int {
	codeWord := make([]int, len(G[0]))
	for i := 0; i < len(G); i++ {
		for j := 0; j < len(G[0]); j++ {
			codeWord[j] ^= informationWord[i] * G[i][j]
		}
	}
	return codeWord
}

// Функция вычисления синдрома
func syndrome(codeWord []int) []int {
	var synd []int
	for i := 0; i < len(G); i++ {
		sum := 0
		for j := 0; j < len(G[0]); j++ {
			sum ^= codeWord[j] * G[i][j]
		}
		synd = append(synd, sum)
	}
	return synd
}

// Исправление ошибок (пример для одиночных и двойных)
func correctErrors(codeWord []int) []int {
	synd := syndrome(codeWord)
	// Простая проверка на нулевой синдром
	isZero := true
	for _, s := range synd {
		if s != 0 {
			isZero = false
			break
		}
	}
	if isZero {
		return codeWord
	}

	// Примерная реализация для исправления ошибок:
	// Пройтись по битам, проверить позиции ошибок и исправить, если возможно
	for i := range codeWord {
		testWord := append([]int{}, codeWord...)
		testWord[i] ^= 1
		testSynd := syndrome(testWord)

		// Проверяем, не обнуляет ли ошибка синдром
		if isZeroSyndrome(testSynd) {
			return testWord // Одиночная ошибка исправлена
		}

		// Попытка исправить двойную ошибку
		for j := i + 1; j < len(codeWord); j++ {
			testWord[j] ^= 1
			testSynd2 := syndrome(testWord)
			if isZeroSyndrome(testSynd2) {
				return testWord // Двойная ошибка исправлена
			}
			testWord[j] ^= 1 // Возврат второго бита
		}
	}

	// Если ошибки не исправлены, возвращаем код как есть
	fmt.Println("Не удалось исправить ошибки")
	return codeWord
}

// Функция проверки на нулевой синдром
func isZeroSyndrome(synd []int) bool {
	for _, s := range synd {
		if s != 0 {
			return false
		}
	}
	return true
}

func main() {
	// Пример информационного слова (7 бит)
	informationWord := []int{1, 0, 1, 1, 0, 1, 0}
	fmt.Println("Информационное слово:", informationWord)

	// Кодирование
	codeWord := encode(informationWord)
	fmt.Println("Кодовое слово:", codeWord)

	// Вносим одиночную ошибку
	codeWord[2] ^= 1 // Инвертируем 3-й бит
	fmt.Println("Кодовое слово с ошибкой:", codeWord)

	// Исправляем ошибки
	correctedCode := correctErrors(codeWord)
	fmt.Println("Исправленное кодовое слово:", correctedCode)
}
