# English Cricket Word Game (CLI)

A nostalgic word-guessing game inspired by cricket, implemented in Go.

## Quick Start

```
go run cmd/ec/main.go
```

## About the Game

English Cricket is a word game that combines vocabulary skills with the structure of cricket. Players "bowl" letters and "bat" by guessing words starting with those letters. The length of each word determines the score.

### Key Features

- Single-player (vs. System) or Two-player modes
- Customizable number of overs (1-3)
- Coin toss to decide who bats first
- Score tracking based on word length

## How to Play

1. Choose game mode (Single Team or Two Team)
2. Participate in coin toss
3. Select to bat or bowl
4. Choose number of overs
5. Start guessing words!

### Example Gameplay

```
$ ec --start
Welcome to English cricket
Choose your team:
1. Single Team (You vs. System)
2. Two Team (You vs. Human)
Input: 1

You selected Single Team
...

System: D
You: Determination
Score: 13

System: E
You: Elephant
Score: 8

...
```

## Project Structure

```
english-cricket/
├── cmd/ec/
│   └── main.go
├── internal/
│   ├── game/
│   ├── ui/
│   └── utils/
├── words.txt
├── go.mod
└── go.sum
```

## Contributing

Contributions are welcome! Feel free to submit issues or pull requests.

**#englishcricket #golang #wordgame #cli**