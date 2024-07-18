package ui

import (
    "fmt"
    "strings"
	
	"github.com/charmbracelet/bubbles/spinner"
    "github.com/charmbracelet/bubbles/textinput"
    tea "github.com/charmbracelet/bubbletea"
    "github.com/charmbracelet/lipgloss"
    "github.com/yasindce1998/english-cricket/internal/game"
    "github.com/yasindce1998/english-cricket/internal/utils"
	
)

var (
    focusedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
    blurredStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
    titleStyle   = lipgloss.NewStyle().MarginLeft(2).Foreground(lipgloss.Color("170")).Bold(true)
    errorStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("196")).Bold(true)
	scoreStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("39")).Bold(true)
    letterStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("226")).Bold(true)
)


type model struct {
    game           *game.Game
    textInput      textinput.Model
    err            error
    letter         rune
    quitting       bool
    inningsOver    bool
    gameStats      string
	lastWordScore int
	spinner        spinner.Model
    loading        bool
}

func InitialModel(g *game.Game) model {
    ti := textinput.New()
    ti.Placeholder = "Enter a word"
    ti.Focus()
    ti.CharLimit = 156
    ti.Width = 20

    s := spinner.New()
    s.Spinner = spinner.Dot
    s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

    return model{
        game:      g,
        textInput: ti,
        err:       nil,
        letter:    g.NextBall(),
        spinner:   s,
        loading:   false,
    }
}

func (m model) Init() tea.Cmd {
    return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    var cmd tea.Cmd

    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.Type {
        case tea.KeyEnter:
            word := m.textInput.Value()
            if !utils.IsValidWord(word) || !strings.HasPrefix(strings.ToLower(word), strings.ToLower(string(m.letter))) {
                m.err = fmt.Errorf("Invalid word")
            } else {
                score, err := m.game.PlayBall(word)
                if err != nil {
                    m.err = err
                } else {
                    m.err = nil
                    m.letter = m.game.NextBall()
					m.lastWordScore = score
                    if m.game.IsOver {
                        m.quitting = true
						m.UpdateGameStats()
                        return m, tea.Quit
                    }
                }
            }
            m.textInput.SetValue("")
        case tea.KeyCtrlC, tea.KeyEsc:
            m.quitting = true
            return m, tea.Quit
        }
	case spinner.TickMsg:
        if m.loading {
            var spinnerCmd tea.Cmd
            m.spinner, spinnerCmd = m.spinner.Update(msg)
            return m, spinnerCmd
        }

    case error:
        m.err = msg
        return m, nil
    }

    m.textInput, cmd = m.textInput.Update(msg)
    return m, cmd
}

func (m model) View() string {
    if m.quitting {
        return m.gameStats
    }
    
    var s strings.Builder
    s.WriteString(titleStyle.Render("English Cricket Word Game"))
    s.WriteString(fmt.Sprintf("\nBall %d/%d\n", m.game.CurrentBall-1, m.game.TotalBalls))
    s.WriteString(fmt.Sprintf("Current letter: %s\n", letterStyle.Render(string(m.letter))))
    s.WriteString(fmt.Sprintf("Current Batting: %s\n", m.game.CurrentBat.Name))
    s.WriteString(fmt.Sprintf("Score: %s\n\n", scoreStyle.Render(fmt.Sprintf("%d", m.game.CurrentBat.Score))))
    
    if m.loading {
        s.WriteString(m.spinner.View() + " Thinking...\n\n")
    } else {
        s.WriteString(focusedStyle.Render(m.textInput.View()))
        s.WriteString("\n\n")
    }
    
    if m.err != nil {
        s.WriteString(errorStyle.Render(fmt.Sprintf("Error: %v\n", m.err)))
    }
    
    if m.inningsOver {
        s.WriteString(titleStyle.Render("\nInnings Over!\n"))
        s.WriteString(fmt.Sprintf("Press Enter to start next innings or Ctrl+C to quit.\n"))
    } else {
        s.WriteString("\nPress Ctrl+C to quit.\n")
    }
	if m.lastWordScore > 0 {
        s.WriteString(fmt.Sprintf("Last word score: %d\n", m.lastWordScore))
    }
    
    return s.String()
}


func (m *model) UpdateGameStats() {
    var s strings.Builder
    s.WriteString(titleStyle.Render("Game Over! Final Statistics:\n\n"))
    s.WriteString(fmt.Sprintf("Player 1 (%s) Score: %d\n", m.game.Player1.Name, m.game.Player1.Score))
    s.WriteString(fmt.Sprintf("Player 2 (%s) Score: %d\n", m.game.Player2.Name, m.game.Player2.Score))
    s.WriteString(fmt.Sprintf("Highest Scoring Word: %s\n", m.game.HighestWord))
    s.WriteString(fmt.Sprintf("Average Word Length: %.2f\n", m.game.GetAverageWordLength()))
    
    winner := "Tie"
    if m.game.Player1.Score > m.game.Player2.Score {
        winner = m.game.Player1.Name
    } else if m.game.Player2.Score > m.game.Player1.Score {
        winner = m.game.Player2.Name
    }
    s.WriteString(fmt.Sprintf("Winner: %s\n", winner))
    
    m.gameStats = s.String()
}

func RunGame(g *game.Game) {
    m := InitialModel(g)
    p := tea.NewProgram(m)
    if err := p.Start(); err != nil {
        fmt.Printf("Error: %v", err)
    }
}