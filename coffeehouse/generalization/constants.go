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

package generalization

// public constants values used in generalization packages.
const (
	// MinSize is the minimum possible size of a generalization table
	// that can be passed to the generalization info.
	MinSize = 1

	// MaxSize is the maximum possible size of a generalization table
	// that can be passed to the generalization info.
	MaxSize = 20

	// JsonIndent is the indetation value that used in the Json.Marshal
	// method for pretty printing the json value. It's two space.
	// users can set the value as a custom value, instead of using this.
	JsonIndent = "  "

	// JsonPrefix is the prefix value that used in the Json.Marshal
	// method for pretty printing the json value. It's an empty string.
	// users can set the value as a custom value, instead of using this.
	JsonPrefix = ""
)
