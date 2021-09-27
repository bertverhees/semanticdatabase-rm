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

func discardMockEReference() {
	_ = assert.Equal
	_ = testing.Coverage
}

// TestMockEReferenceIsContainer tests method IsContainer
func TestMockEReferenceIsContainer(t *testing.T) {
	o := &MockEReference{}
	r := bool(true)
	o.On("IsContainer").Once().Return(r)
	o.On("IsContainer").Once().Return(func() bool {
		return r
	})
	assert.Equal(t, r, o.IsContainer())
	assert.Equal(t, r, o.IsContainer())
	o.AssertExpectations(t)
}

// TestMockEReferenceIsContainment tests method IsContainment
func TestMockEReferenceIsContainment(t *testing.T) {
	o := &MockEReference{}
	r := bool(true)
	o.On("IsContainment").Once().Return(r)
	o.On("IsContainment").Once().Return(func() bool {
		return r
	})
	assert.Equal(t, r, o.IsContainment())
	assert.Equal(t, r, o.IsContainment())
	o.AssertExpectations(t)
}

// TestMockEReferenceSetContainment tests method SetContainment
func TestMockEReferenceSetContainment(t *testing.T) {
	o := &MockEReference{}
	v := bool(true)
	o.On("SetContainment", v).Once()
	o.SetContainment(v)
	o.AssertExpectations(t)
}

// TestMockEReferenceGetEKeys tests method GetEKeys
func TestMockEReferenceGetEKeys(t *testing.T) {
	o := &MockEReference{}
	l := &MockEList{}
	// return a value
	o.On("GetEKeys").Once().Return(l)
	o.On("GetEKeys").Once().Return(func() EList {
		return l
	})
	assert.Equal(t, l, o.GetEKeys())
	assert.Equal(t, l, o.GetEKeys())
	o.AssertExpectations(t)
}

// TestMockEReferenceGetEOpposite tests method GetEOpposite
func TestMockEReferenceGetEOpposite(t *testing.T) {
	o := &MockEReference{}
	r := new(MockEReference)
	o.On("GetEOpposite").Once().Return(r)
	o.On("GetEOpposite").Once().Return(func() EReference {
		return r
	})
	assert.Equal(t, r, o.GetEOpposite())
	assert.Equal(t, r, o.GetEOpposite())
	o.AssertExpectations(t)
}

// TestMockEReferenceSetEOpposite tests method SetEOpposite
func TestMockEReferenceSetEOpposite(t *testing.T) {
	o := &MockEReference{}
	v := new(MockEReference)
	o.On("SetEOpposite", v).Once()
	o.SetEOpposite(v)
	o.AssertExpectations(t)
}

// TestMockEReferenceGetEReferenceType tests method GetEReferenceType
func TestMockEReferenceGetEReferenceType(t *testing.T) {
	o := &MockEReference{}
	r := new(MockEClass)
	o.On("GetEReferenceType").Once().Return(r)
	o.On("GetEReferenceType").Once().Return(func() EClass {
		return r
	})
	assert.Equal(t, r, o.GetEReferenceType())
	assert.Equal(t, r, o.GetEReferenceType())
	o.AssertExpectations(t)
}

// TestMockEReferenceIsResolveProxies tests method IsResolveProxies
func TestMockEReferenceIsResolveProxies(t *testing.T) {
	o := &MockEReference{}
	r := bool(true)
	o.On("IsResolveProxies").Once().Return(r)
	o.On("IsResolveProxies").Once().Return(func() bool {
		return r
	})
	assert.Equal(t, r, o.IsResolveProxies())
	assert.Equal(t, r, o.IsResolveProxies())
	o.AssertExpectations(t)
}

// TestMockEReferenceSetResolveProxies tests method SetResolveProxies
func TestMockEReferenceSetResolveProxies(t *testing.T) {
	o := &MockEReference{}
	v := bool(true)
	o.On("SetResolveProxies", v).Once()
	o.SetResolveProxies(v)
	o.AssertExpectations(t)
}
