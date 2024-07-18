package utils

import (
    "bufio"
    "os"
    "strings"
)

var wordList map[string]bool

func init() {
    wordList = make(map[string]bool)
    file, err := os.Open("words.txt")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        wordList[strings.ToLower(scanner.Text())] = true
    }
}

func IsValidWord(word string) bool {
    return wordList[strings.ToLower(word)]
}