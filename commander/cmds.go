package commander

import "log"

type Commander interface {
	Command(cmd string) (string, error)
}

var RegisteredCmds = map[string]Commander{}

func Register(name string, cmd Commander) {
	if cmd == nil {
		log.Panicf("%v is nil for %v", cmd, name)
	}
	_, exist := RegisteredCmds[name]
	if exist {
		log.Println("register error for %v, already exist", name)
	}
	RegisteredCmds[name] = cmd
}
