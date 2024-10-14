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
	LOADING           = "LOADING"
	RESULT            = "RESULT"
	CHOOSE_ATACK_TYPE = "CHOOSE_ATACK_TYPE"
	CHOOSE_ATACK      = "CHOOSE_ATACK"
	CHOOSE_DEFENSE    = "CHOOSE_DEFENSE"

	// default
	ATACK        = "ATACK"
	DEFENSE      = "DEFENSE"
	SUBMISSION   = "SUBMISSION"
	SUBMISSIONS  = "SUBMISSIONS"
	PROGRESSION  = "PROGRESSION"
	PROGRESSIONS = "PROGRESSIONS"
	POSITIONS    = "POSITIONS"

	// submissions
	KIMURA          = "KIMURA"
	AMERICANA       = "AMERICANA"
	OMOPLATA        = "OMOPLATA"
	TRIANGULO       = "TRIANGULO"
	ARMLOCK         = "ARMLOCK"
	KATAGATAME      = "KATAGATAME"
	ESTRANGULAMENTO = "ESTRANGULAMENTO"
	ATAQUE_PE       = "ATAQUE_PE"

	// progressions
	QUEDA           = "QUEDA"
	RASPAGEM        = "RASPAGEM"
	ABERTURA_GUARDA = "ABERTURA_GUARDA"
	PASSAGEM_GUARDA = "PASSAGEM_GUARDA"
	PEGADA_COSTAS   = "PEGADA_COSTAS"
	JOELHO_BARRIGA  = "JOELHO_BARRIGA"

	// guards
	FECHADA = "FECHADA"
	ONE_LEG = "ONE_LEG"

	// positions
	GUARDA        = "GUARDA"
	MONTADA       = "MONTADA"
	COSTAS        = "COSTAS"
	CEM_QUILOS    = "CEM_QUILOS"
	QUATRO_APOIOS = "QUATRO_APOIOS"
)

type Option struct {
	label string
	value string
}

var ATACKS_TYPES = []Option{
	{label: "Todas as opções", value: ""},
	{label: "Submissões (Kimura, Armlock, etc)", value: SUBMISSIONS},
	{label: "Progressões (Raspagem, Passagem, etc)", value: PROGRESSIONS},
}

var ATACKS = map[string][]Option{
	SUBMISSIONS: {
		{label: "Todas as opções", value: ""},
		{label: "Kimura", value: KIMURA},
		{label: "Americana", value: AMERICANA},
		{label: "Omoplata", value: OMOPLATA},
		{label: "Triângulo", value: TRIANGULO},
		{label: "Armlock", value: ARMLOCK},
		{label: "Katagatame", value: KATAGATAME},
		{label: "Estrangulamento", value: ESTRANGULAMENTO},
		{label: "Ataque de pé", value: ATAQUE_PE},
	},
	PROGRESSIONS: {
		{label: "Todas as opções", value: ""},
		{label: "Queda", value: QUEDA},
		{label: "Raspagem", value: RASPAGEM},
		{label: "Abertura de guarda", value: ABERTURA_GUARDA},
		{label: "Passagem de guarda", value: PASSAGEM_GUARDA},
		{label: "Pegada das costas", value: PEGADA_COSTAS},
		{label: "Joelho na barriga", value: JOELHO_BARRIGA},
	},
}

var DEFENSES = map[string][]Option{
	POSITIONS: {
		{label: "Todas as opções", value: ""},
		{label: "Montada", value: MONTADA},
		{label: "Costas", value: COSTAS},
		{label: "Cem quilos", value: CEM_QUILOS},
		{label: "Quatro apoios", value: QUATRO_APOIOS},
		{label: "Joelho na barriga", value: JOELHO_BARRIGA},
		{label: "Triângulo", value: TRIANGULO},
	},
}

type Content struct {
	name        string
	description string
	url         string
	categories  map[string]bool
}

const (
	dotChar = " • "
)

var (
	title = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#cde291"))
	atackTitle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#111111")).
			Background(lipgloss.Color("#C2DBE9")).
			PaddingRight(1).
			PaddingLeft(1)
	defenseTitle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#111111")).
			Background(lipgloss.Color("#ec96a7")).
			PaddingRight(1).
			PaddingLeft(1)
	subtitle      = lipgloss.NewStyle().Bold(true)
	description   = lipgloss.NewStyle().Width(70).Foreground(lipgloss.Color("#B1B1B1"))
	subtleStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#B1B1B1"))
	checkboxStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#c1dd71"))
	dotStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("236")).Render(dotChar)
	mainStyle     = lipgloss.NewStyle().MarginLeft(2)
)

var VIEWS = map[string]string{
	CHOOSE_ATACK_TYPE: CHOOSE_ATACK_TYPE,
	CHOOSE_ATACK:      CHOOSE_ATACK,
	CHOOSE_DEFENSE:    CHOOSE_DEFENSE,
	LOADING:           LOADING,
	RESULT:            RESULT,
}

type model struct {
	CurrentView  string
	AtackFound   Content
	DefenseFound Content
	AtackType    int
	Atack        int
	Defense      int
}

func main() {
	initialModel := model{CHOOSE_ATACK_TYPE, Content{}, Content{}, 0, 0, 0}
	p := tea.NewProgram(initialModel)
	if _, err := p.Run(); err != nil {
		fmt.Println("could not start program:", err)
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func forceUpdate(msg tea.Msg) tea.Cmd {
	return func() tea.Msg {
		time.Sleep(600 * time.Millisecond)
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
	case VIEWS[CHOOSE_ATACK_TYPE]:
		return updateChooseAtackType(msg, m)
	case VIEWS[CHOOSE_ATACK]:
		return updateChooseAtack(msg, m)
	case VIEWS[CHOOSE_DEFENSE]:
		return updateChooseDefense(msg, m)
	case VIEWS[LOADING]:
		m.CurrentView = VIEWS[RESULT]
		return m, tea.Quit
	default:
		return m, nil
	}
}

// The main view, which just calls the appropriate sub-view
func (m model) View() string {
	var s string

	switch m.CurrentView {
	case VIEWS[CHOOSE_ATACK_TYPE]:
		s = chooseAtackTypeView(m)
	case VIEWS[CHOOSE_ATACK]:
		s = chooseAtackView(m)
	case VIEWS[CHOOSE_DEFENSE]:
		s = chooseDefenseView(m)
	case VIEWS[LOADING]:
		s = "Encontrando resultados..."
	default:
		s = resultView(m)
	}
	return mainStyle.Render("\n" + s + "\n\n")
}

func findResults(m model) model {
	atacks := []Content{}
	defenses := []Content{}

	for _, row := range database {
		atackKeyToSearch := ATACK

		hasSpecificAtack := m.Atack != 0
		if hasSpecificAtack {
			atackType := ATACKS_TYPES[m.AtackType].value
			atackKeyToSearch = ATACKS[atackType][m.Atack].value
		}

		hasAtackType := m.AtackType != 0
		if !hasSpecificAtack && hasAtackType {
			atackKeyToSearch = ATACKS_TYPES[m.AtackType].value
		}

		isSelectedAtack := row.categories[atackKeyToSearch] && row.categories[ATACK]
		if isSelectedAtack {
			atacks = append(atacks, row)
		}

		defenseKeyToSearch := DEFENSE

		hasSpecificDefense := m.Defense != 0
		if hasSpecificDefense {
			defenseKeyToSearch = DEFENSES[POSITIONS][m.Defense].value
		}

		isSelectedDefense := row.categories[defenseKeyToSearch]
		if isSelectedDefense {
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

	m.AtackFound = atackPick
	m.DefenseFound = defensePick

	return m
}

// ANCHOR: Sub-update functions

func updateChooseAtackType(msg tea.Msg, m model) (tea.Model, tea.Cmd) {
	return handleChoices(msg, m, ATACKS_TYPES)
}

func updateChooseAtack(msg tea.Msg, m model) (tea.Model, tea.Cmd) {
	return handleChoices(msg, m, ATACKS[ATACKS_TYPES[m.AtackType].value])
}

func updateChooseDefense(msg tea.Msg, m model) (tea.Model, tea.Cmd) {
	return handleChoices(msg, m, DEFENSES[POSITIONS])
}

func handleChoices(msg tea.Msg, m model, options []Option) (tea.Model, tea.Cmd) {
	field := &m.AtackType
	switch m.CurrentView {
	case VIEWS[CHOOSE_ATACK_TYPE]:
		field = &m.AtackType
	case VIEWS[CHOOSE_ATACK]:
		field = &m.Atack
	case VIEWS[CHOOSE_DEFENSE]:
		field = &m.Defense
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		max := len(options) - 1
		switch msg.String() {
		case "j", "down":
			*field++
			if *field > max {
				*field = max
			}
		case "k", "up":
			*field--
			if *field < 0 {
				*field = 0
			}
		case "enter":
			switch m.CurrentView {
			case VIEWS[CHOOSE_ATACK_TYPE]:
				if m.AtackType == 0 {
					m.CurrentView = VIEWS[CHOOSE_DEFENSE]
				} else {
					m.CurrentView = VIEWS[CHOOSE_ATACK]
				}
				return m, nil
			case VIEWS[CHOOSE_ATACK]:
				m.CurrentView = VIEWS[CHOOSE_DEFENSE]
				return m, nil
			case VIEWS[CHOOSE_DEFENSE]:
				m = findResults(m)
				m.CurrentView = VIEWS[LOADING]
				return m, forceUpdate(msg)
			default:
				return m, nil
			}
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
%v
Vídeo: %v

%v

%v
%v
Vídeo: %v
`
	filledString := fmt.Sprintf(
		content,
		title.Render("Resultados:"),
		atackTitle.Render("Ataque"),
		subtitle.Render(m.AtackFound.name),
		description.Render(m.AtackFound.description),
		m.AtackFound.url,
		defenseTitle.Render("Defesa"),
		subtitle.Render(m.DefenseFound.name),
		description.Render(m.DefenseFound.description),
		m.DefenseFound.url,
	)
	return fmt.Sprint(filledString)
}

func chooseAtackTypeView(m model) string {
	options := ATACKS_TYPES
	return buildChoices(
		m.AtackType,
		atackTitle.Render("ATAQUE"),
		subtitle.Render("Qual tipo de ataque você busca?"),
		options,
	)
}

func chooseAtackView(m model) string {
	options := ATACKS[ATACKS_TYPES[m.AtackType].value]
	return buildChoices(
		m.Atack,
		atackTitle.Render("ATAQUE"),
		subtitle.Render("Qual ATAQUE você busca?"),
		options,
	)
}

func chooseDefenseView(m model) string {
	options := DEFENSES[POSITIONS]
	return buildChoices(
		m.Defense,
		defenseTitle.Render("DEFESA"),
		subtitle.Render("Defesa a partir de qual posição?"),
		options,
	)
}

func buildChoices(field int, title string, subtitle string, options []Option) string {
	tpl := title + "\n\n" + subtitle + "\n\n%s\n\n"
	tpl += subtleStyle.Render("j/k, up/down: selecionar") + dotStyle +
		subtleStyle.Render("enter: escolher") + dotStyle +
		subtleStyle.Render("q, esc: sair")
	idx := 0
	replacements := ""
	questions := []interface{}{}

	for _, option := range options {
		replacements += "%v\n"
		questions = append(questions, checkbox(option.label, field == idx))
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
