package main

import (
	"fmt"
)

func main() {
	// Информационная часть образующей матрицы кода
	infoMatrix := [][]int{
		{0, 0, 0, 0, 1},
		{0, 0, 0, 1, 0},
		{0, 0, 1, 0, 0},
		{0, 1, 0, 0, 0},
		{1, 0, 0, 0, 0},
	}

	fmt.Println("Матрица кода:")
	printMatrix(infoMatrix)

	// Проверочная матрица
	checkMatrix := [][]int{
		{1, 0, 1, 0},
		{0, 0, 1, 1},
		{1, 1, 1, 0},
		{1, 1, 0, 1},
		{0, 1, 1, 1},
	}

	fmt.Println("\nПроверочная матрица:")
	printMatrix(checkMatrix)

	// Полная образующая матрица
	parityMatrixFull := hstack(infoMatrix, checkMatrix)

	// Информационная часть кода
	infoPart := []int{1, 0, 1, 0, 1}

	// Формирование избыточного кода
	redundantCode := modVec(matVecMul(checkMatrix, infoPart), 2)

	// Полный кодовый вектор
	codeWord := append(infoPart, redundantCode...)
	fmt.Println("\nСформированный кодовый вектор:", codeWord)

	// Проверка кодовой комбинации

	/*
		Синдром ошибки вычисляется как произведение кодового слова
		на транспонированную проверочную матрицу с последующим взятием
		остатка по модулю 2:
	*/
	codeword1 := []int{0, 1, 0, 1, 0, 1, 1, 0, 0}
	codeword2 := []int{1, 0, 0, 1, 0, 1, 1, 1, 0}

	fmt.Println("\nПроверка Комбинации 1:")
	syndrome1 := checkCodeword(codeword1, transpose(parityMatrixFull))
	fmt.Println("Синдром:", syndrome1)

	fmt.Println("\nПроверка Комбинации 2:")
	syndrome2 := checkCodeword(codeword2, transpose(parityMatrixFull))
	fmt.Println("Синдром:", syndrome2)

	// Поиск ошибки
	fmt.Println("\nПоиск ошибки для Комбинации 1:")
	findError(syndrome1, parityMatrixFull)

	fmt.Println("\nПоиск ошибки для Комбинации 2:")
	findError(syndrome2, parityMatrixFull)
}

// Функция для вывода матрицы
func printMatrix(matrix [][]int) {
	for _, row := range matrix {
		fmt.Println(row)
	}
}

// Функция горизонтальной конкатенации матриц
func hstack(a, b [][]int) [][]int {
	result := make([][]int, len(a))
	for i := range a {
		result[i] = append(a[i], b[i]...)
	}
	return result
}

// Умножение матрицы на вектор
func matVecMul(matrix [][]int, vector []int) []int {
	result := make([]int, len(matrix[0]))
	for i := range vector {
		for j := range result {
			result[j] += matrix[i][j] * vector[i]
		}
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

// Транспонирование матрицы
func transpose(matrix [][]int) [][]int {
	rows, cols := len(matrix), len(matrix[0])
	result := make([][]int, cols)
	for i := range result {
		result[i] = make([]int, rows)
		for j := range matrix {
			result[i][j] = matrix[j][i]
		}
	}
	return result
}

// Проверка кодового слова
func checkCodeword(codeword []int, checkMatrix [][]int) []int {
	return modVec(matVecMul(checkMatrix, codeword), 2)
}

// Поиск ошибки
func findError(syndrome []int, parityMatrix [][]int) {
	for i, col := range transpose(parityMatrix) {
		if equalVectors(syndrome, col) {
			fmt.Printf("Ошибка в разряде %d\n", i)
			return
		}
	}
	fmt.Println("Ошибка не обнаружена или она двукратная.")
}

/*
Если синдром ошибки не равен нулю, то
он указывает на столбец проверочной матрицы HTHT,
соответствующий ошибочному разряду
*/
// Проверка на равенство двух векторов
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
