package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}
	reply, _ := NewTextReply(string(body), "").Reply()
	fmt.Fprintln(w, reply)

	fmt.Println("body:", string(body))
}

func cmdHandler(w http.ResponseWriter, r *http.Request) {
	cmd := r.FormValue("cmd")
	if cmd == "" {
		cmd = "empty"
	}
	reply, _ := NewTextReply("", cmd).Reply()
	fmt.Fprintln(w, reply)
	fmt.Println("cmd:", cmd)

}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/ui", cmdHandler)
	log.Fatal(http.ListenAndServe(":4000", nil))
}
