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

// genInfo is a private struct which contains
// some important informations about a generalization table in
// Intellivoid's servers. It's used to manage, maintain, save,
// load and list generaliazation table for users.
// Users have to use interface `GenInfo` instead of this struct
// to get, create or list their generalization tables.
// It's supposed to be a cool feature in this library.
// Another feature for this struct and interface is ability to
// get or save the data in Json format (saving in a file).
// NOTICE: the comments in this file, are not supposed to be read
// by users. They are explanations about developers' ideas and
// why these private stuff exist in the first place. They are dev-only.
// So if you are a user and not a dev, it's better to just skip these comments.
type genInfo struct {
	// NewGen is true when this generalization info value
	// is created by function `CreateNew`.
	// Since this struct is private, users don't have to be
	// concerned about this. They however need to know if this
	// generalization value is a new generalization value,
	// it doesn't have any generalization ID (it's empty),
	// of course it's only before we send the request to Intellivoid's
	// servers. After that, it won't be considered as a
	// "new" generalization, so this value should become `false` after
	// we used it and saved it to the servers.
	NewGen bool `json:"is_new"`

	// GenId is the ID of this generalization.
	// It will be an empty string in the case we want to create a new
	// generalization table.
	// Which is why we have to check if this generalization table is
	// new or not. If it's not new, then users HAVE TO PROVIDE us
	// a generalization ID (if they want to update an
	// already existing generalization table that is), so if they
	// didn't pass us any ID, we will return them an error. This is my
	// ordinary idea of how to handle this horrible system of
	// Generalization Tables in Intellivoid-Coffeehous APIs.
	// I also told nektas about some of my ideas, but he refused them
	// all by some I-Cant-Undeastand-These-Stuff-So-They-Are-True words.
	//
	// ------------------------------------------
	//
	// here are some of them:
	// Q: Is there any way to get a list of generalization tables?	.
	// A: You generally want to keep that information on you.
	//
	// Q: Then if we don't save their ID on our side, we can't
	// update them anymore?
	// A: Yes.
	//
	// Q: Still I can't understand why we can't get a list of them.
	// All of them are saved on server side, aren't they?
	// They exist on Intellivoid's databases, then why we can't
	// get a list of them?
	// A: For performance issues, SpamProtectionBot uses this feature
	// heavily on a daily basis, so does many other bots alike
	// and or other software.
	// There are more than 8 million records created and updated on
	// a daily basis, retrieving these records can put a strain on the
	// database and it is more resource efficient to only have to retrieve
	// one record or more and cache that record in the database memory for
	// faster access, plus indexing helps too.
	// But you generally don’t want to ask the database for a large query,
	// even if you limit it to 100 records.
	// ...
	// It is your responsibility to hold this information,
	// same reason why Telegram doesn’t have the ability for bots
	// to retrieve a list of chats or users your bot talked to,
	// it doesn’t work like that and it can be dangerous if
	// your access key is leaked.
	// ...
	// If you’d like you can generally implement the same exact algorithm
	// on the client side but you will have to manage more than just an ID,
	// the reason the feature exists on the api in the first place
	// is for simplicity purposes where the server will create,
	// manage a table for you and keep track of the current pointer
	// and store all the related values. All you have to do is remember an ID.
	// ...
	// In general you would only keep the generalization ids for a temporary
	// amount of time if you are working with that data,
	// for instance chat rooms and whatnot.
	// Or tweets.
	GenId string `json:"generalization_id"`

	// Here is what docs says about it:
	// The size of the generalization table, cannot be 0 or less
	// or greater than what your subscription supports.
	// The mininum for a subscription is 5 and the maximum is 20.
	Size uint8 `json:"generalization_size"`
}
