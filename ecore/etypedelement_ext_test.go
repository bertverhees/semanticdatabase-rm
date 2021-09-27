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

func TestTypedElementIsMany(t *testing.T) {
	e := newETypedElementExt()
	e.SetUpperBound(1)
	assert.False(t, e.IsMany())
	e.SetUpperBound(2)
	assert.True(t, e.IsMany())
	e.SetUpperBound(UNBOUNDED_MULTIPLICITY)
	assert.True(t, e.IsMany())
}

func TestTypedElementIsRequired(t *testing.T) {
	e := newETypedElementExt()
	e.SetLowerBound(0)
	assert.False(t, e.IsRequired())
	e.SetLowerBound(1)
	assert.True(t, e.IsRequired())
}
