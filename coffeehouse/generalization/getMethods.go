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
	"os"
)

//---------------------------------------------------------

// IsNew will return `true` if this is a new generalization.
// but if this generalization already exists in the Intellivoid's
// servers, it will return false.
// Please take note that if this generalization is new,
// it doesn't have generalization ID (it's empty), because for
// new generalization, server will create a new ID automatically.
func (g *genInfo) IsNew() bool {
	return g.NewGen
}

// GetId returns the ID of this generalization. The ID is not
// an integer number, rather it's a string.
// If this generalization is new, the ID will be empty.
func (g *genInfo) GetId() string {
	return g.GenId
}

// GetSize returns the size of this generalization.
// The size of the generalization table cannot be 0 or
// less or greater than what your subscription supports.
// The mininum for a subscription is 5 and the maximum is 20.
func (g *genInfo) GetSize() uint8 {
	return g.Size
}

// ToJsonBytes method will convert this generalization value to
// json. You can use this method to serializa the data in your
// system memory, in files and/or database, so you can load it
// later and update your data in generalization tables.
func (g *genInfo) ToJsonBytes() ([]byte, error) {
	return json.Marshal(g)
}

// ToJsonString method will simply call ToJsonBytes method
// and will convert the value to a string.
// convert this generalization value to
// json. You can use this method to serializa the data in your
// system memory, in files and/or database, so you can load it
// later and update your data in generalization tables.
// This method will return an empty string and non-nil error
// if it fails.
func (g *genInfo) ToJsonString() (s string, err error) {
	var b []byte
	b, err = g.ToJsonBytes()
	if err != nil {
		return
	}

	s = string(b)
	return
}

// ToJsonPretty method will return you a pretty printed
// Json serialized value of this genInfo.
// You can use this method to serializa the data in your
// system memory, in files and/or database, so you can load it
// later and update your data in generalization tables.
// This method will return an empty string and non-nil error
// if it fails.
func (g *genInfo) ToJsonPretty() ([]byte, error) {
	return g.ToJsonIndent(JsonIndent)
}

// ToJsonPrettyStr method will return you a pretty printed
// Json serialized string value of this genInfo.
// You can use this method to serializa the data in your
// system memory, in files and/or database, so you can load it
// later and update your data in generalization tables.
// This method will return an empty string and non-nil error
// if it fails.
func (g *genInfo) ToJsonPrettyStr() (s string, err error) {
	var b []byte
	b, err = g.ToJsonPretty()
	if err != nil {
		return
	}

	s = string(b)
	return
}

// ToJsonIndent method will return you Json serialized value
// of this genInfo with the specified indentation.
// You can use this method to serializa the data in your
// system memory, in files and/or database, so you can load it
// later and update your data in generalization tables.
// This method will return an empty string and non-nil error
// if it fails.
func (g *genInfo) ToJsonIndent(indent string) (b []byte, err error) {
	return json.MarshalIndent(g, JsonPrefix, indent)
}

// ToJsonIndentStr method will return you Json serialized value
// of this genInfo with the specified indentation.
// You can use this method to serializa the data in your
// system memory, in files and/or database, so you can load it
// later and update your data in generalization tables.
// This method will return an empty string and non-nil error
// if it fails.
func (g *genInfo) ToJsonIndentStr(indent string) (s string, err error) {
	var b []byte
	b, err = g.ToJsonIndent(indent)
	if err != nil {
		return
	}

	s = string(b)
	return
}

// SaveToFile will pretty print this generalization info
// and will save it to the specified file on your local
// storage.
// it will return error if you pass an empty file name
// and or if there is any problem when saving the file.
func (g *genInfo) SaveToFile(filename string) error {
	if len(filename) == 0 {
		return errors.New("file name cannot be empty")
	}
	b, err := g.ToJsonPretty()
	if err != nil {
		return err
	}

	return os.WriteFile(filename, b, os.ModeType)
}

// SetID will set the id of this generalization info.
// You don't have to use it unless you want to copy the
// current generalization info into the new one and then
// update it in Intellivoid's servers.
// please notice that if you pass an empty id as the arguments,
// this method won't set the already existing id to empty,
// in fact, it won't do anything.
func (g *genInfo) SetID(id string) {
	if len(id) == 0 {
		return
	}

	g.NewGen = false
	g.GenId = id
}

//---------------------------------------------------------
