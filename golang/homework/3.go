// Частотный анализ

// Написать функцию, которая получает на вход текст и возвращает

// 10 самых часто встречающихся слов без учета словоформ

package main

import (
	"fmt"
	"sort"
	"strings"
)

func word_frequency_analyze(text string, amount int) []string {
	words := strings.Split(text, " ")
	words_map := make(map[string]int)
	result := []string{}

	for _, word := range words {
		if _, exists := words_map[word]; !exists {
			result = append(result, word)
		}
		words_map[word] = words_map[word] + 1

	}
	fmt.Println(words_map)
	result = sort_words_by_freq(result, words_map)

	
	return result[:amount]
}

func sort_words_by_freq(words []string, freq map[string]int) []string {
	sort.Silce(words, func(i, j int) bool {
		return freq[words[i]] > freq[words[j]]
	})
	return words
}

func main() {
	text := "The quick brown fox jumps over the lazy dog. The cat jumps over the brown dog. The fox jumps over the lazy cat. The dog jumps over the lazy fox. The lazy cat jumps over the brown dog. The brown dog jumps over the lazy cat."
	result := word_frequency_analyze(text, 10)
	fmt.Println(result)

}
