package main

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/exp/rand"
)

const (
	// views
	LOADING               = "LOADING"
	RESULT                = "RESULT"
	CHOOSE_ATACK_TYPE     = "CHOOSE_ATACK_TYPE"
	CHOOSE_ATACK          = "CHOOSE_ATACK"
	CHOOSE_ATACK_POSITION = "CHOOSE_ATACK_POSITION"
	CHOOSE_DEFENSE        = "CHOOSE_DEFENSE"

	// default
	ATACK        = "ATACK"
	DEFENSE      = "DEFENSE"
	SUBMISSIONS  = "SUBMISSIONS"
	PROGRESSIONS = "PROGRESSIONS"

	// submissions
	KIMURA          = "KIMURA"
	AMERICANA       = "AMERICANA"
	OMOPLATA        = "OMOPLATA"
	TRIANGULO       = "TRIANGULO"
	ARMBAR          = "ARMBAR"
	KATAGATAME      = "KATAGATAME"
	ESTRANGULAMENTO = "ESTRANGULAMENTO"
	ATAQUE_PE       = "ATAQUE_PE"

	// progressions
	RASPAGEM        = "RASPAGEM"
	ABERTURA_GUARDA = "ABERTURA_GUARDA"
	PASSAGEM_GUARDA = "PASSAGEM_GUARDA"
	PEGADA_COSTAS   = "PEGADA_COSTAS"
	JOELHO_BARRIGA  = "JOELHO_BARRIGA"

	// guards
	FECHADA = "FECHADA"
	ONE_LEG = "ONE_LEG"

	// positions
	GUARDA  = "GUARDA"
	MONTADA = "MONTADA"
)

type Option struct {
	label string
	value string
}

var ATACKS_TYPES = []Option{
	{label: "Submissões (Kimura, Armlock, etc)", value: SUBMISSIONS},
	{label: "Progressões (Raspagem, Passagem, etc)", value: PROGRESSIONS},
}

var ATACKS = map[string][]Option{
	SUBMISSIONS: {
		{label: "Kimura", value: KIMURA},
		{label: "Americana", value: AMERICANA},
		{label: "Omoplata", value: OMOPLATA},
		{label: "Triângulo", value: TRIANGULO},
		{label: "Armbar", value: ARMBAR},
		{label: "Katagatame", value: KATAGATAME},
		{label: "Estrangulamento", value: ESTRANGULAMENTO},
		{label: "Ataque de pé", value: ATAQUE_PE},
	},
	PROGRESSIONS: {
		{label: "Raspagem", value: RASPAGEM},
		{label: "Abertura de guarda", value: ABERTURA_GUARDA},
		{label: "Passagem de guarda", value: PASSAGEM_GUARDA},
		{label: "Pegada das costas", value: PEGADA_COSTAS},
		{label: "Joelho na barriga", value: JOELHO_BARRIGA},
	},
}

type Content struct {
	name       string
	url        string
	technique  string
	categories map[string]bool
}

var database = map[int]Content{
	1: {
		name:      "Kimura trap",
		url:       "https://www.instagram.com/p/C_YZK-1R8Ud/",
		technique: ATACK,
		categories: map[string]bool{
			KIMURA: true,
		},
	},
	2: {
		name:      "Omoplata",
		url:       "https://www.instagram.com/p/C_YZK-1R8Ud/",
		technique: ATACK,
		categories: map[string]bool{
			OMOPLATA: true,
		},
	},
	3: {
		name:      "Americana",
		url:       "https://www.instagram.com/p/C_YZK-1R8Ud/",
		technique: ATACK,
		categories: map[string]bool{
			AMERICANA: true,
		},
	},
	4: {
		name:      "Defesa americana",
		url:       "https://www.instagram.com/p/C_YZK-1R8Ud/",
		technique: DEFENSE,
		categories: map[string]bool{
			AMERICANA: true,
		},
	},
}

const (
	dotChar = " • "
)

var (
	title = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#7D56F4")).
		PaddingRight(1).
		PaddingLeft(1)
	subtitle      = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
	subtleStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
	checkboxStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("212"))
	dotStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("236")).Render(dotChar)
	mainStyle     = lipgloss.NewStyle().MarginLeft(2)
)

var VIEWS = map[string]string{
	LOADING:               LOADING,
	RESULT:                RESULT,
	CHOOSE_ATACK_TYPE:     CHOOSE_ATACK_TYPE,
	CHOOSE_ATACK:          CHOOSE_ATACK,
	CHOOSE_ATACK_POSITION: CHOOSE_ATACK_POSITION,
	CHOOSE_DEFENSE:        CHOOSE_DEFENSE,
}

type model struct {
	CurrentView string
	AtackPick   Content
	DefensePick Content
	AtackType   int
	Atack       int
}

func main() {
	initialModel := model{
		CHOOSE_ATACK_TYPE,
		Content{},
		Content{},
		-1,
		-1,
	}
	p := tea.NewProgram(initialModel)
	if _, err := p.Run(); err != nil {
		fmt.Println("could not start program:", err)
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func forceUpdate(msg tea.Msg, m model) tea.Cmd {
	return func() tea.Msg {
		time.Sleep(800 * time.Millisecond)
		return msg
	}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Make sure these keys always quit
	if msg, ok := msg.(tea.KeyMsg); ok {
		k := msg.String()
		if k == "q" || k == "esc" || k == "ctrl+c" {
			return m, tea.Quit
		}
	}

	switch m.CurrentView {
	case VIEWS[LOADING]:
		m.CurrentView = VIEWS[RESULT]
		return m, nil
	case VIEWS[CHOOSE_ATACK_TYPE]:
		return updateChoseAtackType(msg, m)
	case VIEWS[CHOOSE_ATACK]:
		return updateChoseAtack(msg, m)
	default:
		return m, nil
	}
}

// The main view, which just calls the appropriate sub-view
func (m model) View() string {
	var s string

	switch m.CurrentView {
	case VIEWS[LOADING]:
		s = "Encontrando resultados..."
	case VIEWS[CHOOSE_ATACK_TYPE]:
		s = choseAtackTypeView(m)
	case VIEWS[CHOOSE_ATACK]:
		s = choseAtackView(m)
	default:
		s = resultView(m)
	}
	return mainStyle.Render("\n" + s + "\n\n")
}

func findResults(m model) model {
	atacks := []Content{}
	defenses := []Content{}

	for _, row := range database {
		isAtack := row.technique == ATACK
		if isAtack {
			atacks = append(atacks, row)
		}

		isDefense := row.technique == DEFENSE
		if isDefense {
			defenses = append(defenses, row)
		}
	}

	atackPick := Content{}
	if len(atacks) > 0 {
		atackPick = atacks[rand.Intn(len(atacks))]
	}

	defensePick := Content{}
	if len(defenses) > 0 {
		defensePick = defenses[rand.Intn(len(defenses))]
	}

	m.AtackPick = atackPick
	m.DefensePick = defensePick

	return m
}

// ANCHOR: Sub-update functions

func updateChoseAtackType(msg tea.Msg, m model) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		max := len(ATACKS_TYPES) - 1
		switch msg.String() {
		case "j", "down":
			m.AtackType++
			if m.AtackType > max {
				m.AtackType = max
			}
		case "k", "up":
			m.AtackType--
			if m.AtackType < 0 {
				m.AtackType = 0
			}
		case "enter":
			m.CurrentView = VIEWS[CHOOSE_ATACK]
			return m, nil
		}
	}

	return m, nil
}

func updateChoseAtack(msg tea.Msg, m model) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		options := ATACKS[ATACKS_TYPES[m.AtackType].value]
		max := len(options) - 1
		switch msg.String() {
		case "j", "down":
			m.Atack++
			if m.Atack > max {
				m.Atack = max
			}
		case "k", "up":
			m.Atack--
			if m.Atack < 0 {
				m.Atack = 0
			}
		case "enter":
			m = findResults(m)
			m.CurrentView = VIEWS[LOADING]
			return m, forceUpdate(msg, m)
		}
	}

	return m, nil
}

// ANCHOR: Sub-views

func resultView(m model) string {
	const content = `
%v

%v

%v
Vídeo: %v

%v

%v
Vídeo: %v
`
	filledString := fmt.Sprintf(
		content,
		title.Render("Resultado:"),
		subtitle.Render("# Ataque"),
		m.AtackPick.name,
		m.AtackPick.url,
		subtitle.Render("# Defesa"),
		m.DefensePick.name,
		m.DefensePick.url,
	)

	return fmt.Sprint(filledString)
}

func choseAtackTypeView(m model) string {
	c := m.AtackType

	tpl := "Qual TIPO de ataque você busca?\n\n"
	tpl += "%s\n\n"
	tpl += subtleStyle.Render("j/k, up/down: select") + dotStyle +
		subtleStyle.Render("enter: choose") + dotStyle +
		subtleStyle.Render("q, esc: quit")

	replacements := ""
	questions := []interface{}{}
	idx := 0

	for _, option := range ATACKS_TYPES {
		replacements += "%v\n"
		questions = append(questions, checkbox(option.label, c == idx))
		idx++
	}

	choices := fmt.Sprintf(replacements, questions...)

	return fmt.Sprintf(tpl, choices)
}

func choseAtackView(m model) string {
	c := m.Atack

	tpl := "Qual ATAQUE você busca?\n\n"
	tpl += "%s\n\n"
	tpl += subtleStyle.Render("j/k, up/down: select") + dotStyle +
		subtleStyle.Render("enter: choose") + dotStyle +
		subtleStyle.Render("q, esc: quit")

	replacements := ""
	questions := []interface{}{}
	idx := 0

	options := ATACKS[ATACKS_TYPES[m.AtackType].value]

	for _, option := range options {
		replacements += "%v\n"
		questions = append(questions, checkbox(option.label, c == idx))
		idx++
	}

	choices := fmt.Sprintf(replacements, questions...)

	return fmt.Sprintf(tpl, choices)
}

func checkbox(label string, checked bool) string {
	if checked {
		return checkboxStyle.Render("[x] " + label)
	}
	return fmt.Sprintf("[ ] %s", label)
}
