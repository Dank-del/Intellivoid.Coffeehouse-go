package posTagging

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	cf "github.com/Dank-del/Intellivoid.Coffeehouse-go/coffeehouse"
)

func TagPOSFull(inp, lang, sen string) (pos *POSApiResponse, err error) {
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

	dt := url.Values{}
	dt.Set(accessKeyKey, cf.GetKey())
	dt.Set(inputKey, inp)
	dt.Set(languageKey, lang)
	dt.Set(sentenceSplitKey, sen)
	resp, err := http.PostForm(endpointurl, dt)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	log.Println(string(b))

	pos = new(POSApiResponse)

	err = json.Unmarshal(b, pos)
	if err != nil {
		return nil, err
	}

	return
}
