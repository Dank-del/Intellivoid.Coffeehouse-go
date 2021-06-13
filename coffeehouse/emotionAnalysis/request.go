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


package emotionAnalysis

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Dank-del/Intellivoid.Coffeehouse-go/coffeehouse"
	"io/ioutil"
	"net/http"
)

func DoRequest(Input string, Language string, sentenceSplit string, Generalize string, GeneralizationSize string, GeneralizationID string) (data **EmotionAnalysisAPIResponse, err error){
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
	if len(GeneralizationSize) == 0 {
		GeneralizationSize = " "
	}
	if len(GeneralizationID) == 0 {
		GeneralizationID = " "
	}
	if len(Generalize) == 0 {
		Generalize = "0"
	}
	url := fmt.Sprintf(endpointurl, coffeehouse.CofeeHouseAPIKey, Input, Language, sentenceSplit, Generalize, GeneralizationID, GeneralizationSize)
	res, err := http.Post(url, "text/plain", nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var Response *EmotionAnalysisAPIResponse

	err = json.Unmarshal(b, &Response)
	if err != nil {
		return nil, err
	}
	return &Response, nil
}
