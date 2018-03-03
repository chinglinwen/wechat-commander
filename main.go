package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/tidwall/gjson"
)

func handler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}
	fmt.Println("wechat:", string(body))

	reply, err := NewTextReply(string(body), "").Reply()
	if err != nil {
		fmt.Fprint(w, "internal error: ", err.Error())
		return
	}

	replyType := gjson.Get(string(reply), "type").String()
	replyData := gjson.Get(string(reply), "data").String()
	replyErr := gjson.Get(string(reply), "error").String()
	var n int
	if len(replyData) < 10 {
		n = len(replyData)
	}
	fmt.Printf("results type: %v, len: %v, data: %v, err: %v", replyType, len(replyData), replyData[0:n], replyErr)

	fmt.Fprint(w, reply)
}

func cmdHandler(w http.ResponseWriter, r *http.Request) {
	cmd := r.FormValue("cmd")
	if cmd == "" {
		cmd = "empty"
	}
	fmt.Println("cmd:", cmd)
	reply, err := NewTextReply("", cmd).Reply()
	if err != nil {
		fmt.Fprint(w, "internal error: ", err.Error())
		return
	}

	replyType := gjson.Get(string(reply), "type").String()
	replyData := gjson.Get(string(reply), "data").String()
	replyErr := gjson.Get(string(reply), "error").String()
	fmt.Println("replyfields:", replyType, replyData, replyErr)
	fmt.Println("reply:", reply)

	fmt.Fprint(w, string(reply))
}

func main() {
	log.Println("starting...")
	http.HandleFunc("/", handler)
	http.HandleFunc("/ui", cmdHandler)
	log.Fatal(http.ListenAndServe(":4000", nil))
}
