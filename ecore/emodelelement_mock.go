// Code generated by soft.generator.go. DO NOT EDIT.

// *****************************************************************************
// Copyright(c) 2021 MASA Group
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// *****************************************************************************

package ecore

type MockEModelElement struct {
	MockEObjectInternal
}

// GetEAnnotations get the value of eAnnotations
func (eModelElement *MockEModelElement) GetEAnnotations() EList {
	ret := eModelElement.Called()

	var r EList
	if rf, ok := ret.Get(0).(func() EList); ok {
		r = rf()
	} else {
		if ret.Get(0) != nil {
			r = ret.Get(0).(EList)
		}
	}

	return r
}

// GetEAnnotation provides mock implementation
func (eModelElement *MockEModelElement) GetEAnnotation(source string) EAnnotation {
	ret := eModelElement.Called(source)

	var r EAnnotation
	if rf, ok := ret.Get(0).(func() EAnnotation); ok {
		r = rf()
	} else {
		if ret.Get(0) != nil {
			r = ret.Get(0).(EAnnotation)
		}
	}

	return r
}
