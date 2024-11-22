package main

import (
	"fmt"
)

func main() {
	// Параметры кода (23, 18)
	n := 23 // длина кодового слова
	k := 18 // длина информационной части

	// Создание образующей матрицы
	infoMatrix := identityMatrix(k)
	parityMatrix := randomMatrix(k, n-k)
	generatorMatrix := hstack(infoMatrix, parityMatrix)

	// Формирование проверочной матрицы
	parityCheckMatrix := hstack(transpose(parityMatrix), identityMatrix(n-k))

	// Генерация информационной части
	infoPart := randomVector(k)
	fmt.Println("Информационная часть:", infoPart)

	// Кодирование
	codeword := encode(infoPart, generatorMatrix)
	fmt.Println("Сформированный кодовый вектор:", codeword)

	// Введение ошибок
	codewordWithError := append([]int{}, codeword...)
	codewordWithError[2] = 1 - codewordWithError[2] // Одиночная ошибка
	codewordWithError[7] = 1 - codewordWithError[7] // Вторая ошибка
	fmt.Println("Кодовое слово с ошибками:", codewordWithError)

	// Проверка синдрома
	syndrome := checkCodeword(codewordWithError, parityCheckMatrix)
	fmt.Println("Синдром:", syndrome)

	// Исправление ошибок
	correctedCodeword := correctErrors(codewordWithError, syndrome, parityCheckMatrix)
	fmt.Println("Исправленное кодовое слово:", correctedCodeword)

	// Проверка на корректность после исправления
	finalSyndrome := checkCodeword(correctedCodeword, parityCheckMatrix)
	if isZeroVector(finalSyndrome) {
		fmt.Println("Ошибки исправлены. Кодовое слово корректно.")
	} else {
		fmt.Println("Ошибки не удалось исправить.")
	}
}

// Создание единичной матрицы
func identityMatrix(size int) [][]int {
	matrix := make([][]int, size)
	for i := range matrix {
		matrix[i] = make([]int, size)
		matrix[i][i] = 1
	}
	return matrix
}

// Создание случайной матрицы
func randomMatrix(rows, cols int) [][]int {
	matrix := make([][]int, rows)
	for i := range matrix {
		matrix[i] = make([]int, cols)
		for j := range matrix[i] {
			matrix[i][j] = randomBit()
		}
	}
	return matrix
}

// Создание случайного вектора
func randomVector(size int) []int {
	vector := make([]int, size)
	for i := range vector {
		vector[i] = randomBit()
	}
	return vector
}

// Генерация случайного бита (0 или 1)
func randomBit() int {
	return int(uint8(randomByte()) % 2)
}

// Вспомогательная функция для случайного числа
func randomByte() byte {
	return byte(42) // Для воспроизводимости; можно заменить на случайное значение
}

// Горизонтальная конкатенация двух матриц
func hstack(a, b [][]int) [][]int {
	rows := len(a)
	result := make([][]int, rows)
	for i := range a {
		result[i] = append(a[i], b[i]...)
	}
	return result
}

// Транспонирование матрицы
func transpose(matrix [][]int) [][]int {
	rows := len(matrix)
	cols := len(matrix[0])
	result := make([][]int, cols)
	for i := 0; i < cols; i++ {
		result[i] = make([]int, rows)
		for j := 0; j < rows; j++ {
			result[i][j] = matrix[j][i]
		}
	}
	return result
}

// Кодирование информационного вектора
func encode(infoPart []int, generatorMatrix [][]int) []int {
	return modVec(matVecMul(generatorMatrix, infoPart), 2)
}

// Проверка кодового слова (вычисление синдрома)
func checkCodeword(codeword []int, parityCheckMatrix [][]int) []int {
	return modVec(matVecMul(parityCheckMatrix, codeword), 2)
}

// Исправление ошибок
func correctErrors(codeword, syndrome []int, parityCheckMatrix [][]int) []int {
	// Проверка одиночной ошибки
	for i, col := range transpose(parityCheckMatrix) {
		if equalVectors(syndrome, col) {
			fmt.Printf("Обнаружена одиночная ошибка в разряде %d. Исправляем.\n", i)
			codeword[i] = 1 - codeword[i] // Инвертируем бит
			return codeword
		}
	}

	// Проверка двойной ошибки
	for i, col1 := range transpose(parityCheckMatrix) {
		for j, col2 := range transpose(parityCheckMatrix) {
			if i < j && equalVectors(syndrome, modVec(vecAdd(col1, col2), 2)) {
				fmt.Printf("Обнаружена двойная ошибка в разрядах %d и %d. Исправляем.\n", i, j)
				codeword[i] = 1 - codeword[i]
				codeword[j] = 1 - codeword[j]
				return codeword
			}
		}
	}

	fmt.Println("Ошибка не обнаружена или она превышает допустимые пределы (2 ошибки).")
	return codeword
}

// Умножение матрицы на вектор
func matVecMul(matrix [][]int, vector []int) []int {
	result := make([]int, len(matrix))
	for i := range matrix {
		for j, v := range vector {
			result[i] += matrix[i][j] * v
		}
	}
	return result
}

// Сложение двух векторов поэлементно
func vecAdd(a, b []int) []int {
	result := make([]int, len(a))
	for i := range a {
		result[i] = a[i] + b[i]
	}
	return result
}

// Взятие остатка по модулю для вектора
func modVec(vector []int, mod int) []int {
	for i := range vector {
		vector[i] %= mod
	}
	return vector
}

// Проверка равенства двух векторов
func equalVectors(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// Проверка, является ли вектор нулевым
func isZeroVector(vector []int) bool {
	for _, v := range vector {
		if v != 0 {
			return false
		}
	}
	return true
}
