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
	"github.com/stretchr/testify/mock"
)

func TestMockEStoreGet(t *testing.T) {
	o := &MockEStore{}
	mockObject := &MockEObject{}
	mockFeature := &MockEStructuralFeature{}
	mockResult := &MockEObject{}
	o.On("Get", mockObject, mockFeature, 0).Return(mockResult).Once()
	o.On("Get", mockObject, mockFeature, 0).Return(func(EObject, EStructuralFeature, int) interface{} {
		return mockResult
	}).Once()
	assert.Equal(t, mockResult, o.Get(mockObject, mockFeature, 0))
	assert.Equal(t, mockResult, o.Get(mockObject, mockFeature, 0))
	o.AssertExpectations(t)
}

func TestMockEStoreSet(t *testing.T) {
	o := &MockEStore{}
	mockObject := &MockEObject{}
	mockFeature := &MockEStructuralFeature{}
	mockValue := &MockEObject{}
	mockOld := &MockEObject{}
	o.On("Set", mockObject, mockFeature, 0, mockValue).Return(mockOld).Once()
	o.On("Set", mockObject, mockFeature, 0, mockValue).Return(func(object EObject, feature EStructuralFeature, index int, value interface{}) interface{} {
		return mockOld
	}).Once()
	assert.Equal(t, mockOld, o.Set(mockObject, mockFeature, 0, mockValue))
	assert.Equal(t, mockOld, o.Set(mockObject, mockFeature, 0, mockValue))
	o.AssertExpectations(t)
}

func TestMockEStoreIsSet(t *testing.T) {
	o := &MockEStore{}
	mockObject := &MockEObject{}
	mockFeature := &MockEStructuralFeature{}
	o.On("IsSet", mockObject, mockFeature).Return(false).Once()
	o.On("IsSet", mockObject, mockFeature).Return(func(EObject, EStructuralFeature) bool {
		return true
	}).Once()
	assert.False(t, o.IsSet(mockObject, mockFeature))
	assert.True(t, o.IsSet(mockObject, mockFeature))
	o.AssertExpectations(t)
}

func TestMockEStoreUnSet(t *testing.T) {
	o := &MockEStore{}
	mockObject := &MockEObject{}
	mockFeature := &MockEStructuralFeature{}
	o.On("UnSet", mockObject, mockFeature).Once()
	o.UnSet(mockObject, mockFeature)
	o.AssertExpectations(t)
}

func TestMockEStoreIsEmpty(t *testing.T) {
	o := &MockEStore{}
	mockObject := &MockEObject{}
	mockFeature := &MockEStructuralFeature{}
	o.On("IsEmpty", mockObject, mockFeature).Return(false).Once()
	o.On("IsEmpty", mockObject, mockFeature).Return(func(EObject, EStructuralFeature) bool {
		return true
	}).Once()
	assert.False(t, o.IsEmpty(mockObject, mockFeature))
	assert.True(t, o.IsEmpty(mockObject, mockFeature))
	o.AssertExpectations(t)
}

func TestMockEStoreContains(t *testing.T) {
	o := &MockEStore{}
	mockObject := &MockEObject{}
	mockFeature := &MockEStructuralFeature{}
	mockValue := &MockEObject{}
	o.On("Contains", mockObject, mockFeature, mockValue).Return(false).Once()
	o.On("Contains", mockObject, mockFeature, mockValue).Return(func(EObject, EStructuralFeature, interface{}) bool {
		return true
	}).Once()
	assert.False(t, o.Contains(mockObject, mockFeature, mockValue))
	assert.True(t, o.Contains(mockObject, mockFeature, mockValue))
	o.AssertExpectations(t)
}

func TestMockEStoreSize(t *testing.T) {
	o := &MockEStore{}
	mockObject := &MockEObject{}
	mockFeature := &MockEStructuralFeature{}
	o.On("Size", mockObject, mockFeature).Return(1).Once()
	o.On("Size", mockObject, mockFeature).Return(func(EObject, EStructuralFeature) int {
		return 2
	}).Once()
	assert.Equal(t, 1, o.Size(mockObject, mockFeature))
	assert.Equal(t, 2, o.Size(mockObject, mockFeature))
	o.AssertExpectations(t)
}

func TestMockEStoreIndexOf(t *testing.T) {
	o := &MockEStore{}
	mockObject := &MockEObject{}
	mockFeature := &MockEStructuralFeature{}
	mockValue := &MockEObject{}
	o.On("IndexOf", mockObject, mockFeature, mockValue).Return(1).Once()
	o.On("IndexOf", mockObject, mockFeature, mockValue).Return(func(EObject, EStructuralFeature, interface{}) int {
		return 2
	}).Once()
	assert.Equal(t, 1, o.IndexOf(mockObject, mockFeature, mockValue))
	assert.Equal(t, 2, o.IndexOf(mockObject, mockFeature, mockValue))
	o.AssertExpectations(t)
}

func TestMockEStoreLastIndexOf(t *testing.T) {
	o := &MockEStore{}
	mockObject := &MockEObject{}
	mockFeature := &MockEStructuralFeature{}
	mockValue := &MockEObject{}
	o.On("LastIndexOf", mockObject, mockFeature, mockValue).Return(1).Once()
	o.On("LastIndexOf", mockObject, mockFeature, mockValue).Return(func(EObject, EStructuralFeature, interface{}) int {
		return 2
	}).Once()
	assert.Equal(t, 1, o.LastIndexOf(mockObject, mockFeature, mockValue))
	assert.Equal(t, 2, o.LastIndexOf(mockObject, mockFeature, mockValue))
	o.AssertExpectations(t)
}

func TestMockEStoreAdd(t *testing.T) {
	o := &MockEStore{}
	mockObject := &MockEObject{}
	mockFeature := &MockEStructuralFeature{}
	mockValue := &MockEObject{}
	o.On("Add", mockObject, mockFeature, 0, mockValue).Once()
	o.Add(mockObject, mockFeature, 0, mockValue)
	o.AssertExpectations(t)
}

func TestMockEStoreRemove(t *testing.T) {
	o := &MockEStore{}
	mockObject := &MockEObject{}
	mockFeature := &MockEStructuralFeature{}
	mockOld := &MockEObject{}
	o.On("Remove", mockObject, mockFeature, 0).Return(mockOld).Once()
	o.On("Remove", mockObject, mockFeature, 0).Return(func(object EObject, feature EStructuralFeature, index int) interface{} {
		return mockOld
	}).Once()
	assert.Equal(t, mockOld, o.Remove(mockObject, mockFeature, 0))
	assert.Equal(t, mockOld, o.Remove(mockObject, mockFeature, 0))
	mock.AssertExpectationsForObjects(t, mockObject, mockFeature, mockOld, o)
}

func TestMockEStoreMove(t *testing.T) {
	o := &MockEStore{}
	mockObject := &MockEObject{}
	mockFeature := &MockEStructuralFeature{}
	mockOld := &MockEObject{}
	o.On("Move", mockObject, mockFeature, 0, 1).Return(mockOld).Once()
	o.On("Move", mockObject, mockFeature, 0, 1).Return(func(object EObject, feature EStructuralFeature, index int, old int) interface{} {
		return mockOld
	}).Once()
	assert.Equal(t, mockOld, o.Move(mockObject, mockFeature, 0, 1))
	assert.Equal(t, mockOld, o.Move(mockObject, mockFeature, 0, 1))
	mock.AssertExpectationsForObjects(t, mockObject, mockFeature, mockOld, o)
}

func TestMockEStoreClear(t *testing.T) {
	o := &MockEStore{}
	mockObject := &MockEObject{}
	mockFeature := &MockEStructuralFeature{}
	o.On("Clear", mockObject, mockFeature).Once()
	o.Clear(mockObject, mockFeature)
	o.AssertExpectations(t)
}

func TestMockEStoreToArray(t *testing.T) {
	o := &MockEStore{}
	mockObject := &MockEObject{}
	mockFeature := &MockEStructuralFeature{}
	mockResult := []interface{}{}
	o.On("ToArray", mockObject, mockFeature).Return(mockResult).Once()
	o.On("ToArray", mockObject, mockFeature).Return(func(EObject, EStructuralFeature) []interface{} {
		return mockResult
	}).Once()
	assert.Equal(t, mockResult, o.ToArray(mockObject, mockFeature))
	assert.Equal(t, mockResult, o.ToArray(mockObject, mockFeature))
	o.AssertExpectations(t)
}

func TestMockEStoreGetContainer(t *testing.T) {
	o := &MockEStore{}
	mockObject := &MockEObject{}
	mockResult := &MockEObject{}
	o.On("GetContainer", mockObject).Return(mockResult).Once()
	o.On("GetContainer", mockObject).Return(func(EObject) EObject {
		return mockResult
	}).Once()
	assert.Equal(t, mockResult, o.GetContainer(mockObject))
	assert.Equal(t, mockResult, o.GetContainer(mockObject))
	o.AssertExpectations(t)
}

func TestMockEStoreGetContainingFeature(t *testing.T) {
	o := &MockEStore{}
	mockObject := &MockEObject{}
	mockResult := &MockEStructuralFeature{}
	o.On("GetContainingFeature", mockObject).Return(mockResult).Once()
	o.On("GetContainingFeature", mockObject).Return(func(EObject) EStructuralFeature {
		return mockResult
	}).Once()
	assert.Equal(t, mockResult, o.GetContainingFeature(mockObject))
	assert.Equal(t, mockResult, o.GetContainingFeature(mockObject))
	o.AssertExpectations(t)
}

func TestMockEStoreCreate(t *testing.T) {
	o := &MockEStore{}
	mockClass := &MockEClass{}
	mockResult := &MockEObject{}
	o.On("Create", mockClass).Return(mockResult).Once()
	o.On("Create", mockClass).Return(func(EClass) EObject {
		return mockResult
	}).Once()
	assert.Equal(t, mockResult, o.Create(mockClass))
	assert.Equal(t, mockResult, o.Create(mockClass))
	o.AssertExpectations(t)
}
