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

package tests

import (
	"github.com/Dank-del/Intellivoid.Coffeehouse-go/coffeehouse"
	"github.com/Dank-del/Intellivoid.Coffeehouse-go/coffeehouse/emotionAnalysis"
	"log"
	"testing"
)

func TestEmotionAnal(t *testing.T) {
	coffeehouse.SetKey("<API Key here>")
	res, err := emotionAnalysis.DoRequest("Hey there, I'm Sayan. How are you doing?", "en", "", "", "", "")
	if err != nil {
		log.Fatal(err.Error())
	}
	if (*res).Success != true {
		t.Errorf("[Intellivoid.Coffeehouse-go (emotionAnalysis)] Failed request, response code: %d", (*res).ResponseCode)

	}
}
