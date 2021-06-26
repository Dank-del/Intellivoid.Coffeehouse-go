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

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

// CreateNew will create a new generalization info with the
// specified size.
// The returned value doesn't have any ID.
func CreateNew(size uint8) (info GenInfo, err error) {
	if size < MinSize {
		return nil, errors.New("size cannot be less than 1")
	}

	if size > MaxSize {
		return nil, errors.New("size cannot be grater than 20")
	}

	info = &genInfo{
		NewGen: true,
		Size:   size,
		GenId:  "",
	}

	return
}

// LoadFromBytes will load a generalization info value from the
// given file name.
// It will return error if the file doesn't exist or it doesn't
// contain correct values.
func LoadFromFile(filename string) (info GenInfo, err error) {
	if len(filename) == 0 {
		return nil, errors.New("filename cannot be empty")
	}
	var b []byte
	b, err = ioutil.ReadFile(filename)
	if err != nil {
		return
	}

	return LoadFromBytes(b)
}

// LoadFromBytes will load a generalization info value from the
// given byte array.
// it will return error if the given value is not correct.
func LoadFromBytes(b []byte) (info GenInfo, err error) {
	if b == nil {
		return nil, errors.New("byes cannot be nil")
	}

	var gInfo genInfo
	err = json.Unmarshal(b, &gInfo)
	if err != nil {
		return
	}

	info = &gInfo
	return
}
