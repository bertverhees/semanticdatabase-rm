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

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func discardMockEFactory() {
	_ = assert.Equal
	_ = testing.Coverage
}

// TestMockEFactoryGetEPackage tests method GetEPackage
func TestMockEFactoryGetEPackage(t *testing.T) {
	o := &MockEFactory{}
	r := new(MockEPackage)
	o.On("GetEPackage").Once().Return(r)
	o.On("GetEPackage").Once().Return(func() EPackage {
		return r
	})
	assert.Equal(t, r, o.GetEPackage())
	assert.Equal(t, r, o.GetEPackage())
	o.AssertExpectations(t)
}

// TestMockEFactorySetEPackage tests method SetEPackage
func TestMockEFactorySetEPackage(t *testing.T) {
	o := &MockEFactory{}
	v := new(MockEPackage)
	o.On("SetEPackage", v).Once()
	o.SetEPackage(v)
	o.AssertExpectations(t)
}

// TestMockEFactoryConvertToString tests method ConvertToString
func TestMockEFactoryConvertToString(t *testing.T) {
	o := &MockEFactory{}
	eDataType := new(MockEDataType)
	instanceValue := interface{}(nil)
	r := string("Test String")
	o.On("ConvertToString", eDataType, instanceValue).Return(r).Once()
	o.On("ConvertToString", eDataType, instanceValue).Return(func() string {
		return r
	}).Once()
	assert.Equal(t, r, o.ConvertToString(eDataType, instanceValue))
	assert.Equal(t, r, o.ConvertToString(eDataType, instanceValue))
	o.AssertExpectations(t)
}

// TestMockEFactoryCreate tests method Create
func TestMockEFactoryCreate(t *testing.T) {
	o := &MockEFactory{}
	eClass := new(MockEClass)
	r := new(MockEObjectInternal)
	o.On("Create", eClass).Return(r).Once()
	o.On("Create", eClass).Return(func() EObject {
		return r
	}).Once()
	assert.Equal(t, r, o.Create(eClass))
	assert.Equal(t, r, o.Create(eClass))
	o.AssertExpectations(t)
}

// TestMockEFactoryCreateFromString tests method CreateFromString
func TestMockEFactoryCreateFromString(t *testing.T) {
	o := &MockEFactory{}
	eDataType := new(MockEDataType)
	literalValue := string("Test String")
	r := interface{}(nil)
	o.On("CreateFromString", eDataType, literalValue).Return(r).Once()
	o.On("CreateFromString", eDataType, literalValue).Return(func() interface{} {
		return r
	}).Once()
	assert.Equal(t, r, o.CreateFromString(eDataType, literalValue))
	assert.Equal(t, r, o.CreateFromString(eDataType, literalValue))
	o.AssertExpectations(t)
}
