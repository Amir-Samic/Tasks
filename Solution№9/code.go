package main

import (
	"fmt"
	"os"
)

//Шифр Атбаш
func atbash(text string) string {
	result := ""
	for _, char := range text {
		if char >= 'A' && char <= 'Z' {
			result += string(90 - (char - 65))
		} else if char >= 'a' && char <= 'z' {
			result += string(122 - (char - 97))
		} else {
			result += string(char)
		}
	}
	return result
}

//Шифр Цезаря
func caesar(text string, shift int) string {
	result := ""
	for _, char := range text {
		if char >= 'A' && char <= 'Z' {
			result += string((int(char)-65+shift)%26 + 65)
		} else if char >= 'a' && char <= 'z' {
			result += string((int(char)-97+shift)%26 + 97)
		} else {
			result += string(char)
		}
	}
	return result
}

//Главная функция
func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Ошибка чтения файла:", err)
		return
	}
	inputText := string(data)

	var shift int
	fmt.Print("Введите сдвиг для шифра Цезаря: ")
	_, err = fmt.Scan(&shift)
	if err != nil {
		fmt.Println("Ошибка ввода сдвига:", err)
		return
	}

	atbashText := atbash(inputText)
	caesarText := caesar(shift, inputText)

	outputText := fmt.Sprintf(
		"Атбаш: %s\nЦезарь: (сдвиг %d): %s",
		atbashText,
		shift,
		caesarText,
	)

	err = os.WriteFile("output.txt", []byte(outputText), 0644)
	if err != nil {
		fmt.Println("Ошибка записи в файл:", err)
		return
	}

	fmt.Println("Готово! Результат записан в output.txt")
}
