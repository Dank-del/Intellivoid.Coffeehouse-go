package posTagging

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	cf "github.com/Dank-del/Intellivoid.Coffeehouse-go/coffeehouse"
)

func TagPOSFull(inp, lang, sen string) (data *POSApiResponse, err error) {
	if len(inp) == 0 {
		err = errors.New("[NLP][POSTagging] input not provided")
		return nil, err
	}

	if len(lang) == 0 {
		lang = cf.DefaultLang
	}
	if len(sen) == 0 {
		sen = cf.DefaultIndex
	}

	key := url.QueryEscape(cf.GetKey())
	inp = url.QueryEscape(inp)
	lang = url.QueryEscape(lang)
	sen = url.QueryEscape(sen)

	url := fmt.Sprintf(endpointurl, key, inp, lang, sen)

	resp, err := http.Post(url, cf.ContentType, nil)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	log.Println(string(b))

	data = new(POSApiResponse)

	err = json.Unmarshal(b, data)
	if err != nil {
		return nil, err
	}

	return
}
