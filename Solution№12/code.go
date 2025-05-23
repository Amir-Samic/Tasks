package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

//Решение квадратного уравнения
func Binary(a, b, c float64) (string, string) {
	var D = (b * b) - (4 * a * c)

	if D >= 0 {
		x1 := ((-b) + math.Sqrt(D)) / (2 * a)
		x2 := ((-b) - math.Sqrt(D)) / (2 * a)
		return strconv.FormatFloat(x1, 'f', 2, 64), strconv.FormatFloat(x2, 'f', 2, 64)
	} else {
		sqrt := strconv.FormatFloat(math.Sqrt(-D)/(2*a), 'f', 2, 64)

		x1 := strconv.FormatFloat(-b/(2*a), 'f', 1, 64) + "+" + sqrt + "i"
		x2 := strconv.FormatFloat(-b/(2*a), 'f', 2, 64) + "-" + sqrt + "i"

		return x1, x2
	}

}

//главная функция
func main() {
	filepath := "input.txt"
	content, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println("Error")
		return
	}
	nums := strings.Fields(string(content))

	fmt.Println(nums)

	var a, b, c float64
	a, _ = strconv.ParseFloat(nums[0], 64)
	b, _ = strconv.ParseFloat(nums[1], 64)
	c, _ = strconv.ParseFloat(nums[2], 64)

	x1, x2 := Binary(a, b, c)

	fmt.Println(x1, x2)


	f, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error")
		return
	}
	defer f.Close()

f.WriteString(x1)
f.WriteString("\n")
f.WriteString(x2)

}
