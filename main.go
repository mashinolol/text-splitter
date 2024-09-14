package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func splitText(text string, partSize int) []string {
	var parts []string
	textRunes := []rune(text)
	textLen := len(textRunes)
	partTextSize := partSize - 5

	if textLen <= partTextSize {
		part := fmt.Sprintf("[1\\1] %s", text)
		parts = append(parts, part)
		return parts
	}

	numParts := (textLen + partTextSize - 1) / partTextSize

	for i := 0; i < numParts; i++ {
		start := i * partTextSize
		end := start + partTextSize
		if end > textLen {
			end = textLen
		}
		partText := string(textRunes[start:end])
		part := fmt.Sprintf("[%d\\%d] %s", i+1, numParts, partText)
		parts = append(parts, part)
	}

	return parts
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите текст: ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)

	var partSize int
	fmt.Print("Введите размер части: ")
	fmt.Scanf("%d", &partSize)

	if partSize <= 5 {
		fmt.Println("Размер части должен быть больше 5, чтобы учитывать метку [x\\y]")
		return
	}

	if utf8.RuneCountInString(text) == 0 {
		fmt.Println("Текст не может быть пустым")
		return
	}

	parts := splitText(text, partSize)

	for _, part := range parts {
		fmt.Println(part)
	}
}
