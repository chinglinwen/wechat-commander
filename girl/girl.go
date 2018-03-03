package girl

import (
	"io/ioutil"
	"net/http"
)

func Pic() (pic []byte, err error) {
	/*
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		uri := "http://www.mzitu.com/zipai/"
		s, err := spider.CreateSpiderFromUrl(uri)
		if err != nil {
			return
		}
		srcs, _ := s.GetAttr("div.main>div.main-content>div.postlist>div>ul>li>div.comment-body>p>img", "src")
		if len(srcs) < 1 {
			err = fmt.Errorf("cannot get mzitu images")
			return
		}
		img := srcs[r.Intn(len(srcs))]
		fmt.Println("img", img)
	*/
	img := "http://wx2.sinaimg.cn/mw1024/9d52c073gy1foxoszeu10j20sg0zkk4y.jpg"
	resp, err := http.Get(img)
	if err != nil {
		return
	}
	return ioutil.ReadAll(resp.Body)
}
