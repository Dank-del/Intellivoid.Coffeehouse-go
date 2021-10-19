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

package classificationNSFW

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	url2 "net/url"
	"strconv"
	"strings"

	cf "github.com/intellivoid/Intellivoid.Coffeehouse-go/coffeehouse"
	gen "github.com/intellivoid/Intellivoid.Coffeehouse-go/coffeehouse/generalization"
)

func toBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func classify(filename, gen, genSize, genId string) (classified *NSFWClassified, err error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	imgBase64 := toBase64(bytes)

	v := url2.Values{}
	v.Set(imageKey, imgBase64)
	v.Set(accessKeyKey, cf.GetKey())
	v.Set(generalizationKey, gen)
	v.Set(generalizationSizeKey, genSize)

	v.Set(generalizationIdKey, genId)

	resp, err := http.PostForm(endpointurl, v)

	if err != nil {
		str := err.Error()
		str = strings.ReplaceAll(str, imgBase64, "{<base64> value of image}")
		log.Fatal(str)
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(b))

	classified = new(NSFWClassified)

	err = json.Unmarshal(b, classified)

	return
}

func ClassifyFile(filename string) (*NSFWClassified, error) {
	return classify(filename, cf.DefaultIndex,
		cf.DefaultIndex, cf.DefaultIndex)
}

func ClassifyWithGeneralize(filename string,
	gInfo gen.GenInfo) (classified *NSFWClassified, err error) {
	if gInfo == nil {
		return nil, errors.New("generalization info cannot be nil")
	}

	size := strconv.Itoa(int(gInfo.GetSize()))

	classified, err = classify(filename, cf.TrueIndex,
		size, cf.DefaultIndex)

	if classified.Results != nil {
		if classified.Results.Generalization != nil {
			id := classified.Results.Generalization.ID
			// set the generalization id, so it will set
			// isnew to `false`.
			gInfo.SetID(id)
		}
	}

	return
}
