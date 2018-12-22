package quote

import (
	"encoding/json"
	"fmt"

	"github.com/chinglinwen/wxrobot-backend/commander"
	"github.com/go-resty/resty"
)

type QuoteResult []struct {
	Quote    string `json:"quote"`
	Author   string `json:"author"`
	Category string `json:"category"`
}

var (
	url = "https://andruxnet-random-famous-quotes.p.mashape.com/?cat=movies&count=10"
	key = "dKHmQTWJM0mshA6NKKPhOt9zboacp1AABODjsnTbTPMv5KhafO"
)

func GetQuote() (QuoteResult, error) {
	resp, err := resty.R().
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetHeader("X-Mashape-Key", key).
		SetHeader("Accept", "application/json").
		Post(url)
	if err != nil {
		return nil, err
	}
	var q QuoteResult
	err = json.Unmarshal(resp.Body(), &q)
	return q, nil

}

type Quote struct {
	quotes QuoteResult
	i      int
}

func (b *Quote) Command(cmd string) (data string, err error) {
	if cmd != "quote" {
		return
	}
	//log.Printf("got cmd %v from quote", cmd)
	if b.i == 0 || b.i > 10 {
		//get another 10 quotes
		b.quotes, err = GetQuote()
		b.i = 1
	}
	data = fmt.Sprintf("%v\n  --%v\n", b.quotes[b.i].Quote, b.quotes[b.i].Author)
	b.i += 1
	return
}

func init() {
	commander.Register("quote", &Quote{})
}
