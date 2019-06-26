package basic

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/chinglinwen/wxrobot-backend/commander"
)

type Sh struct {
}

func (*Sh) Command(cmd string) (out string, err error) {
	s := strings.TrimPrefix(cmd, "sh")
	c := exec.Command("bash", "-l", "-c", s)

	output, err := c.CombinedOutput()
	if err != nil {
		err = fmt.Errorf("execute cmds: %v\noutput: %v", err, string(output))
		return
	}
	out = string(output)

	return
}

func (*Sh) Help() string {
	return "execute shell commands"
}

func init() {
	commander.Register("sh", &Sh{})
}
