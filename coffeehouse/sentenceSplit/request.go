/*
 * This file is part of Intellivoid.Coffeehouse-go (https://github.com/Dank-del/Intellivoid.Coffeehouse-go).
 * Copyright (c) 2021 Sayan Biswas.
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

package sentenceSplit

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Dank-del/Intellivoid.Coffeehouse-go/coffeehouse"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func DoRequest(Sentence string) (result *SentenceSplitResponse, err error){
	if len(Sentence) == 0 {
		err = errors.New("[NLP][SentenceSplit] input not provided")
		return nil, err
	}
	requestBody := strings.NewReader(fmt.Sprintf(`
		{
			"input": %s,
		}
	`, Sentence))

	// post some data
	req, err := http.NewRequest("POST", endpointurl,
		requestBody,
	)
	// check for response error
	if err != nil {
		log.Fatal( err )
	}
	req.Header.Set("access_key", coffeehouse.CofeeHouseAPIKey)

	res, err := http.DefaultClient.Do(req)
	if err != nil{
		log.Println(err.Error())
	}
	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response SentenceSplitResponse

	err = json.Unmarshal(b, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}