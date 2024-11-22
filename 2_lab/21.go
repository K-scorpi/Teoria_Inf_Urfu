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

	// Проверочная матрица
	checkMatrix := [][]int{
		{1, 0, 1, 0},
		{0, 0, 1, 1},
		{1, 1, 1, 0},
		{1, 1, 0, 1},
		{0, 1, 1, 1},
	}

	// Полная образующая матрица
	parityMatrixFull := hstack(infoMatrix, checkMatrix)

	// Пример правильного кодового слова
	codeword := []int{1, 0, 1, 0, 1, 1, 1, 0, 0} // Правильное кодовое слово

	// Вносим ошибку в один из разрядов
	codewordWithError := make([]int, len(codeword))
	copy(codewordWithError, codeword)
	codewordWithError[3] ^= 1 // Инверсия бита в разряде 3 (внесли ошибку)

	fmt.Println("Исходное кодовое слово:", codeword)
	fmt.Println("Кодовое слово с ошибкой:", codewordWithError)

	// Вычисление синдрома для кодового слова с ошибкой
	fmt.Println("\nПроверка кодового слова с ошибкой:")
	syndrome := checkCodeword(codewordWithError, transpose(parityMatrixFull))
	fmt.Println("Синдром:", syndrome)

	// Поиск ошибки
	fmt.Println("\nПоиск ошибки в кодовом слове:")
	findError(syndrome, parityMatrixFull)
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
	result := make([]int, len(matrix[0])) // Длина результата равна числу столбцов
	for j := range result {               // Проходим по столбцам результата
		for i := range matrix { // Проходим по строкам матрицы
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

// Проверка кодового слова с отладочными сообщениями
func checkCodeword(codeword []int, checkMatrix [][]int) []int {
	fmt.Println("\nПроверяем кодовое слово:", codeword)

	// Транспонированная проверочная матрица
	transposedMatrix := transpose(checkMatrix)
	fmt.Println("\nТранспонированная проверочная матрица:")
	printMatrix(transposedMatrix)

	// Результат умножения перед модулем
	result := matVecMul(transposedMatrix, codeword)
	fmt.Println("\nРезультат умножения перед модулем:", result)

	// Результат умножения после модуля
	modResult := modVec(result, 2)
	fmt.Println("Результат умножения после модуля:", modResult)

	return modResult
}

// Поиск ошибки
func findError(syndrome []int, parityMatrix [][]int) {
	fmt.Println("\nИщем столбец в транспонированной проверочной матрице, совпадающий с синдромом:")
	for i, col := range transpose(parityMatrix) {
		fmt.Printf("Проверяем столбец %d: %v\n", i, col)
		if equalVectors(syndrome, col) {
			fmt.Printf("Ошибка в разряде %d\n", i)
			return
		}
	}
	fmt.Println("Ошибка не обнаружена или она двукратная.")
}

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
