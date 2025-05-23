package main

import (
 "fmt"
 "io"
 "os"
 "strconv"
 "strings"
)

func sortNumbers(slice []int64) {
 length := len(slice)
 for i := 0; i < length; i++ {
  changed := false
  for j := 0; j < length-1; j++ {
   if slice[j] > slice[j+1] {
    slice[j], slice[j+1] = slice[j+1], slice[j]
    changed = true
   }
  }
  if !changed {
   break
  }
 }
}

func main() {
 inputFile, err := os.Open("input.txt")
 if err != nil {
  fmt.Println("No input")
  os.Exit(1)
 }

 data := make([]byte, 64)
 for {
  _, err := inputFile.Read(data)
  if err == io.EOF {
   break
  }
 }

 var numbers []int64
 for _, part := range strings.Split(string(data), " ") {
  for _, subpart := range strings.Split(string(part), "\r") {
   for _, item := range strings.Split(string(subpart), "\x00") {
    value, _ := strconv.ParseInt(item, 10, 64)
    if value != 0 {
     numbers = append(numbers, value)
    }
   }
  }
 }
 inputFile.Close()
 
 outputFile, err := os.Create("output.txt")
 defer outputFile.Close()

 sortNumbers(numbers)

 result := ""
 for _, num := range numbers {
  result += fmt.Sprintf("%d ", num)
 }

 outputFile.WriteString(result)

 fmt.Println(numbers)
}
