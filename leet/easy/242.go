package leet

import (
	"strings"
)

// Question link https://leetcode.com/problems/valid-anagram/
func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    sMap := make(map[string]int)
    tMap := make(map[string]int)

    sArray := strings.Split(s, "")
    tArray := strings.Split(t, "")

    for _, char := range sArray {
        if _, found := sMap[char]; found {
            sMap[char]++
        } else {
            sMap[char] = 1
        }
    }

    for _, char := range tArray {
        if _, found := tMap[char]; found {
            tMap[char]++
        } else {
            tMap[char] = 1
        }
    }

    // comparing the two maps
    for idx := range sMap {
        if sMap[idx] != tMap[idx] {
            return false
        }
    }
    return true
}