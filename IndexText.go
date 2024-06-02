package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Індексуємо текст по словам
func indexText(text []string) map[string]map[int]struct{} {
	index := make(map[string]map[int]struct{})
	for i, line := range text {
		words := strings.Fields(line)
		for _, word := range words {
			word = strings.ToLower(word)
			if _, exists := index[word]; !exists {
				index[word] = make(map[int]struct{})
			}
			index[word][i] = struct{}{}
		}
	}
	return index
}

// Знаходимо всі рядки за словом
func searchByWord(index map[string]map[int]struct{}, text []string, word string) []string {
	word = strings.ToLower(word)
	lines := []string{}
	if lineIndices, found := index[word]; found {
		for lineIndex := range lineIndices {
			lines = append(lines, text[lineIndex])
		}
	}
	return lines
}

// Зчитуємо текст з файлу
func readTextFromFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var text []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return text, nil
}

func main() {
	// Зчитуємо текст з файлу
	filename := "text.txt"
	text, err := readTextFromFile(filename)
	if err != nil {
		fmt.Println("Помилка читання файлу:", err)
		return
	}

	// Індексуємо текст
	index := indexText(text)

	// Отримання пошукового запиту від користувача
	var query string
	fmt.Print("Введіть слово для пошуку: ")
	fmt.Scanln(&query)

	// Пошук рядків, які містять пошукове слово
	results := searchByWord(index, text, query)

	// Виведення результатів пошуку
	fmt.Println("Результати пошуку:")
	for _, result := range results {
		fmt.Println(result)
	}
}
