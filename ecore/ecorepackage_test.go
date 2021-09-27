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

func TestEcorePackageInstance(t *testing.T) {
	ecorePackage := GetPackage()
	assert.Equal(t, ecorePackage.GetNsURI(), "http://www.eclipse.org/emf/2002/Ecore")
	assert.Equal(t, ecorePackage.GetNsPrefix(), "ecore")
}

func TestEcorePackageClassAndAttribute(t *testing.T) {
	ecorePackage := GetPackage()
	eClass := ecorePackage.GetEClass()
	assert.Equal(t, eClass.GetName(), "EClass")
	eClassAbstract := ecorePackage.GetEClass_Abstract()
	assert.Equal(t, eClassAbstract.GetName(), "abstract")
	assert.Equal(t, eClassAbstract.GetEAttributeType().GetName(), "EBoolean")

	eFeatures := eClass.GetEStructuralFeatures()
	assert.True(t, eFeatures.Contains(eClassAbstract))
}
