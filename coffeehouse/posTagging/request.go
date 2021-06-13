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

package posTagging

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

func DoRequest(Input string, Language string, sentenceSplit string) (data *POSApiResponse, err error){
	// request body (payload)
	if len(Input) == 0 {
		err = errors.New("[NLP][POSTagging] input not provided")
		return nil, err
	}
	if len(Language) == 0 {
		Language = "en"
	}
	if len(sentenceSplit) == 0 {
		sentenceSplit = "0"
	}
	requestBody := strings.NewReader(fmt.Sprintf(`
		{
			"input": %s,
			"language":%s,
			"sentence_split":%s
		}
	`, Input, Language, sentenceSplit))

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

	var response POSApiResponse

	err = json.Unmarshal(b, &response)
	if err != nil {
		return nil, err
	}
    fmt.Println(&response)
	return &response, nil
}