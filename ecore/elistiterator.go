// *****************************************************************************
// Copyright(c) 2021 MASA Group
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// *****************************************************************************

package ecore

type listIterator struct {
	cursor int
	list   EList
}

// Next return the current value of the iterator
func (it *listIterator) Next() interface{} {
	i := it.cursor
	if i >= it.list.Size() {
		panic("Not such an element")
	}
	it.cursor = i + 1
	return it.list.Get(i)
}

// HasNext make the iterator go further in the array
func (it *listIterator) HasNext() bool {
	return it.cursor < it.list.Size()
}
