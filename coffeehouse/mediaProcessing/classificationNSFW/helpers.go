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
	"fmt"
	"github.com/Dank-del/Intellivoid.Coffeehouse-go/coffeehouse"
	"io/ioutil"
	"log"
	"net/http"
	url2 "net/url"
)

func toBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}
func DoRequest(filename string) (res *NSFWClassificationResponse, err error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	var base64Encoding string
	mimeType := http.DetectContentType(bytes)

	switch mimeType {
	case "image/jpeg":
		base64Encoding += "data:image/jpeg;base64,"
	case "image/png":
		base64Encoding += "data:image/png;base64,"

	}
	base64Encoding += toBase64(bytes)
	key := url2.QueryEscape(coffeehouse.GetKey())
	image := url2.QueryEscape(base64Encoding)
	url := fmt.Sprintf("%s?=accesskey=%s&image=%s", endpointUrl, key, image)
	resp, err := http.Post(url, coffeehouse.ContentType, nil)
	if err != nil {
		log.Fatal(err)
	}
	d := json.NewDecoder(resp.Body)
	var n NSFWClassificationResponse
	err = d.Decode(&n)
	return &n, err
}
