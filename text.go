package main

import (
	"context"
	"time"

	pb "github.com/chinglinwen/wxrobot/api"
)

func sendText(name, text string) (msg string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	r, err := client.Text(ctx, &pb.TextRequest{Name: name, Text: text})
	if err != nil {
		return
	}
	msg = r.Data
	return
}
