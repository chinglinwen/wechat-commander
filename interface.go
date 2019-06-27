package main

import (
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/chinglinwen/wxrobot-backend/commander"
)

//from wechat
/*
{
  "IsGroup": false,
  "MsgId": "3451551602901927588",
  "Content": "disable",
  "FromUserName": "@a99651a071b3adfe9d4fea18915cb09e",
  "ToUserName": "@fe447f00f7ef71089b35244b706fcbd22e9ed44855bfa6fc7b3dba19ff5ee8bc",
  "Who": "@a99651a071b3adfe9d4fea18915cb09e",
  "MsgType": 1,
  "SubType": 0,
  "OriginContent": "disable",
  "At": "",
  "Url": "",
  "RecommendInfo": null
}

*/
const (
	textRobot = "Hi there!"
)

type Handler interface {
	Reply() (string, error)
}

type Ask struct {
	Cmd     string
	Body    string
	From    string
	IsGroup bool
}

type Reply struct {
	Type string
	Data string
}

func NewAsk(from, cmd string) *Ask {
	return &Ask{Cmd: cmd, From: from}
}

// func NewAsk(body, cmd, from string) *Ask {
// 	isgroup := false
// 	if cmd == "" {
// 		cmd = gjson.Get(body, "Content").String()
// 		isgroup = gjson.Get(body, "IsGroup").Bool()
// 	}
// 	cmd = strings.ToLower(cmd)
// 	return &Ask{Body: body, Cmd: cmd, From: from, IsGroup: isgroup}
// }

func formatCheck(cmd, from string) bool {
	return regexp.MustCompile(`^/`).MatchString(cmd) // && from == fromwechat
}

type Helper interface {
	Help() string
}

func GenHelp() string {
	data := "list of commands:\n"
	for name, v := range commander.RegisteredCmds {
		var help string
		if h, ok := v.(Helper); ok {
			help = h.Help()
		}
		data += fmt.Sprintf("/%v  %v\n", name, help)
	}
	return data
}

// list of commands
// /ping
// /k8s
func Process(cmd string) (kind, data string, err error) {
	kind = "text"

	c := strings.Fields(cmd)[0]
	cmdargs := strings.TrimPrefix(cmd, c)

	if c == "help" {
		data = GenHelp()
		return
	}

	found := false
	for name, v := range commander.RegisteredCmds {
		if c != name {
			continue
		}
		found = true
		log.Printf("got Commander %v for args %v\n", name, cmdargs)
		data, err = v.Command(cmdargs)
		if err != nil {
			log.Printf("do cmd: %v err: %v\n", cmdargs, err)
			data = err.Error()
			return
		}
	}
	if !found {
		data = fmt.Sprintf("unknown command: %v\n%v", c, GenHelp())
		return
	}

	return
}

func found(data string) bool {
	return data != ""
}

func match(cmd string, words ...string) bool {
	for _, word := range words {
		if strings.Contains(cmd, word) {
			return true
		}
	}
	return false
}

func (t *Ask) Reply() (reply string, err error) {
	if !formatCheck(t.Cmd, t.From) {
		log.Println("ignore cmd", t.Cmd)
		return
	}
	t.Cmd = strings.TrimPrefix(t.Cmd, "/")
	return encode(Process(t.Cmd))
}

func encode(kind, data string, err error) (string, error) {
	var errtext string
	if err != nil {
		errtext = err.Error()
	}
	b, err := json.MarshalIndent(&struct {
		Type  string `json:"type"`
		Data  string `json:"data"`
		Error string `json:"error"`
	}{
		Type:  kind,
		Data:  data,
		Error: errtext,
	}, "", "  ")
	return string(b), err
}
