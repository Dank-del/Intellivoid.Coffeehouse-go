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
	"testing"

	"github.com/Dank-del/Intellivoid.Coffeehouse-go/coffeehouse"
	"github.com/Dank-del/Intellivoid.Coffeehouse-go/coffeehouse/mediaProcessing/classificationNSFW"
)

func TestNSFWClassification(t *testing.T) {
	coffeehouse.SetKey("9af1dc94e606dc199e4d38ad2f80193742aee3fabc4c714b4bacf9ccc690ebfde6af8def68427540dbc4f0e3e57cc90ba21b4f27e81d367b551ae8acbaf99471")
	res, err := classificationNSFW.DoRequest("nsfw/owo.jpg")
	if err != nil {
		t.Errorf(err.Error())
	}
	if res.Success != true {
		t.Errorf("[Intellivoid.Coffeehouse-go (classificationNSFW)] Failed request, response code: %d", res.ResponseCode)
	}
}
