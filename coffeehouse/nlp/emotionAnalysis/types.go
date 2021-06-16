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

package emotionAnalysis

import (
	cf "github.com/Dank-del/Intellivoid.Coffeehouse-go/coffeehouse"
)

// Emotionysis is a short name for `Emotion Analysis`.
// it contains the results for our analysing operation,
// and predicted emotions of a given text.
// this struct is much like sentimental analysis
// but more to do with emotions.
type Emotionysis struct {
	Success      bool      `json:"success"`
	ResponseCode int       `json:"response_code"`
	Results      *Results  `json:"results"`
	Error        *cf.Error `json:"error"`
}

type Results struct {
	Text           string             `json:"text"`
	SourceLanguage string             `json:"source_language"`
	Emotion        *Emotion           `json:"emotion"`
	Sentences      *Sentences         `json:"sentences"`
	Generalization *cf.Generalization `json:"generalization"`
}

type Emotion struct {
	Emotion     string       `json:"emotion"`
	Prediction  float64      `json:"prediction"`
	Predictions *Predictions `json:"predictions"`
}

type Predictions struct {
	Neutral   float64 `json:"NEUTRAL"`
	Happiness float64 `json:"HAPPINESS"`
	Affection float64 `json:"AFFECTION"`
	Sadness   float64 `json:"SADNESS"`
	Anger     float64 `json:"ANGER"`
}

type Sentences struct {
	Text        string   `json:"text"`
	OffsetBegin int      `json:"offset_begin"`
	OffsetEnd   int      `json:"offset_end"`
	Emotion     *Emotion `json:"emotion"`
}
