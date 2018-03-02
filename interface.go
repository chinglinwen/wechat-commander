package main

import (
	"strings"

	"github.com/tidwall/gjson"
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

func NewTextReply(body, cmd string) *TextReply {
	if cmd == "" {
		cmd = gjson.Get(body, "Content").String()
	}
	cmd = strings.ToLower(cmd)
	return &TextReply{Body: body, Cmd: cmd}
}

func (t *TextReply) Reply() (text string, err error) {
	if t.Cmd == "empty" {
		text = "Your command is empty"
		return
	}

	if strings.Contains(t.Cmd, "robot") {
		text = textRobot
		return
	}
	return
	//not found, just skip, always return err=nil
}
