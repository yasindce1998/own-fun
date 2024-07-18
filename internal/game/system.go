package game

import (
    "math/rand"
    "bufio"
    "strings"
    "time"
    "os"
    "unicode"
)

type System struct {
    Player *Player
    words  map[rune][]string
}

func NewSystem() *System {
    return &System{
        Player: NewPlayer("System"),
        words:  loadWords(),
    }
}

func (s *System) Bowl() rune {
    rand.Seed(time.Now().UnixNano())
    return rune('A' + rand.Intn(26))
}

func (s *System) ChooseWord(letter rune, difficulty DifficultyLevel) string {
    words := s.words[letter]
    if len(words) == 0 {
        return ""
    }
    
    var filteredWords []string
    switch difficulty {
    case Easy:
        filteredWords = filterWordsByLength(words, 3, 5)
    case Medium:
        filteredWords = filterWordsByLength(words, 5, 8)
    case Hard:
        filteredWords = filterWordsByLength(words, 8, 100)
    }
    
    if len(filteredWords) == 0 {
        return words[rand.Intn(len(words))]
    }
    return filteredWords[rand.Intn(len(filteredWords))]
}

func filterWordsByLength(words []string, minLength, maxLength int) []string {
    var filtered []string
    for _, word := range words {
        if len(word) >= minLength && len(word) <= maxLength {
            filtered = append(filtered, word)
        }
    }
    return filtered
}

func loadWords() map[rune][]string {
    wordMap := make(map[rune][]string)

    file, err := os.Open("words.txt")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        word := strings.TrimSpace(scanner.Text())
        if word == "" {
            continue
        }

        firstLetter := unicode.ToLower(rune(word[0]))
        wordMap[firstLetter] = append(wordMap[firstLetter], word)
    }

    if err := scanner.Err(); err != nil {
        panic(err)
    }

    return wordMap
}