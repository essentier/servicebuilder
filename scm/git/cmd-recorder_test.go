package git_test

import (
	"strings"

	"github.com/essentier/nomockutil/cmd"
)

func createRecordCmd() *recordCmd {
	return &recordCmd{CmdRunner: cmd.CreateCmdConsole()}
}

type recordCmd struct {
	cmd.CmdRunner
	lastCommand string
}

func (c *recordCmd) NewRunner() cmd.CmdRunner {
	return createRecordCmd()
}

func (c *recordCmd) RunCmd(name string, args ...string) (string, cmd.CmdError) {
	c.lastCommand = name + " " + strings.Join(args[:], " ")
	return "", c.LastError()
}

func (c *recordCmd) RunInNewRunner(name string, args ...string) (string, cmd.CmdError) {
	newRunner := c.NewRunner()
	return newRunner.RunCmd(name, args...)
}
