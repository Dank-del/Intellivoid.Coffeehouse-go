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

type POSApiResponse struct {
	Success      bool `json:"success"`
	ResponseCode int  `json:"response_code"`
	Results *Results  `json:"results"`
	Error *Error `json:"error"`
}

type Error struct {
    ErrorCode int    `json:"error_code"`
    Type      string `json:"type"`
    Message   string `json:"message"`
}

type Results struct {
	Text           string `json:"text"`
	SourceLanguage string `json:"source_language"`
	Tags           *Tags `json:"tags"`
    Sentences      *Sentences `json:"sentences"`
}

type Sentences struct {
	Text        string `json:"text"`
	OffsetBegin int    `json:"offset_begin"`
	OffsetEnd   int    `json:"offset_end"`
	Tags        *Tags  `json:"tags"`
}

type Tags struct {
	Text        string `json:"text"`
	OffsetBegin int    `json:"offset_begin"`
	OffsetEnd   int    `json:"offset_end"`
	TagValue    string `json:"tag_value"`
}


