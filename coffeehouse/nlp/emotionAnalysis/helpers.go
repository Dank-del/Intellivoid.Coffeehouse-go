/*
 * This file is part of Intellivoid.Coffeehouse-go (https://github.com/Dank-del/Intellivoid.Coffeehouse-go).
 * Copyright (c) 2021 Sayan Biswas, ALiwoto.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, version 3.
 *
 * This program is distributed in the hope that it will be useful, but
 * WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU
 * General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program. If not, see <http://www.gnu.org/licenses/>.
 */

package emotionAnalysis

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	cf "github.com/Dank-del/Intellivoid.Coffeehouse-go/coffeehouse"
)

// AnalysisText will analysis a text and return the data
// recieved from Intellivoid's server.
// this function allows you to predict the emotions of a given text, and
// it will return an error if something went wrong.
func AnalysisText(text string) (data *Emotionysis, err error) {
	return AnalysisFull(text, cf.DefaultLang,
		cf.DefaultIndex, cf.DefaultIndex,
		cf.DefaultIndex, cf.DefaultIndex)
}

//
func AnalysisFull(inp, lang, sen, gen, genSize, genId string) (data *Emotionysis, err error) {
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
	if len(genSize) == 0 {
		genSize = cf.DefaultIndex
	}
	if len(genId) == 0 {
		genId = cf.DefaultIndex
	}
	if len(gen) == 0 {
		gen = cf.DefaultIndex
	}

	key := url.QueryEscape(cf.GetKey())
	inp = url.QueryEscape(inp)
	lang = url.QueryEscape(lang)
	sen = url.QueryEscape(sen)
	gen = url.QueryEscape(gen)
	genSize = url.QueryEscape(genSize)
	genId = url.QueryEscape(genId)

	url := fmt.Sprintf(endpointurl, key, inp, lang, sen, gen, genSize, genId)

	res, err := http.Post(url, cf.ContentType, nil)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	data = new(Emotionysis)

	err = json.Unmarshal(b, data)
	if err != nil {
		return nil, err
	}

	return
}
