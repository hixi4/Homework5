package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

// TextIndex представляє індекс тексту
type TextIndex struct {
	index map[string][]int
}

// NewTextIndex створює новий TextIndex
func NewTextIndex() *TextIndex {
	return &TextIndex{
		index: make(map[string][]int),
	}
}

// Індексує текст по словам (case insensitive)
func (ti *TextIndex) IndexText(lines []string) {
	for i, line := range lines {
		words := strings.Fields(line)
		for _, word := range words {
			normalizedWord := normalizeWord(word)
			ti.index[normalizedWord] = append(ti.index[normalizedWord], i)
		}
	}
}

// Знаходить всі рядки за словом (case insensitive)
func (ti *TextIndex) SearchByWord(lines []string, query string) []string {
	normalizedQuery := normalizeWord(query)
	var results []string
	if indices, found := ti.index[normalizedQuery]; found {
		for _, idx := range indices {
			results = append(results, lines[idx])
		}
	}
	return results
}

// Зчитує текст з файлу
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

// Пошук тексту
func searchText(lines []string, ti *TextIndex) {
	fmt.Print("Введіть слово для пошуку: ")
	reader := bufio.NewReader(os.Stdin)
	query, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	query = strings.TrimSpace(query) // Видаляємо символ нового рядка і пробіли

	results := ti.SearchByWord(lines, query)
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

	// Створюємо та індексуємо текст
	ti := NewTextIndex()
	ti.IndexText(lines)
	fmt.Println("Проіндексовані слова (для тестування):")
	for word, indices := range ti.index {
		fmt.Printf("%s: %v\n", word, indices)
	}

	// Пошук тексту
	searchText(lines, ti)
}
