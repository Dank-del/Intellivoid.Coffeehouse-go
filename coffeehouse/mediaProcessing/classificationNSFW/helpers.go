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

package classificationNSFW

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	url2 "net/url"
	"strings"

	"github.com/Dank-del/Intellivoid.Coffeehouse-go/coffeehouse"
)

func toBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}
func DoRequest(filename string) (res *NSFWClassificationResponse, err error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	imgBase64 := toBase64(bytes)

	dt := url2.Values{}
	dt.Set("image", imgBase64)
	dt.Set(accessKeyKey, coffeehouse.GetKey())
	resp, err := http.PostForm(endpointurl, dt)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//req.Header.Set("image", image)
	//req.PostForm = url2.Values{}
	//req.PostForm.Set("image", image)
	//req.Header.Set("Content-Type", "multipart/form-data")
	//resp, err := http.Post(url, coffeehouse.ContentType, nil)
	//resp, err := http.DefaultClient.Do(req)
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

	var n NSFWClassificationResponse

	err = json.Unmarshal(b, &n)

	return &n, err
}
