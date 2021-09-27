// *****************************************************************************
// Copyright(c) 2021 MASA Group
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// *****************************************************************************

package ecore

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEListIterator(t *testing.T) {
	mockList := &MockEList{}
	mockList.On("Size").Return(3)
	for i := 0; i < 3; i++ {
		mockList.On("Get", i).Return(i)
	}
	it := &listIterator{list: mockList}
	for i := 0; it.HasNext(); i++ {
		assert.Equal(t, i, it.Next())
	}
	assert.Panics(t, func() { it.Next() })
}
