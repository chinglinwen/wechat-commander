package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/tidwall/gjson"

	"github.com/chinglinwen/wxrobot-backend/girl"
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

type TextReply struct {
	Cmd  string
	Body string
}

type Reply struct {
	Type string
	Data string
}

func NewTextReply(body, cmd string) *TextReply {
	if cmd == "" {
		cmd = gjson.Get(body, "Content").String()
	}
	cmd = strings.ToLower(cmd)
	return &TextReply{Body: body, Cmd: cmd}
}

func (t *TextReply) Reply() (reply string, err error) {
	var kind = "text"
	var data string

	if t.Cmd == "empty" {
		err = fmt.Errorf("Your command is empty")
		return encode(kind, data, err)
	}

	if match(t.Cmd, "robot", "机器人") {
		data = textRobot
		return encode(kind, data, err)
	}

	if match(t.Cmd, "error", "bug") {
		err = fmt.Errorf("robot is in trouble")
		return encode(kind, data, err)
	}

	if match(t.Cmd, "girl", "美女") {
		kind = "image"
		data, err := girl.Pic()
		if err != nil {
			return encode(kind, "", err)
		}
		encoded := base64.StdEncoding.EncodeToString(data)
		return encode(kind, encoded, err)
	}

	//fmt.Printf("There's no data for this cmd.\n")
	return
	//not found, just skip, always return err=nil
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

func match(cmd string, words ...string) bool {
	for _, word := range words {
		if strings.Contains(cmd, word) {
			return true
		}
	}
	return false
}
