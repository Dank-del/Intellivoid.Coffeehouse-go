/*
 * This file is part of Intellivoid.Coffeehouse-go (https://github.com/intellivoid/Intellivoid.Coffeehouse-go).
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

package spamPrediction

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"

	cf "github.com/intellivoid/Intellivoid.Coffeehouse-go/coffeehouse"
)

func DoRequest(inp, lang, sen, gen, genSize, genId string) (predict *SpamPredictionAPIResponse, err error) {
	if len(inp) == 0 {
		err = errors.New("[NLP][POSTagging] input not provided")
		return
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
	v.Set(generalizationIdKey, genId)
	v.Set(generalizationSizeKey, genSize)

	resp, err := http.PostForm(endpointurl, v)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	predict = new(SpamPredictionAPIResponse)

	err = json.Unmarshal(b, predict)
	if err != nil {
		return nil, err
	}

	return
}
