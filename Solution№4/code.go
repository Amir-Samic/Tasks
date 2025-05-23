package main

import (
 "bufio"
 "fmt"
 "os"
 "strconv"
 "strings"
)

func loadMatrix(fileName string) [][]float64 {
 file, err := os.Open(fileName)
 if err != nil {
  panic(err)
 }
 defer file.Close()

 var data [][]float64
 scanner := bufio.NewScanner(file)

 for scanner.Scan() {
  textLine := scanner.Text()
  if textLine == "" {
   continue
  }

  values := strings.Fields(textLine)
  rowData := make([]float64, len(values))

  for idx, strValue := range values {
   numValue, err := strconv.ParseFloat(strValue, 64)
   if err != nil {
    panic(err)
   }
   rowData[idx] = numValue
  }

  data = append(data, rowData)
 }

 return data
}

func saveResults(fileName string, detValue, traceValue float64, transposedData [][]float64) {
 file, err := os.Create(fileName)
 if err != nil {
  panic(err)
 }
 defer file.Close()

 bufferedWriter := bufio.NewWriter(file)
 defer bufferedWriter.Flush()

 fmt.Fprintf(bufferedWriter, "Определитель: %.2f\n", detValue)
 fmt.Fprintf(bufferedWriter, "След: %.2f\n", traceValue)
 fmt.Fprintln(bufferedWriter, "Транспонированная матрица:")

 for _, row := range transposedData {
  for _, value := range row {
   fmt.Fprintf(bufferedWriter, "%.2f ", value)
  }
  fmt.Fprintln(bufferedWriter)
 }
}

func calculateDeterminant(mat [][]float64) float64 {
 size := len(mat)
 if size == 1 {
  return mat[0][0]
 }
 if size == 2 {
  return mat[0][0]*mat[1][1] - mat[0][1]*mat[1][0]
 }

 result := 0.0
 for col := 0; col < size; col++ {
  subMatrix := make([][]float64, size-1)
  for i := range subMatrix {
   subMatrix[i] = make([]float64, size-1)
  }

  for i := 1; i < size; i++ {
   colIndex := 0
   for j := 0; j < size; j++ {
    if j == col {
     continue
    }
    subMatrix[i-1][colIndex] = mat[i][j]
    colIndex++
   }
  }

  sign := 1.0
  if col%2 == 1 {
   sign = -1.0
  }
  result += sign * mat[0][col] * calculateDeterminant(subMatrix)
 }
 return result
}

func calculateTrace(mat [][]float64) float64 {
 sum := 0.0
 for i := 0; i < len(mat); i++ {
  sum += mat[i][i]
 }
 return sum
}

func transposeMatrix(mat [][]float64) [][]float64 {
 size := len(mat)
 result := make([][]float64, size)
 for i := range result {
  result[i] = make([]float64, size)
 }

 for i := 0; i < size; i++ {
  for j := 0; j < size; j++ {
   result[j][i] = mat[i][j]
  }
 }
 return result
}

func main() {
 matrixData := loadMatrix("input.txt")

 matrixSize := len(matrixData)
 for _, row := range matrixData {
  if len(row) != matrixSize {
   fmt.Println("Ошибка: матрица должна быть квадратной")
   return
  }
 }

 det := calculateDeterminant(matrixData)
 tr := calculateTrace(matrixData)
 transposed := transposeMatrix(matrixData)

 saveResults("output.txt", det, tr, transposed)

 fmt.Printf("Результаты сохранены в output.txt\n")
 fmt.Printf("Определитель: %.2f\n", det)
 fmt.Printf("След: %.2f\n", tr)
 fmt.Println("Транспонированная матрица:")
 for _, row := range transposed {
  fmt.Println(row)
 }
}
