package main

import (
	"fmt"
	"math"
)

// Функция для вычисления энтропии
func entropy(p []float64) float64 {
	ent := 0.0
	for _, pi := range p {
		if pi > 0 {
			ent -= pi * math.Log2(pi)
		}
	}
	return ent
}

// Функция для вычисления маргинальных вероятностей для A
func marginalProbabilities(matrix [][]float64, columnCount int) []float64 {
	marginal := make([]float64, columnCount)
	for i := 0; i < columnCount; i++ {
		for j := 0; j < columnCount; j++ {
			marginal[i] += matrix[i][j]
		}
	}
	return marginal
}

// Функция для вычисления совместной энтропии H(A, B)
func jointEntropy(matrix [][]float64) float64 {
	ent := 0.0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] > 0 {
				ent -= matrix[i][j] * math.Log2(matrix[i][j])
			}
		}
	}
	return ent
}

// Функция для вычисления условной энтропии H(A|B)
func conditionalEntropy(matrix [][]float64, marginalB []float64) float64 {
	condEnt := 0.0
	for j := 0; j < len(marginalB); j++ {
		for i := 0; i < len(matrix); i++ {
			pA_B := matrix[i][j] / marginalB[j]
			if pA_B > 0 {
				condEnt -= matrix[i][j] * math.Log2(pA_B)
			}
		}
	}
	return condEnt
}

func main() {
	// Пример матрицы совместных вероятностей p(A, B)
	matrix := [][]float64{
		{0.75, 0.15, 0.1, 0},
		{0.1, 0.7, 0.05, 0.15},
		{0, 0.1, 0.8, 0.01},
		{0.15, 0.05, 0.05, 0.75},
	}

	// Количество состояний A и B
	stateCount := 4

	// Вычисляем маргинальные вероятности для B (путем суммирования по строкам)
	marginalB := make([]float64, stateCount)
	for j := 0; j < stateCount; j++ {
		for i := 0; i < stateCount; i++ {
			marginalB[j] += matrix[i][j]
		}
	}

	// 1. Вычисление энтропии источника H(A)
	marginalA := marginalProbabilities(matrix, stateCount)
	H_A := entropy(marginalA)
	fmt.Printf("Энтропия источника H(A): %.4f\n", H_A)

	// 2. Вычисление совместной энтропии H(A, B)
	H_A_B := jointEntropy(matrix)
	fmt.Printf("Совместная энтропия H(A, B): %.4f\n", H_A_B)

	// 3. Вычисление условной энтропии H(A|B)
	H_A_B_cond := conditionalEntropy(matrix, marginalB)
	fmt.Printf("Условная энтропия H(A|B): %.4f\n", H_A_B_cond)

	// 4. Вычисление информационных потерь (взаимной информации) I(A;B)
	I_A_B := H_A - H_A_B_cond
	fmt.Printf("Информационные потери I(A;B): %.4f\n", I_A_B)
}
