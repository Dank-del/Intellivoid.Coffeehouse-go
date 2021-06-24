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
	"github.com/Dank-del/Intellivoid.Coffeehouse-go/coffeehouse/nlp/posTagging"
)

func TestPOSTagging(t *testing.T) {
	coffeehouse.SetKey("<access_key>")
	res, err := posTagging.TagPOSFull("hello, how are you?", "en", "0")
	if err != nil {
		log.Fatal(err.Error())
		t.Log(err)
	}

	if res.Success != true {
		t.Log(res)
		t.Errorf("[Intellivoid.Coffeehouse-go (posTagging)] Failed request, response code: %d", res.ResponseCode)

	}
}
