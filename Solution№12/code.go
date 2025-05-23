package main

import (
 "fmt"
 "math"
 "os"
 "strconv"
 "strings"
)

// Функция решения квадратного уравнения
func solveQuadraticEquation(coeffA, coeffB, coeffC float64) (root1, root2 string) {
 discriminant := (coeffB * coeffB) - (4 * coeffA * coeffC)

 if discriminant >= 0 {
  // Действительные корни
  realRoot1 := ((-coeffB) + math.Sqrt(discriminant)) / (2 * coeffA)
  realRoot2 := ((-coeffB) - math.Sqrt(discriminant)) / (2 * coeffA)
  return strconv.FormatFloat(realRoot1, 'f', 2, 64), 
         strconv.FormatFloat(realRoot2, 'f', 2, 64)
 } else {
  // Комплексные корни
  imaginaryPart := strconv.FormatFloat(math.Sqrt(-discriminant)/(2*coeffA), 'f', 2, 64)
  realPart := strconv.FormatFloat(-coeffB/(2*coeffA), 'f', 2, 64)

  complexRoot1 := realPart + "+" + imaginaryPart + "i"
  complexRoot2 := realPart + "-" + imaginaryPart + "i"

  return complexRoot1, complexRoot2
 }
}

func main() {
 inputFilePath := "input.txt"
 fileContent, err := os.ReadFile(inputFilePath)
 if err != nil {
  fmt.Println("Ошибка чтения файла")
  return
 }

 coefficients := strings.Fields(string(fileContent))
 fmt.Println(coefficients)

 var a, b, c float64
 a, _ = strconv.ParseFloat(coefficients[0], 64)
 b, _ = strconv.ParseFloat(coefficients[1], 64)
 c, _ = strconv.ParseFloat(coefficients[2], 64)

 solution1, solution2 := solveQuadraticEquation(a, b, c)
 fmt.Println(solution1, solution2)

 outputFile, err := os.Create("output.txt")
 if err != nil {
  fmt.Println("Ошибка создания файла")
  return
 }
 defer outputFile.Close()

 outputFile.WriteString(solution1 + "\n")
 outputFile.WriteString(solution2)
}
