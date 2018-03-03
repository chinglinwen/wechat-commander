package main

import (
	"testing"
)

func TestMatch(t *testing.T) {
	if !match("robot", "robot") {
		t.Error("robot match robot failed")
	}
}
