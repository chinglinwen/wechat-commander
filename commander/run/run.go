package basic

import (
	"fmt"
	"os/exec"

	"github.com/chinglinwen/wxrobot-backend/commander"
)

type Run struct {
}

func (*Run) Command(cmd string) (out string, err error) {
	s := "./run.sh " + cmd
	c := exec.Command("bash", "-l", "-c", s)
	c.Dir = "/home/wen/git/commanders"
	output, err := c.CombinedOutput()
	if err != nil {
		err = fmt.Errorf("execute cmds: %v\noutput: %v", err, string(output))
		return
	}
	out = string(output)

	return
}

func (*Run) Help() string {
	return "run commanders from git"
}

func init() {
	commander.Register("run", &Run{})
}
