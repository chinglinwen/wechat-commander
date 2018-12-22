package basic

import (
	"github.com/chinglinwen/wxrobot-backend/commander"
)

type Ping struct {
}

func (b *Ping) Command(cmd string) (string, error) {
	//log.Printf("got cmd %v from ping", cmd)
	return "pong", nil
}

func (b *Ping) Help() string {
	return "do a basic ping"
}

func init() {
	commander.Register("ping", &Ping{})
}
