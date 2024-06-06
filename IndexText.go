package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

// Зчитуємо текст з файлу
func readFile(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return lines
}

// Функція для нормалізації слова: видаляє пунктуацію і переводить в нижній регістр
func normalizeWord(word string) string {
	var b strings.Builder
	for _, r := range word {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			b.WriteRune(unicode.ToLower(r))
		}
	}
	return b.String()
}

// Індексуємо текст по словам (case insensitive)
func indexText(lines []string) map[string][]int {
	index := make(map[string][]int)
	for i, line := range lines {
		words := strings.Fields(line)
		for _, word := range words {
			normalizedWord := normalizeWord(word)
			index[normalizedWord] = append(index[normalizedWord], i)
		}
	}
	return index
}

// Знаходимо всі рядки за словом (case insensitive)
func searchByWord(lines []string, index map[string][]int, query string) []string {
	normalizedQuery := normalizeWord(query)
	var results []string
	if indices, found := index[normalizedQuery]; found {
		for _, idx := range indices {
			results = append(results, lines[idx])
		}
	}
	return results
}

// Пошук тексту
func searchText(lines []string, index map[string][]int) {
	fmt.Print("Введіть слово для пошуку: ")
	reader := bufio.NewReader(os.Stdin)
	query, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	query = strings.TrimSpace(query) // Видаляємо символ нового рядка і пробіли

	results := searchByWord(lines, index, query)
	if len(results) == 0 {
		fmt.Println("Рядок не знайдено.")
		return
	}
	fmt.Println("Знайдені рядки:")
	for _, line := range results {
		fmt.Println(line)
	}
}

func main() {
	filename := "text.txt"
	lines := readFile(filename)

	// Індексуємо текст
	index := indexText(lines)
	fmt.Println("Проіндексовані слова (для тестування):")
	for word, indices := range index {
		fmt.Printf("%s: %v\n", word, indices)
	}

	// Пошук тексту
	searchText(lines, index)
}
