package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func parseInts(strs []string) ([]int64, error) {
	var result []int64
	for _, s := range strs {
		arr := strings.TrimSpace(s)
		if arr == "" {
			continue
		}
		i, err := strconv.ParseInt(arr, 10, 64)
		if err != nil {
			log.Printf("Не удалось преобразовать '%s': %v", s, err)
			continue
		}
		result = append(result, i)
	}
	return result, nil
}


func main() {
	inputFile, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Не удалось открыть файл: %v", err)
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	scanner.Split(bufio.ScanWords)
	var words []string
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	ints, err := parseInts(words)
	if err != nil {
		log.Fatalf("Ошибка при разборе целых чисел: %v", err)
	}
	sort.Slice(ints, func(i, j int) bool {
		return ints[i] < ints[j]
	})

	outputFile, err := os.Create("output.txt")
	if err != nil {
		log.Fatalf("Не удается создать файл: %v", err)
	}
	defer outputFile.Close()

	writer := bufio.NewWriter(outputFile)
	for _, val := range ints {
		fmt.Fprintf(writer, "%d ", val)
	}
	writer.Flush()

	fmt.Println("Данные записываются в output.txt.")
}
