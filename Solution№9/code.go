package main

import (
 "fmt"
 "os"
)

func applyAtbashCipher(originalText string) string {
 encryptedText := ""
 for _, character := range originalText {
  if character >= 'A' && character <= 'Z' {
   // Преобразование для заглавных букв (A=65, Z=90)
   encryptedText += string(90 - (character - 65))
  } else if character >= 'a' && character <= 'z' {
   // Преобразование для строчных букв (a=97, z=122)
   encryptedText += string(122 - (character - 97))
  } else {
   encryptedText += string(character)
  }
 }
 return encryptedText
}

func applyCaesarCipher(originalText string, shiftAmount int) string {
 encryptedText := ""
 for _, character := range originalText {
  if character >= 'A' && character <= 'Z' {
   newPosition := (int(character)-65+shiftAmount)%26 + 65
   encryptedText += string(newPosition)
  } else if character >= 'a' && character <= 'z' {
   newPosition := (int(character)-97+shiftAmount)%26 + 97
   encryptedText += string(newPosition)
  } else {
   encryptedText += string(character)
  }
 }
 return encryptedText
}

func main() {
 fileContent, err := os.ReadFile("input.txt")
 if err != nil {
  fmt.Println("Ошибка при чтении файла:", err)
  return
 }
 textToEncrypt := string(fileContent)

 var shiftValue int
 fmt.Print("Укажите величину сдвига для шифра Цезаря: ")
 _, err = fmt.Scan(&shiftValue)
 if err != nil {
  fmt.Println("Ошибка при вводе сдвига:", err)
  return
 }

 atbashEncrypted := applyAtbashCipher(textToEncrypt)
 caesarEncrypted := applyCaesarCipher(textToEncrypt, shiftValue)

 result := fmt.Sprintf(
  "Атбаш: %s\n Цезаря (сдвиг %d): %s",
  atbashEncrypted,
  shiftValue,
  caesarEncrypted,
 )

 err = os.WriteFile("output.txt", []byte(result), 0644)
 if err != nil {
  fmt.Println("Ошибка при записи в файл:", err)
  return
 }

 fmt.Println("Операция завершена! Результаты сохранены в output.txt")
}
