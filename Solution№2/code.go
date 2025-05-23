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

func convertToNumbers(stringValues []string) ([]int64, error) {
 var numbers []int64
 for _, str := range stringValues {
  cleanedStr := strings.TrimSpace(str)
  if cleanedStr == "" {
   continue
  }
  num, err := strconv.ParseInt(cleanedStr, 10, 64)
  if err != nil {
   log.Printf("Ошибка преобразования '%s': %v", str, err)
   continue
  }
  numbers = append(numbers, num)
 }
 return numbers, nil
}

func main() {
 sourceFile, err := os.Open("input.txt")
 if err != nil {
  log.Fatalf("Ошибка открытия файла: %v", err)
 }
 defer sourceFile.Close()

 fileScanner := bufio.NewScanner(sourceFile)
 fileScanner.Split(bufio.ScanWords)
 var textParts []string
 for fileScanner.Scan() {
  textParts = append(textParts, fileScanner.Text())
 }

 numericValues, err := convertToNumbers(textParts)
 if err != nil {
  log.Fatalf("Ошибка преобразования данных: %v", err)
 }


 sort.Slice(numericValues, func(i, j int) bool {
  return numericValues[i] < numericValues[j]
 })

 resultFile, err := os.Create("output.txt")
 if err != nil {
  log.Fatalf("Ошибка создания файла: %v", err)
 }
 defer resultFile.Close()

 fileWriter := bufio.NewWriter(resultFile)
 for _, value := range numericValues {
  fmt.Fprintf(fileWriter, "%d ", value)
 }
 fileWriter.Flush()

 fmt.Println("Результат сохранен в output.txt.")
}
