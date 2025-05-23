package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readMatrix(filename string) [][]float64 {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var matrix [][]float64
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		fields := strings.Fields(line)
		row := make([]float64, len(fields))

		for i, numStr := range fields {
			num, err := strconv.ParseFloat(numStr, 64)
			if err != nil {
				panic(err)
			}
			row[i] = num
		}

		matrix = append(matrix, row)
	}

	return matrix
}

func writeResults(filename string, det, tr float64, ts [][]float64) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	fmt.Fprintf(writer, "Определитель: %.2f\n", det)
	fmt.Fprintf(writer, "След: %.2f\n", tr)
	fmt.Fprintln(writer, "Транспонированная матрица:")

	for _, row := range ts {
		for _, val := range row {
			fmt.Fprintf(writer, "%.2f ", val)
		}
		fmt.Fprintln(writer)
	}
}

// Определитель
func determinant(matrix [][]float64) float64 {
	n := len(matrix)
	if n == 1 {
		return matrix[0][0]
	}
	if n == 2 {
		return matrix[0][0]*matrix[1][1] - matrix[0][1]*matrix[1][0]
	}

	det := 0.0
	for col := 0; col < n; col++ {
		minor := make([][]float64, n-1)
		for i := range minor {
			minor[i] = make([]float64, n-1)
		}

		for i := 1; i < n; i++ {
			colIdx := 0
			for j := 0; j < n; j++ {
				if j == col {
					continue
				}
				minor[i-1][colIdx] = matrix[i][j]
				colIdx++
			}
		}

		sign := 1.0
		if col%2 == 1 {
			sign = -1.0
		}
		det += sign * matrix[0][col] * determinant(minor)
	}
	return det
}

// След
func trace(matrix [][]float64) float64 {
	tr := 0.0
	for i := 0; i < len(matrix); i++ {
		tr += matrix[i][i]
	}
	return tr
}

// Транспонирование
func transpose(matrix [][]float64) [][]float64 {
	n := len(matrix)
	transposed := make([][]float64, n)
	for i := range transposed {
		transposed[i] = make([]float64, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			transposed[j][i] = matrix[i][j]
		}
	}
	return transposed
}

func main() {
	matrix := readMatrix("input.txt")

	n := len(matrix)
	for _, row := range matrix {
		if len(row) != n {
			fmt.Println("Ошибка: матрица должна быть квадратной")
			return
		}
	}

	det := determinant(matrix)
	tr := trace(matrix)
	ts := transpose(matrix)

	writeResults("output.txt", det, tr, ts)

	fmt.Printf("Результаты сохранены в output.txt\n")
	fmt.Printf("Определитель: %.2f\n", det)
	fmt.Printf("След: %.2f\n", tr)
	fmt.Println("Транспонированная матрица:")
	for _, row := range ts {
		fmt.Println(row)
	}
}
