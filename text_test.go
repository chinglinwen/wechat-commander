package main

import (
	"testing"
)

func TestSendText(t *testing.T) {
	close := conn()
	defer close()

	r, err := sendText("wen", "data hello")
	if err != nil {
		t.Error("err: ", err)
	}
	t.Log(r)
}
