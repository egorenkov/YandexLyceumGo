package main

import (
	"fmt"
	"strings"
)

func CheckLetters(text string) string {
	if len(text) <= 1 {
		if text == "e" {
			return fmt.Sprintf("Количество возможных ошибок: %d, перепроверьте текст", 1)
		}
		return "Текст готов к публикации!"
	}

	if strings.ContainsRune(text, 'е') {
		return fmt.Sprintf("Количество возможных ошибок: %d, перепроверьте текст", strings.Count(text, string("е")))
	}
	return "Текст готов к публикации!"

}

func main() {
	fmt.Println(CheckLetters(""))
	fmt.Println(CheckLetters("Тут точно нет сложных букв, хотя ж выглядит подозрительно"))
	fmt.Println(CheckLetters("Ох уж эти е и ё"))
	fmt.Println(CheckLetters("Е ее ее тум турун тум турун тум"))
}
