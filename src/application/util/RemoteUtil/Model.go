package RemoteUtil

import (
	tea "github.com/charmbracelet/bubbletea"
	"os/exec"
	"script-go/src/application/pojo/dto/Log"
	"script-go/src/application/util/GenUtil"
	"script-go/src/application/util/LogUtil"
)

type execModel struct {
	err  error
	bin  string
	args []string
}

func newExecModel(bin string, args []string) *execModel {
	return &execModel{bin: bin, args: args}
}

func (m *execModel) Init() tea.Cmd {
	return tea.ExecProcess(
		exec.Command(m.bin, m.args...),
		func(err error) tea.Msg {
			LogUtil.LoggerLine(Log.Of("Model", "Init", "tea.ExecProcess", err))
			return tea.Quit()
		},
	)
}

func (m *execModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		LogUtil.LoggerLine(Log.Of("execModel", "Update", "KeyMsg", msg.String()))
	default:
		if GenUtil.IsEmpty(msg) {
			LogUtil.LoggerLine(Log.Of("execModel", "Update", "default msg", "quit"))
			tea.Quit()
		}
	}
	return m, nil
}

func (m *execModel) View() string {
	return ""
}
