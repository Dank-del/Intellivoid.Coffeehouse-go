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

const (
	endpointurl = "https://api.intellivoid.net/" +
		"coffeehouse/" +
		"v1/nlp/" +
		"emotion_analysis?" +
		accessKeyKey + sAnd +
		inputKey + sAnd +
		languageKey + sAnd +
		sentenceSplitKey + sAnd +
		generalizationKey + sAnd +
		generalizationSizeKey + sAnd +
		generalizationIdKey + sFormat
)

// keys used in endpoint
const (
	accessKeyKey          = "access_key"
	inputKey              = "input"
	languageKey           = "language"
	sentenceSplitKey      = "sentence_split"
	generalizationKey     = "generalize"
	generalizationSizeKey = "generalization_size"
	generalizationIdKey   = "generalization_id"
	sFormat               = "=%s"
	sAnd                  = sFormat + "&"
)
