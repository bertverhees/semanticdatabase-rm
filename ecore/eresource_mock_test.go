// *****************************************************************************
//
// This file is part of a MASA library or program.
// Refer to the included end-user license agreement for restrictions.
//
// Copyright (c) 2020 MASA Group
//
// *****************************************************************************

package ecore

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestMockEResourceGetResourceSet tests method GetResourceSet
func TestMockEResourceGetResourceSet(t *testing.T) {
	r := &MockEResource{}
	set := &MockEResourceSet{}
	r.On("GetResourceSet").Return(set).Once()
	r.On("GetResourceSet").Return(func() EResourceSet {
		return set
	}).Once()
	assert.Equal(t, set, r.GetResourceSet())
	assert.Equal(t, set, r.GetResourceSet())
	r.AssertExpectations(t)
}

// TestMockEResourceGetURI tests method GetURI
func TestMockEResourceGetURI(t *testing.T) {
	r := &MockEResource{}
	uri := &URI{}
	r.On("GetURI").Return(uri).Once()
	r.On("GetURI").Return(func() *URI {
		return uri
	}).Once()
	assert.Equal(t, uri, r.GetURI())
	assert.Equal(t, uri, r.GetURI())
	r.AssertExpectations(t)
}

// TestMockEResourceSetURI tests method SetURI
func TestMockEResourceSetURI(t *testing.T) {
	r := &MockEResource{}
	uri := &URI{}
	r.On("SetURI", uri).Once()
	r.SetURI(uri)
	r.AssertExpectations(t)
}

// TestMockEResourceGetContents tests method GetContents
func TestMockEResourceGetContents(t *testing.T) {
	r := &MockEResource{}
	c := &MockEList{}
	r.On("GetContents").Return(c).Once()
	r.On("GetContents").Return(func() EList {
		return c
	}).Once()
	assert.Equal(t, c, r.GetContents())
	assert.Equal(t, c, r.GetContents())
	r.AssertExpectations(t)
}

// TestMockEResourceGetAllContents tests method GetAllContents
func TestMockEResourceGetAllContents(t *testing.T) {
	r := &MockEResource{}
	i := &MockEIterator{}
	r.On("GetAllContents").Return(i).Once()
	r.On("GetAllContents").Return(func() EIterator {
		return i
	}).Once()
	assert.Equal(t, i, r.GetAllContents())
	assert.Equal(t, i, r.GetAllContents())
	r.AssertExpectations(t)
}

// TestMockEResourceAttached tests method Attached
func TestMockEResourceAttached(t *testing.T) {
	r := &MockEResource{}
	o := &MockEObject{}
	r.On("Attached", o)
	r.Attached(o)
	r.AssertExpectations(t)
}

// TestMockEResourceAttached tests method Detached
func TestMockEResourceDetached(t *testing.T) {
	r := &MockEResource{}
	o := &MockEObject{}
	r.On("Detached", o)
	r.Detached(o)
	r.AssertExpectations(t)
}

// TestMockEResourceGetEObject tests method GetEObject
func TestMockEResourceGetEObject(t *testing.T) {
	r := &MockEResource{}
	o := &MockEObject{}
	r.On("GetEObject", "test").Return(o).Once()
	r.On("GetEObject", "test").Return(func(string) EObject {
		return o
	}).Once()
	assert.Equal(t, o, r.GetEObject("test"))
	assert.Equal(t, o, r.GetEObject("test"))
	r.AssertExpectations(t)
}

// TestMockEResourceGetURIFragment tests method GetURIFragment
func TestMockEResourceGetURIFragment(t *testing.T) {
	r := &MockEResource{}
	o := &MockEObject{}
	r.On("GetURIFragment", o).Return("fragment").Once()
	r.On("GetURIFragment", o).Return(func(EObject) string {
		return "fragment"
	}).Once()
	assert.Equal(t, "fragment", r.GetURIFragment(o))
	assert.Equal(t, "fragment", r.GetURIFragment(o))
	r.AssertExpectations(t)
}

// TestMockEResourceLoad tests method Load
func TestMockEResourceLoad(t *testing.T) {
	r := &MockEResource{}
	r.On("Load").Once()
	r.Load()
	r.AssertExpectations(t)
}

func TestMockEResourceLoadWithOptions(t *testing.T) {
	r := &MockEResource{}
	m := make(map[string]interface{})
	r.On("LoadWithOptions", m).Once()
	r.LoadWithOptions(m)
	r.AssertExpectations(t)
}

// TestMockEResourceLoadWithReader tests method LoadWithReader
func TestMockEResourceLoadWithReader(t *testing.T) {
	r := &MockEResource{}
	m := make(map[string]interface{})
	r.On("LoadWithReader", nil, m).Once()
	r.LoadWithReader(nil, m)
	r.AssertExpectations(t)
}

// TestMockEResourceUnLoad tests method UnLoad
func TestMockEResourceUnLoad(t *testing.T) {
	r := &MockEResource{}
	r.On("Unload").Once()
	r.Unload()
	r.AssertExpectations(t)
}

// TestMockEResourceSave tests method Save
func TestMockEResourceSave(t *testing.T) {
	r := &MockEResource{}
	r.On("Save").Once()
	r.Save()
	r.AssertExpectations(t)
}

func TestMockEResourceSaveWithOptions(t *testing.T) {
	r := &MockEResource{}
	m := make(map[string]interface{})
	r.On("SaveWithOptions", m).Once()
	r.SaveWithOptions(m)
	r.AssertExpectations(t)
}

// TestMockEResourceSaveWithReader tests method SaveWithReader
func TestMockEResourceSaveWithReader(t *testing.T) {
	r := &MockEResource{}
	m := make(map[string]interface{})
	r.On("SaveWithWriter", nil, m).Once()
	r.SaveWithWriter(nil, m)
	r.AssertExpectations(t)
}

// TestMockEResourceGetErrors tests method GetErrors
func TestMockEResourceGetErrors(t *testing.T) {
	r := &MockEResource{}
	c := &MockEList{}
	r.On("GetErrors").Return(c).Once()
	r.On("GetErrors").Return(func() EList {
		return c
	}).Once()
	assert.Equal(t, c, r.GetErrors())
	assert.Equal(t, c, r.GetErrors())
	r.AssertExpectations(t)
}

// TestMockEResourceGetWarnings tests method GetWarnings
func TestMockEResourceGetWarnings(t *testing.T) {
	r := &MockEResource{}
	c := &MockEList{}
	r.On("GetWarnings").Return(c).Once()
	r.On("GetWarnings").Return(func() EList {
		return c
	}).Once()
	assert.Equal(t, c, r.GetWarnings())
	assert.Equal(t, c, r.GetWarnings())
	r.AssertExpectations(t)
}

// TestMockEResourceIsLoaded tests method IsLoaded
func TestMockEResourceIsLoaded(t *testing.T) {
	r := &MockEResource{}
	r.On("IsLoaded").Return(true).Once()
	r.On("IsLoaded").Return(func() bool {
		return false
	}).Once()
	assert.True(t, r.IsLoaded())
	assert.False(t, r.IsLoaded())
	r.AssertExpectations(t)
}

func TestMockEResourceGetIDManager(t *testing.T) {
	r := &MockEResource{}
	m := &MockEObjectIDManager{}
	r.On("GetObjectIDManager").Return(m).Once()
	r.On("GetObjectIDManager").Return(func() EObjectIDManager {
		return m
	}).Once()
	assert.Equal(t, m, r.GetObjectIDManager())
	assert.Equal(t, m, r.GetObjectIDManager())
	r.AssertExpectations(t)
}

func TestMockEResourceSetIDManager(t *testing.T) {
	r := &MockEResource{}
	m := &MockEObjectIDManager{}
	r.On("SetObjectIDManager", m).Once()
	r.SetObjectIDManager(m)
	r.AssertExpectations(t)

}
