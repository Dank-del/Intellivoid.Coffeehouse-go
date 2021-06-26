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

package classificationNSFW

// endpoint url used for sending HTTP request for this
// API.
const (
	endpointurl = "https://api.intellivoid.net/" +
		"coffeehouse/" +
		"v1/image/" +
		"nsfw_classification?"
)

// keys used in endpoint key-value map
const (
	accessKeyKey          = "access_key"
	imageKey              = "image"
	generalizationKey     = "generalize"
	generalizationSizeKey = "generalization_size"
	generalizationIdKey   = "generalization_id"
)
