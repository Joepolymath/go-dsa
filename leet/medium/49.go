package leet

import (
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
    hashMap := make(map[string][]string)

    for _, word := range strs {
        wordArray := strings.Split(word, "")
			sort.Strings(wordArray)
        sortedString := strings.Join(wordArray, "")
        if _, exists := hashMap[sortedString]; exists {
            hashMap[sortedString] = append(hashMap[sortedString], word)
        } else {
            hashMap[sortedString] = []string{word}
        }
    }
    result := make([][]string, 0)
    for _, value := range hashMap {
        result = append(result, value)
    }
    return result
}