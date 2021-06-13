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

package sentimentAnalysis

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

func DoRequest(Input string, Language string, sentenceSplit string, Generalize string, GeneralizationSize string, GeneralizationID string) (data *SentimentAnalysisResponse, err error){
	// request body (payload)
	if len(Language) == 0 {
		err = errors.New("[NLP][POSTagging] input not provided")
		return nil, err
	}
	if len(Language) == 0 {
		Language = "en"
	}
	if len(GeneralizationSize) == 0 {
		GeneralizationSize = " "
	}
	if len(GeneralizationID) == 0 {
		GeneralizationID = " "
	}
	if len(sentenceSplit) == 0 {
		sentenceSplit = "0"
	}
	if len(Generalize) == 0 {
		Generalize = "0"
	}

	requestBody := strings.NewReader(fmt.Sprintf(`
		{
			"input": %s,
			"language":%s,
			"sentence_split":%s,
            "generalization_id": %s,
            "generalize": %s,
            "generalization_size": %s,

		}
	`, Input, Language, sentenceSplit, GeneralizationID, Generalize, GeneralizationSize))

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

	var response SentimentAnalysisResponse

	err = json.Unmarshal(b, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
