package girl

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestPic(t *testing.T) {
	data, err := Pic()
	if err != nil {
		t.Error(err)
	}
	err = ioutil.WriteFile("testpic.png", data, 0644)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("see file testpic.png")
}
