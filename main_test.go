package main

import (
	"testing"
)

var body = `
{
  "IsGroup": false,
  "MsgId": "8927133500120292699",
  "Content": "Robot",
  "FromUserName": "@a99651a071b3adfe9d4fea18915cb09e",
  "ToUserName": "@fe447f00f7ef71089b35244b706fcbd22e9ed44855bfa6fc7b3dba19ff5ee8bc",
  "Who": "@a99651a071b3adfe9d4fea18915cb09e",
  "MsgType": 1,
  "SubType": 0,
  "OriginContent": "robot",
  "At": "",
  "Url": "",
  "RecommendInfo": null
}
`

func TestTextReply(t *testing.T) {
	reply, _ := NewTextReply("", "robot").Reply()
	//spew.Dump(reply)
	if reply != textRobot {
		t.Error("reply incorrect")
	}
}
