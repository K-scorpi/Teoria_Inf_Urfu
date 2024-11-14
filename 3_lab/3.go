package main

import (
	"bytes"
	"compress/lzw"
	"fmt"
	"io"
	"os"
)

// Compress сжимает данные с использованием алгоритма LZW.
func Compress(input []byte) ([]byte, error) {
	var buf bytes.Buffer
	writer := lzw.NewWriter(&buf, lzw.LSB, 8)
	_, err := writer.Write(input)
	if err != nil {
		return nil, err
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// Decompress распаковывает данные, сжатые с помощью алгоритма LZW.
func Decompress(input []byte) ([]byte, error) {
	buf := bytes.NewReader(input)
	reader := lzw.NewReader(buf, lzw.LSB, 8)
	defer reader.Close()

	var result bytes.Buffer
	_, err := io.Copy(&result, reader)
	if err != nil {
		return nil, err
	}
	return result.Bytes(), nil
}

// SaveToFile сохраняет данные в файл.
func SaveToFile(filename string, data []byte) error {
	return os.WriteFile(filename, data, 0644)
}

// ReadFromFile считывает данные из файла.
func ReadFromFile(filename string) ([]byte, error) {
	return os.ReadFile(filename)
}

// DeleteFile удаляет файл.
func DeleteFile(filename string) error {
	return os.Remove(filename)
}

func main() {
	// Тестовые данные для сжатия и распаковки
	data := []byte("dd9900-M41 ASCII символы 123 123 444!")
	fmt.Println("Оригинальные данные:", string(data))
	fmt.Printf("Размер оригинальных данных: %d байт\n", len(data))

	// Сжатие данных
	compressedData, err := Compress(data)
	if err != nil {
		fmt.Println("Ошибка сжатия:", err)
		return
	}
	fmt.Printf("Сжатые данные: %v\n", compressedData)
	fmt.Printf("Размер сжатых данных: %d байт\n", len(compressedData))

	// Сохранение сжатых данных в файл
	filename := "compressed.lzw"
	err = SaveToFile(filename, compressedData)
	if err != nil {
		fmt.Println("Ошибка сохранения файла:", err)
		return
	}
	//fmt.Printf("Сжатые данные сохранены в файл '%s'\n", filename)

	// Загрузка и распаковка данных из файла
	compressedDataFromFile, err := ReadFromFile(filename)
	if err != nil {
		fmt.Println("Ошибка чтения файла:", err)
		return
	}

	// Распаковка данных
	decompressedData, err := Decompress(compressedDataFromFile)
	if err != nil {
		fmt.Println("Ошибка распаковки:", err)
		return
	}
	fmt.Println("Распакованные данные:", string(decompressedData))
	fmt.Printf("Размер распакованных данных: %d байт\n", len(decompressedData))

	// Удаление файла после завершения работы
	err = DeleteFile(filename)
	if err != nil {
		fmt.Println("Ошибка удаления файла:", err)
		return
	}
	//fmt.Printf("Файл '%s' был успешно удален\n", filename)
}
