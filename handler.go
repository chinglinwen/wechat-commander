package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	// where does msg from
	fromwechat = "wechat"
	fromcmd    = "cmd"
)

// curl localhost:4000 -F from=me -F cmd=help
func handler(w http.ResponseWriter, r *http.Request) {
	ip := r.RemoteAddr
	// fmt.Printf("r: %#v\n", r)
	cmd := r.FormValue("cmd")
	from := r.FormValue("from")
	if cmd == "" {
		cmd = "empty"
	}
	log.Printf("from %v(ip: %v), cmd: %v", from, ip, cmd)

	reply, err := NewAsk(from, cmd).Reply()
	if err != nil {
		fmt.Fprintln(w, "internal error: ", err.Error())
		log.Println("error: ", err.Error())
		return
	}

	// replyType := gjson.Get(string(reply), "type").String()
	// replyData := gjson.Get(string(reply), "data").String()
	// replyErr := gjson.Get(string(reply), "error").String()
	// var n int
	// if len(replyData) < 10 {
	// 	n = len(replyData)
	// }
	// log.Printf("results type: %v, len: %v, data: %v, err: %v\n",
	// 	replyType, len(replyData), replyData[0:n], replyErr)

	fmt.Fprintln(w, reply)
}

// // handler received command from wxrobot.
// func handler(w http.ResponseWriter, r *http.Request) {
// 	body, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		log.Printf("error reading body: %v\n", err)
// 		http.Error(w, "can't read body\n", http.StatusBadRequest)
// 		return
// 	}
// 	log.Println("wechat:", string(body))

// 	reply, err := NewAsk(string(body), "", fromwechat).Reply()
// 	if err != nil {
// 		fmt.Fprintln(w, "internal error: ", err.Error())
// 		log.Println("error: ", err.Error())
// 		return
// 	}

// 	replyType := gjson.Get(string(reply), "type").String()
// 	replyData := gjson.Get(string(reply), "data").String()
// 	replyErr := gjson.Get(string(reply), "error").String()
// 	var n int
// 	if len(replyData) < 10 {
// 		n = len(replyData)
// 	}
// 	log.Printf("results type: %v, len: %v, data: %v, err: %v\n",
// 		replyType, len(replyData), replyData[0:n], replyErr)

// 	fmt.Fprintln(w, reply)
// }

// // cmdHandler receive command from the web.
// func cmdHandler(w http.ResponseWriter, r *http.Request) {
// 	cmd := r.FormValue("cmd")
// 	if cmd == "" {
// 		cmd = "empty"
// 	}
// 	log.Println("cmd:", cmd)
// 	reply, err := NewAsk(fromcmd, cmd).Reply()
// 	if err != nil {
// 		fmt.Fprintln(w, "internal error: ", err.Error())
// 		log.Println("error: ", err.Error())
// 		return
// 	}

// 	replyType := gjson.Get(string(reply), "type").String()
// 	replyData := gjson.Get(string(reply), "data").String()
// 	replyErr := gjson.Get(string(reply), "error").String()
// 	log.Println("replyfields:", replyType, replyData, replyErr)

// 	log.Println("reply:", reply)
// 	fmt.Fprint(w, string(reply))
// }

// // textHandler actively send text to wxrobot.
// // curl localhost:4000/text -F name="休比" -F text=hello
// //
// // this may need valid check ( specific group, and id limit )
// func textHandler(w http.ResponseWriter, r *http.Request) {
// 	name := r.FormValue("name")
// 	text := r.FormValue("text")
// 	if name == "" || text == "" {
// 		fmt.Fprintln(w, "name or text form value is empty")
// 		return
// 	}
// 	log.Println("try send to name: ", name, " text: ", text)

// 	reply, err := sendText(name, text)
// 	if err != nil {
// 		fmt.Fprint(w, "error: ", err.Error())
// 		log.Println("error: ", err.Error())
// 		return
// 	}
// 	log.Printf("reply: %v\n", reply)
// 	fmt.Fprintln(w, reply)
// }
