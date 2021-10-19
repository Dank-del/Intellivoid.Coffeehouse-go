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

package langDetect

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	cf "github.com/intellivoid/Intellivoid.Coffeehouse-go/coffeehouse"
)

func DetectText(text string) (*LangDetector, error) {
	return DetectFull(text, cf.DefaultLang,
		cf.DefaultIndex, cf.DefaultIndex,
		cf.DefaultIndex, cf.DefaultIndex)
}

func DetectFull(inp, lang, sen, gen, genSize, genId string) (detector *LangDetector, err error) {
	if !cf.IsSet() {
		return nil, errors.New("[NLP][LangDetect] access key is not set")
	}

	if len(inp) == 0 {
		err = errors.New("[NLP][LangDetect] input not provided")
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

	log.Println(string(b))

	detector = new(LangDetector)

	err = json.Unmarshal(b, detector)
	if err != nil {
		return nil, err
	}

	return
}
