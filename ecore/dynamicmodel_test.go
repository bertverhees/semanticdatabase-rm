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

type DynamicMetaModel struct {
	bookStoreEPackage EPackage
	bookStoreEFactory EFactory
	bookStoreEClass   EClass
	bookStoreOwner    EAttribute
	bookStoreLocation EAttribute
	bookStoreBooks    EReference
	bookEClass        EClass
	bookName          EAttribute
	bookISBN          EAttribute
}

func createDynamicMetaModel() *DynamicMetaModel {
	m := new(DynamicMetaModel)
	/*
	* Create EClass instance to model BookStore class
	 */
	m.bookStoreEClass = GetFactory().CreateEClass()

	/*
	 * Create EClass instance to model Book class
	 */
	m.bookEClass = GetFactory().CreateEClass()

	/*
	 * Instantiate EPackage and provide unique URI
	 * to identify this package
	 */
	m.bookStoreEFactory = GetFactory().CreateEFactory()

	m.bookStoreEPackage = GetFactory().CreateEPackage()
	m.bookStoreEPackage.SetName("BookStorePackage")
	m.bookStoreEPackage.SetNsPrefix("bookStore")
	m.bookStoreEPackage.SetNsURI("http:///com.ibm.dynamic.example.bookstore.ecore")
	m.bookStoreEPackage.SetEFactoryInstance(m.bookStoreEFactory)

	/*
	 * Create attributes for BookStore class as specified in the model
	 */
	m.bookStoreOwner = GetFactory().CreateEAttribute()
	m.bookStoreOwner.SetName("owner")
	m.bookStoreOwner.SetEType(GetPackage().GetEString())

	m.bookStoreLocation = GetFactory().CreateEAttribute()
	m.bookStoreLocation.SetName("location")
	m.bookStoreLocation.SetEType(GetPackage().GetEString())

	m.bookStoreBooks = GetFactory().CreateEReference()
	m.bookStoreBooks.SetName("books")
	m.bookStoreBooks.SetEType(m.bookEClass)
	m.bookStoreBooks.SetUpperBound(UNBOUNDED_MULTIPLICITY)
	m.bookStoreBooks.SetContainment(true)

	/*
	 * Create attributes for Book class as defined in the model
	 */
	m.bookName = GetFactory().CreateEAttribute()
	m.bookName.SetName("name")
	m.bookName.SetEType(GetPackage().GetEString())

	m.bookISBN = GetFactory().CreateEAttribute()
	m.bookISBN.SetName("isbn")
	m.bookISBN.SetEType(GetPackage().GetEInt())

	/*
	 * Add owner, location and books attributes/references
	 * to BookStore class
	 */
	m.bookStoreEClass.GetEStructuralFeatures().Add(m.bookStoreOwner)
	m.bookStoreEClass.GetEStructuralFeatures().Add(m.bookStoreLocation)
	m.bookStoreEClass.GetEStructuralFeatures().Add(m.bookStoreBooks)

	/*
	 * Add name and isbn attributes to Book class
	 */
	m.bookEClass.GetEStructuralFeatures().Add(m.bookName)
	m.bookEClass.GetEStructuralFeatures().Add(m.bookISBN)

	/*
	 * Place BookStore and Book classes in bookStoreEPackage
	 */
	m.bookStoreEPackage.GetEClassifiers().Add(m.bookStoreEClass)
	m.bookStoreEPackage.GetEClassifiers().Add(m.bookEClass)

	return m
}

type DynamicModel struct {
	bookObject      EObject
	bookStoreObject EObject
}

func instanciateDynamicModel(mm *DynamicMetaModel) *DynamicModel {
	m := new(DynamicModel)

	bookFactoryInstance := mm.bookStoreEPackage.GetEFactoryInstance()
	/*
	 * Create dynamic instance of BookStoreEClass and BookEClass
	 */
	m.bookObject = bookFactoryInstance.Create(mm.bookEClass)
	m.bookStoreObject = bookFactoryInstance.Create(mm.bookStoreEClass)

	/*
	 * Set the values of bookStoreObject attributes
	 */
	m.bookStoreObject.ESet(mm.bookStoreOwner, "David Brown")
	m.bookStoreObject.ESet(mm.bookStoreLocation, "Street#12, Top Town, NY")
	allBooks := m.bookStoreObject.EGet(mm.bookStoreBooks).(EList)
	allBooks.Add(m.bookObject)

	/*
	 * Set the values of bookObject attributes
	 */
	m.bookObject.ESet(mm.bookName, "Harry Potter and the Deathly Hallows")
	m.bookObject.ESet(mm.bookISBN, 157221)
	return m
}

func TestDynamicModel(t *testing.T) {
	mm := createDynamicMetaModel()
	m := instanciateDynamicModel(mm)

	assert.Equal(t, 0, mm.bookStoreEClass.GetClassifierID())
	assert.Equal(t, 1, mm.bookEClass.GetClassifierID())

	assert.Equal(t, mm.bookStoreEClass, m.bookStoreObject.EClass())
	assert.Equal(t, mm.bookEClass, m.bookObject.EClass())

	/*
	 * Read/Get the values of bookStoreObject attributes
	 */
	assert.Equal(t, "David Brown", m.bookStoreObject.EGet(mm.bookStoreOwner).(string))
	assert.Equal(t, "Street#12, Top Town, NY", m.bookStoreObject.EGet(mm.bookStoreLocation).(string))

	/*
	 * Read/Get the values of bookObject attributes
	 */
	assert.Equal(t, "Harry Potter and the Deathly Hallows", m.bookObject.EGet(mm.bookName).(string))
	assert.Equal(t, 157221, m.bookObject.EGet(mm.bookISBN).(int))

}

func TestGetURINoResource(t *testing.T) {
	mm := createDynamicMetaModel()
	m := instanciateDynamicModel(mm)
	assert.Equal(t, &URI{Fragment: "//"}, GetURI(m.bookStoreObject))
	assert.Equal(t, &URI{Fragment: "//@books.0"}, GetURI(m.bookObject))
}

func TestGetURIResource(t *testing.T) {
	mm := createDynamicMetaModel()
	m := instanciateDynamicModel(mm)
	r := NewEResourceImpl()
	r.SetURI(&URI{Scheme: "file",
		Path: "a.test",
	})
	c := r.GetContents()
	c.Add(m.bookStoreObject)
	assert.Equal(t, &URI{Scheme: "file",
		Path:     "a.test",
		Fragment: "/",
	}, GetURI(m.bookStoreObject))
	assert.Equal(t, &URI{Scheme: "file",
		Path:     "a.test",
		Fragment: "//@books.0",
	}, GetURI(m.bookObject))
}

func TestGetEObjectFromRoot(t *testing.T) {
	mm := createDynamicMetaModel()
	m := instanciateDynamicModel(mm)
	assert.Equal(t, m.bookObject, GetEObject(m.bookStoreObject, "@books.0"))
}
