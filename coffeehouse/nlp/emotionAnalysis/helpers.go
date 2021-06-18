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
func AnalysisFull(inp, lang, sen, gen, genSize, genId string) (emo *Emotionysis, err error) {
	if !cf.IsSet() {
		return nil, errors.New("[NLP][EmotionAnalysis] access key is not set")
	}

	if len(inp) == 0 {
		err = errors.New("[NLP][EmotionAnalysis] input not provided")
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

	v := url.Values{}
	v.Set(accessKeyKey, cf.GetKey())
	v.Set(inputKey, inp)
	v.Set(languageKey, lang)
	v.Set(sentenceSplitKey, sen)
	v.Set(generalizationKey, gen)
	v.Set(generalizationSizeKey, genSize)
	v.Set(generalizationIdKey, genId)

	// Post the form with the data
	resp, err := http.PostForm(endpointurl, v)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	emo = new(Emotionysis)

	err = json.Unmarshal(b, emo)
	if err != nil {
		return nil, err
	}

	return
}
