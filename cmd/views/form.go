package views

import (
	"fmt"
	"strconv"

	"github.com/codeaspiras/gotravel/calculator"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type (
	errMsg error
)

const (
	distanceToGo = iota
	distanceToReturn
	costPerLiter
	distancePerLiter
)

const (
	lightBlue = lipgloss.Color("#06DAFF")
	darkGray  = lipgloss.Color("#767676")
)

var (
	inputStyle  = lipgloss.NewStyle().Foreground(lightBlue)
	outputStyle = lipgloss.NewStyle().Foreground(darkGray)
	validInputs = make(map[int]float64)
)

type model struct {
	inputs  []textinput.Model
	focused int
	err     error
}

func numberValidator(s string, key int) error {
	num, err := strconv.ParseFloat(s, 64)
	if err != nil {
		delete(validInputs, key)
		return fmt.Errorf("informe somente números (com ponto no lugar da vírgula)")
	}

	if num <= 0 {
		delete(validInputs, key)
		return fmt.Errorf("informe um valor positivo")
	}

	validInputs[key] = num
	return nil
}

func NewFormModel() model {
	var inputs []textinput.Model = make([]textinput.Model, 4)
	inputs[distanceToGo] = textinput.New()
	inputs[distanceToGo].Placeholder = "123.45"
	inputs[distanceToGo].Focus()
	inputs[distanceToGo].CharLimit = 10
	inputs[distanceToGo].Width = 30
	inputs[distanceToGo].Prompt = ""
	inputs[distanceToGo].Validate = func(s string) error {
		return numberValidator(s, distanceToGo)
	}

	inputs[distanceToReturn] = textinput.New()
	inputs[distanceToReturn].Placeholder = "123.45"
	inputs[distanceToReturn].CharLimit = 10
	inputs[distanceToReturn].Width = 30
	inputs[distanceToReturn].Prompt = ""
	inputs[distanceToReturn].Validate = func(s string) error {
		return numberValidator(s, distanceToReturn)
	}

	inputs[costPerLiter] = textinput.New()
	inputs[costPerLiter].Placeholder = "10.5"
	inputs[costPerLiter].CharLimit = 7
	inputs[costPerLiter].Width = 29
	inputs[costPerLiter].Prompt = ""
	inputs[costPerLiter].Validate = func(s string) error {
		return numberValidator(s, costPerLiter)
	}

	inputs[distancePerLiter] = textinput.New()
	inputs[distancePerLiter].Placeholder = "14.0"
	inputs[distancePerLiter].CharLimit = 6
	inputs[distancePerLiter].Width = 30
	inputs[distancePerLiter].Prompt = ""
	inputs[distancePerLiter].Validate = func(s string) error {
		return numberValidator(s, distancePerLiter)
	}

	return model{
		inputs:  inputs,
		focused: 0,
		err:     nil,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd = make([]tea.Cmd, len(m.inputs))

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			if m.focused == len(m.inputs)-1 {
				return m, tea.Quit
			}
			m.nextInput()
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyShiftTab, tea.KeyCtrlP:
			m.prevInput()
		case tea.KeyTab, tea.KeyCtrlN:
			m.nextInput()
		}
		for i := range m.inputs {
			m.inputs[i].Blur()
		}
		m.inputs[m.focused].Focus()

	case errMsg:
		m.err = msg
		return m, nil
	}

	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}
	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	output := "Preencha os campos..."
	if len(validInputs) == 4 {
		totalDistance := validInputs[distanceToGo] + validInputs[distanceToReturn]
		costToCover := calculator.CostToCover(
			totalDistance,
			validInputs[distancePerLiter],
			validInputs[costPerLiter],
		)
		output = fmt.Sprintf("Resultado: $%.2f para percorrer %.2fkm", costToCover, totalDistance)
	}
	return fmt.Sprintf(
		` Bem-vindo(a) à calculadora de custo de combustível!
 Preencha os campos abaixo para que a conta seja feita!
 
 Observação: Para valores numéricos, se quiser informar uma
 fração, insira somente números e ponto (.) no lugar da vírgula.

 %s  %s
 %s %s

 %s  %s
 $%s %s

 %s
`,
		inputStyle.Width(30).Render("Distância ida (km)"),
		inputStyle.Width(30).Render("Distância volta (km)"),
		m.inputs[distanceToGo].View(),
		m.inputs[distanceToReturn].View(),
		inputStyle.Width(30).Render("Combustível ($/L)"),
		inputStyle.Width(30).Render("Eficiência (km/L)"),
		m.inputs[costPerLiter].View(),
		m.inputs[distancePerLiter].View(),
		outputStyle.Render(output),
	) + "\n"
}

// nextInput focuses the next input field
func (m *model) nextInput() {
	m.focused = (m.focused + 1) % len(m.inputs)
}

// prevInput focuses the previous input field
func (m *model) prevInput() {
	m.focused--
	// Wrap around
	if m.focused < 0 {
		m.focused = len(m.inputs) - 1
	}
}
