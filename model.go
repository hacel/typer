package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/reflow/wordwrap"
)

func getText() tea.Msg {
	res, err := http.Get("https://hasel.xyz/texts")
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("error getting text from server: %s", res.Status)
	}
	defer res.Body.Close()
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	return string(b)
}

var (
	green = lipgloss.NewStyle().Background(lipgloss.Color("#009900"))
	red   = lipgloss.NewStyle().Background(lipgloss.Color("#990000"))
)

type model struct {
	width     int
	textInput textinput.Model
	text      string
	c         int
	i         int
	start     time.Time
}

func initialModel() model {
	ti := textinput.NewModel()
	ti.Prompt = ""
	ti.CharLimit = 1000
	ti.Width = 20
	ti.Focus()

	return model{
		textInput: ti,
	}
}

func (m model) Init() tea.Cmd {
	return getText
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.text != "" && m.c == len(m.text) {
		return m, tea.Quit
	}

	var cmd tea.Cmd
	switch msg := msg.(type) {
	case string:
		m.text = msg
		m.start = time.Now()
		return m, nil

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.textInput.Width = msg.Width

	case tea.KeyMsg:
		if m.text == "" {
			return m, nil
		}
		switch msg.Type {
		case tea.KeyEnter, tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit

		case tea.KeyBackspace:
			if m.i == m.c {
				m.c--
				if m.c < 0 {
					m.c = 0
				}
			}
			m.i--
			if m.i < 0 {
				m.i = 0
			}

		default:
			if m.i == m.c && msg.String() == string(m.text[m.c]) {
				m.c++
			}
			m.i++
		}
		m.textInput, cmd = m.textInput.Update(msg)
		return m, cmd

	case error:
		fmt.Fprintln(os.Stderr, msg)
		return m, tea.Quit
	}
	return m, cmd
}

func (m model) View() string {
	if m.text == "" {
		return ""
	}

	if m.c == len(m.text) {
		words := strings.Count(m.text, " ") + 1
		wpm := float64(words) / time.Since(m.start).Minutes()
		return fmt.Sprintf("%s\n\nWords: %d, WPM: %.f\n", wordwrap.String(m.text, m.width), words, wpm)
	}

	i := m.i
	if i > len(m.text) {
		i = len(m.text)
	}
	correct := m.text[:m.c]
	wrong := m.text[m.c:i]
	rest := m.text[i:]
	text := wordwrap.String(green.Render(correct)+red.Render(wrong)+rest, m.width)
	return text + "\n\n" + m.textInput.View()
}
