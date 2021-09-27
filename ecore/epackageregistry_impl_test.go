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

func TestMockEPackageRegistryImpl_RegisterPackage(t *testing.T) {
	rp := NewEPackageRegistryImpl()
	p := &MockEPackage{}
	p.On("GetNsURI").Return("uri").Once()
	rp.RegisterPackage(p)
	mock.AssertExpectationsForObjects(t, p)
}

func TestMockEPackageRegistryImpl_UnRegisterPackage(t *testing.T) {
	rp := NewEPackageRegistryImpl()
	p := &MockEPackage{}
	p.On("GetNsURI").Return("uri").Once()
	rp.UnregisterPackage(p)
	mock.AssertExpectationsForObjects(t, p)
}

func TestMockEPackageRegistryImpl_RegisterPutPackage(t *testing.T) {
	rp := NewEPackageRegistryImpl()
	p := &MockEPackage{}
	rp.PutPackage("nsURI", p)
	assert.Equal(t, p, rp.GetPackage("nsURI"))
}

func TestMockEPackageRegistryImpl_RegisterPutSupplier(t *testing.T) {
	rp := NewEPackageRegistryImpl()
	p := &MockEPackage{}
	f := func() EPackage {
		return p
	}
	rp.PutSupplier("nsURI", f)
	assert.Equal(t, p, rp.GetPackage("nsURI"))
}

func TestMockEPackageRegistryImpl_Remove(t *testing.T) {
	rp := NewEPackageRegistryImpl()
	p := &MockEPackage{}
	rp.PutPackage("nsURI", p)
	assert.Equal(t, p, rp.GetPackage("nsURI"))
	rp.Remove("nsURI")
	assert.Nil(t, rp.GetPackage("nsURI"))
}

func TestMockEPackageRegistryImpl_GetPackage(t *testing.T) {
	{
		rp := NewEPackageRegistryImpl()
		assert.Nil(t, rp.GetPackage("uri"))
	}
	{
		rp := NewEPackageRegistryImpl()
		p := &MockEPackage{}
		p.On("GetNsURI").Return("uri").Once()
		rp.RegisterPackage(p)
		mock.AssertExpectationsForObjects(t, p)
		assert.Equal(t, p, rp.GetPackage("uri"))
	}
	{
		delegate := &MockEPackageRegistry{}
		p := &MockEPackage{}
		rp := NewEPackageRegistryImplWithDelegate(delegate)
		delegate.On("GetPackage", "uri").Return(p).Once()
		assert.Equal(t, p, rp.GetPackage("uri"))
		mock.AssertExpectationsForObjects(t, p, delegate)
	}
}

func TestMockEPackageRegistryImpl_GetFactory(t *testing.T) {
	{
		rp := NewEPackageRegistryImpl()
		assert.Nil(t, rp.GetFactory("uri"))
	}
	{
		rp := NewEPackageRegistryImpl()
		p := &MockEPackage{}
		p.On("GetNsURI").Return("uri").Once()
		rp.RegisterPackage(p)
		mock.AssertExpectationsForObjects(t, p)

		f := &MockEFactory{}
		p.On("GetEFactoryInstance").Return(f).Once()
		assert.Equal(t, f, rp.GetFactory("uri"))
		mock.AssertExpectationsForObjects(t, p, f)
	}
	{
		delegate := &MockEPackageRegistry{}
		p := &MockEPackage{}
		f := &MockEFactory{}
		rp := NewEPackageRegistryImplWithDelegate(delegate)
		delegate.On("GetFactory", "uri").Return(f).Once()
		assert.Equal(t, f, rp.GetFactory("uri"))
		mock.AssertExpectationsForObjects(t, p, delegate)
	}
}
