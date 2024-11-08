package main

import (
	"container/heap"
	"fmt"
	"math"
	"sort"
)

// Структура для хранения символа и его вероятности
type SymbolProb struct {
	Symbol      string
	Probability float64
}

// Реализация кода Шеннона-Фано
func ShannonFano(symbols []SymbolProb) map[string]string {
	sort.Slice(symbols, func(i, j int) bool {
		return symbols[i].Probability > symbols[j].Probability
	})

	// Рекурсивная функция для построения кодов
	var buildCodes func([]SymbolProb, string) map[string]string
	buildCodes = func(symbols []SymbolProb, prefix string) map[string]string {
		// Если остался один символ, возвращаем его с текущим префиксом
		if len(symbols) == 1 {
			return map[string]string{symbols[0].Symbol: prefix}
		}

		// Находим точку для разделения символов на две группы с равной суммой вероятностей
		totalProb := 0.0
		for _, sp := range symbols {
			totalProb += sp.Probability
		}

		accProb := 0.0
		splitIdx := 0
		for i, sp := range symbols {
			accProb += sp.Probability
			if accProb >= totalProb/2 {
				splitIdx = i + 1
				break
			}
		}

		// Рекурсивно строим коды для левой и правой части
		codes := make(map[string]string)
		for k, v := range buildCodes(symbols[:splitIdx], prefix+"0") {
			codes[k] = v
		}
		for k, v := range buildCodes(symbols[splitIdx:], prefix+"1") {
			codes[k] = v
		}
		return codes
	}

	return buildCodes(symbols, "")
}

// Реализация кода Хаффмана
type HuffmanNode struct {
	Symbol      string
	Probability float64
	Left, Right *HuffmanNode
}

type PriorityQueue []*HuffmanNode

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Probability < pq[j].Probability
}
func (pq PriorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*HuffmanNode))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func Huffman(symbols []SymbolProb) map[string]string {
	// Создаем очередь с приоритетом
	pq := &PriorityQueue{}
	heap.Init(pq)
	for _, sp := range symbols {
		heap.Push(pq, &HuffmanNode{Symbol: sp.Symbol, Probability: sp.Probability})
	}

	// Строим дерево Хаффмана
	for pq.Len() > 1 {
		left := heap.Pop(pq).(*HuffmanNode)
		right := heap.Pop(pq).(*HuffmanNode)
		merged := &HuffmanNode{
			Probability: left.Probability + right.Probability,
			Left:        left,
			Right:       right,
		}
		heap.Push(pq, merged)
	}

	// Генерируем коды из дерева
	var buildCodes func(*HuffmanNode, string) map[string]string
	buildCodes = func(node *HuffmanNode, prefix string) map[string]string {
		if node == nil {
			return nil
		}
		if node.Left == nil && node.Right == nil {
			return map[string]string{node.Symbol: prefix}
		}
		codes := make(map[string]string)
		for k, v := range buildCodes(node.Left, prefix+"0") {
			codes[k] = v
		}
		for k, v := range buildCodes(node.Right, prefix+"1") {
			codes[k] = v
		}
		return codes
	}

	root := heap.Pop(pq).(*HuffmanNode)
	return buildCodes(root, "")
}

// Вычисление метрик кодирования
func CalculateMetrics(symbols []SymbolProb, codes map[string]string) (avgLength, entropy, efficiency, compressionRatio float64) {
	// Средняя длина кода
	for _, sp := range symbols {
		avgLength += sp.Probability * float64(len(codes[sp.Symbol]))
	}

	// Энтропия источника
	for _, sp := range symbols {
		if sp.Probability > 0 {
			entropy -= sp.Probability * math.Log2(sp.Probability)
		}
	}

	// Относительная эффективность (энтропия / средняя длина)
	efficiency = (entropy / avgLength) * 100

	// Коэффициент статистического сжатия (сравнение с фиксированной длиной кода)
	fixedLength := math.Ceil(math.Log2(float64(len(symbols))))
	compressionRatio = fixedLength / avgLength

	return
}

func main() {
	// Вероятности символов
	symbols := []SymbolProb{
		{"A", 0.204}, {"B", 0.184}, {"C", 0.176}, {"D", 0.146},
		{"E", 0.134}, {"F", 0.077}, {"G", 0.071}, {"H", 0.008},
		{"I", 0.02},
	}

	// Построение кода Шеннона-Фано
	shannonFanoCodes := ShannonFano(symbols)
	avgLengthSF, entropySF, efficiencySF, compressionRatioSF := CalculateMetrics(symbols, shannonFanoCodes)

	// Построение кода Хаффмана
	huffmanCodes := Huffman(symbols)
	avgLengthH, entropyH, efficiencyH, compressionRatioH := CalculateMetrics(symbols, huffmanCodes)

	// Вывод результатов
	fmt.Println("Коды Шеннона-Фано:", shannonFanoCodes)
	fmt.Printf("Метрики Шеннона-Фано (Средняя длина: %.4f, Энтропия: %.4f, Эффективность: %.2f%%, Коэффициент сжатия: %.4f)\n",
		avgLengthSF, entropySF, efficiencySF, compressionRatioSF)

	fmt.Println("Коды Хаффмана:", huffmanCodes)
	fmt.Printf("Метрики Хаффмана (Средняя длина: %.4f, Энтропия: %.4f, Эффективность: %.2f%%, Коэффициент сжатия: %.4f)\n",
		avgLengthH, entropyH, efficiencyH, compressionRatioH)
}
