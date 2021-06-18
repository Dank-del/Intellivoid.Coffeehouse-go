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

package posTagging

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"

	cf "github.com/Dank-del/Intellivoid.Coffeehouse-go/coffeehouse"
)

func TagPOSFull(inp, lang, sen string) (pos *POSApiResponse, err error) {
	if !cf.IsSet() {
		return nil, errors.New("[NLP][POSTagging] access key is not set")
	}

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

	v := url.Values{}
	v.Set(accessKeyKey, cf.GetKey())
	v.Set(inputKey, inp)
	v.Set(languageKey, lang)
	v.Set(sentenceSplitKey, sen)

	resp, err := http.PostForm(endpointurl, v)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	pos = new(POSApiResponse)

	err = json.Unmarshal(b, pos)
	if err != nil {
		return nil, err
	}

	return
}
