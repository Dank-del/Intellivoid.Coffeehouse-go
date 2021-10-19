package tests

import (
	"log"
	"testing"

	"github.com/intellivoid/Intellivoid.Coffeehouse-go/coffeehouse"
	"github.com/intellivoid/Intellivoid.Coffeehouse-go/coffeehouse/nlp/langDetect"
)

func TestLangDetect(t *testing.T) {
	coffeehouse.SetKey("<access_key>")
	res, err := langDetect.DetectText("こんにちわ")
	if err != nil {
		log.Fatal(err.Error())
		t.Log(err)
	}

	if res.Success != true {
		t.Log(res)
		t.Errorf("[Intellivoid.Coffeehouse-go (langDetect)] Failed request, response code: %d", res.ResponseCode)

	}
}
