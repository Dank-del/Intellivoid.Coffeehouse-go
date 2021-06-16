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

package tests

import (
	"log"
	"testing"

	"github.com/Dank-del/Intellivoid.Coffeehouse-go/coffeehouse"
	"github.com/Dank-del/Intellivoid.Coffeehouse-go/coffeehouse/nlp/sentimentAnalysis"
)

func TestSentimentAnal(t *testing.T) {
	coffeehouse.SetKey("3056b00704d72611e19e5c6df580798864155d9818175ffa1a4ae4a1c1496eeca48ad7aa9a625918431edbc1fa3ea712b2d591dd842e0bdb237c14a545cdd068")
	// res, err := sentimentAnalysis.DoRequest("Hey there, I'm Sayan. How are you doing?", "en", "", "", "", "")
	// res, err := sentimentAnalysis.DoRequest("After I've been accepted, I wanted to cry out from happiness", "en", "", "", "", "")
	res, err := sentimentAnalysis.DoRequest("I'm very happy", "en", "", "", "", "")
	if err != nil {
		log.Fatal(err.Error())
		t.Log(err)
	}

	log.Println(res)

	if res.Success != true {
		t.Log(res)
		t.Errorf("[Intellivoid.Coffeehouse-go (sentimentAnalysis)] Failed request, response code: %d", res.ResponseCode)

	}
}
