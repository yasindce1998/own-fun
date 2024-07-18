package ui

import (
    "fmt"
    "strings"

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
)

type model struct {
    game      *game.Game
    textInput textinput.Model
    err       error
    letter    rune
    quitting  bool
}

func InitialModel(g *game.Game) model {
    ti := textinput.New()
    ti.Placeholder = "Enter a word"
    ti.Focus()
    ti.CharLimit = 156
    ti.Width = 20

    return model{
        game:      g,
        textInput: ti,
        err:       nil,
        letter:    g.NextBall(),
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
                    if m.game.IsOver {
                        m.quitting = true
                        return m, tea.Quit
                    }
                }
            }
            m.textInput.SetValue("")
        case tea.KeyCtrlC, tea.KeyEsc:
            m.quitting = true
            return m, tea.Quit
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
        return fmt.Sprintf("Game over! Final score: %d\n", m.game.CurrentBat.Score)
    }
    
    var s strings.Builder
    s.WriteString(titleStyle.Render("English Cricket Word Game"))
    s.WriteString(fmt.Sprintf("\nBall %d/%d\n", m.game.CurrentBall-1, m.game.TotalBalls))
    s.WriteString(fmt.Sprintf("Current letter: %c\n", m.letter))
    s.WriteString(fmt.Sprintf("Score: %d\n\n", m.game.CurrentBat.Score))
    
    s.WriteString(focusedStyle.Render(m.textInput.View()))
    s.WriteString("\n\n")
    
    if m.err != nil {
        s.WriteString(errorStyle.Render(fmt.Sprintf("Error: %v\n", m.err)))
    }
    
    s.WriteString("\nPress Ctrl+C to quit.\n")
    
    return s.String()
}

func RunGame(g *game.Game) {
    m := InitialModel(g)
    p := tea.NewProgram(m)
    if err := p.Start(); err != nil {
        fmt.Printf("Error: %v", err)
    }
}