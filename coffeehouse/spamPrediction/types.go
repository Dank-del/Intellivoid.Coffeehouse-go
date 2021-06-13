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

package spamPrediction

type SpamPredictionAPIResponse struct {
	Success      bool     `json:"success"`
	ResponseCode int      `json:"response_code"`
	Results      *Results `json:"results"`
}

type Results struct {
	Text           string          `json:"text"`
	SourceLanguage string          `json:"source_language"`
	SpamPrediction *SpamPrediction `json:"spam_prediction"`
	Sentences      *Sentences      `json:"sentences"`
	Generalization *Generalization `json:"generalization"`
}

type Sentences struct {
	Text           string          `json:"text"`
	OffsetBegin    int             `json:"offset_begin"`
	OffsetEnd      int             `json:"offset_end"`
	SpamPrediction *SpamPrediction `json:"spam_prediction"`
}

type SpamPrediction struct {
	IsSpam      bool         `json:"is_spam"`
	Prediction  float64      `json:"prediction"`
	Predictions *Predictions `json:"predictions"`
}

type Predictions struct {
	Ham  float64 `json:"ham"`
	Spam float64 `json:"spam"`
}

type Generalization struct {
	ID             string         `json:"id"`
	Size           int            `json:"size"`
	TopLabel       string         `json:"top_label"`
	TopProbability float64        `json:"top_probability"`
	Probabilities  *Probabilities `json:"probabilities"`
}

type Probabilities struct {
	Label                 string    `json:"label"`
	CalculatedProbability float64   `json:"calculated_probability"`
	CurrentPointer        int       `json:"current_pointer"`
	Probabilities         []float64 `json:"probabilities"`
}
