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

type GenInfo interface {
	// IsNew will check if the current generalization
	// information is related to a new generalization table
	// or not. if yes, the generalization table will be empty.
	IsNew() bool

	// GetId will return you the ID of this generalization table.
	// Please notice that if this generalization info is related
	// to a new generalization table, this method will return
	// an empty string.
	GetId() string

	// GetSize returns the size of this
	// generalization table. It's a non-zero uint8.
	// The size of the generalization table cannot be 0 or
	// less or greater than what your subscription supports.
	// The mininum for a subscription is 5 and the maximum is 20.
	GetSize() uint8

	// SetID will set the ID of this generalization info.
	// This method is mostly used for internal operation purpose.
	// You don't have to use it unless you want to copy the
	// current generalization info into the new one and then
	// update it in Intellivoid's servers.
	// Please notice that if you pass an empty id as the arguments,
	// this method won't set the already existing id to empty,
	// in fact, it won't do anything.
	SetID(id string)

	// ToJsonBytes method will convert this generalization info
	// value to Json serialized byte array. If you need a string
	// value instead of byte array, you can use `ToJsonString`
	// method, so it will give a string value straightly.
	// You can use this method to serializa the data in your
	// system memory, in files and/or database, so you can load it
	// later and update your data in generalization tables.
	ToJsonBytes() ([]byte, error)

	// ToJsonPretty method will return you a pretty printed
	// Json serialized value of this generalization table.
	// You can use this method to serializa the data in your
	// system memory, in files and/or database, so you can load it
	// later and update your data in generalization tables.
	// This method will return an empty string and non-nil error
	// if it fails.
	ToJsonPretty() ([]byte, error)

	// ToJsonString method will simply call ToJsonBytes method
	// and will convert the value to a string.
	// convert this generalization value to
	// json. You can use this method to serializa the data in your
	// system memory, in files and/or database, so you can load it
	// later and update your data in generalization tables.
	// This method will return an empty string and non-nil error
	// if it fails.
	ToJsonString() (string, error)

	// ToJsonPrettyStr method will return you a pretty printed
	// Json serialized string value of this genInfo.
	// You can use this method to serializa the data in your
	// system memory, in files and/or database, so you can load it
	// later and update your data in generalization tables.
	// This method will return an empty string and non-nil error
	// if it fails.
	ToJsonPrettyStr() (string, error)

	// SaveToFile will pretty print this generalization info
	// and will save it to the specified file on your local
	// storage.
	// it will return error if you pass an empty file name
	// and or if there is any problem when saving the file.
	SaveToFile(filename string) error

	// ToJsonIndent method will return you Json serialized value
	// of this generalization info with the specified indentation.
	// we mostly recommend you to use double space or more.
	// But you can also use tab character ('\t' or '\u0x09').
	// You can use this method to serializa the data in your
	// system memory, in files and/or database, so you can load it
	// later and update your data in generalization tables.
	// This method will return an empty string and non-nil error
	// if it fails.
	// Please notice that if you enter a non-whitespace
	// character for this method, it won't return you any
	// kind of errors.
	ToJsonIndent(indent string) ([]byte, error)

	// ToJsonIndent method will return you **string** value of
	// Json serialized value of this generalization info
	// with the specified indentation.
	// we mostly recommend you to use double space or more.
	// But you can also use tab character ('\t' or '\u0x09').
	// You can use this method to serializa the data in your
	// system memory, in files and/or database, so you can load it
	// later and update your data in generalization tables.
	// This method will return an empty string and non-nil error
	// if it fails.
	// Please notice that if you enter a non-whitespace
	// character for this method, it won't return you any
	// kind of errors.
	ToJsonIndentStr(indent string) (string, error)
}
