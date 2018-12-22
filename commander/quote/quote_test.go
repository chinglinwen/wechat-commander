package quote

import (
	"testing"
)

func TestGetQuote(t *testing.T) {
	q, err := GetQuote()
	if err != nil {
		t.Error(err)
	}
	t.Log("q:", q)
}
