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

type SentimentAnalysisResponse struct {
	Success      bool `json:"success"`
	ResponseCode int  `json:"response_code"`
	Results  *Results `json:"results"`
}

type Results struct {
	Text           string `json:"text"`
	SourceLanguage string `json:"source_language"`
	Sentiment *Sentiment`json:"sentiment"`
	Sentences *Sentences `json:"sentences"`
	Generalization interface{} `json:"generalization"`
}

type Sentences struct {
	Text        string `json:"text"`
	OffsetBegin int    `json:"offset_begin"`
	OffsetEnd   int    `json:"offset_end"`
	Sentiment *Sentiment`json:"sentiment"`
}

type Sentiment struct {
	Sentiment   string  `json:"sentiment"`
	Prediction  float64 `json:"prediction"`
	Predictions *Predictions `json:"predictions"`
}

type Predictions struct {
	VeryNegative float64 `json:"VERY_NEGATIVE"`
	Negative     float64 `json:"NEGATIVE"`
	Neutral      float64 `json:"NEUTRAL"`
	Positive     float64 `json:"POSITIVE"`
	VeryPositive float64 `json:"VERY_POSITIVE"`
}