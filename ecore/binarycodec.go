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
	"io"
)

const (
	BINARY_OPTION_ID_ATTRIBUTE = "ID_ATTRIBUTE" // if true, save id attribute of the object
)

type BinaryCodec struct {
}

func (bc *BinaryCodec) NewEncoder(resource EResource, w io.Writer, options map[string]interface{}) EResourceEncoder {
	return NewBinaryEncoder(resource, w, options)
}
func (bc *BinaryCodec) NewDecoder(resource EResource, r io.Reader, options map[string]interface{}) EResourceDecoder {
	return NewBinaryDecoder(resource, r, options)
}

type binaryFeatureKind int

var binaryDebug = false

const (
	bfkObjectContainer binaryFeatureKind = iota
	bfkObjectContainerProxy
	bfkObject
	bfkObjectProxy
	bfkObjectList
	bfkObjectListProxy
	bfkObjectContainment
	bfkObjectContainmentProxy
	bfkObjectContainmentList
	bfkObjectContainmentListProxy
	bfkFloat64
	bfkFloat32
	bfkInt
	bfkInt64
	bfkInt32
	bfkInt16
	bfkByte
	bfkBool
	bfkString
	bfkByteArray
	bfkData
	bfkDataList
	bfkEnum
	bfkDate
)

func getBinaryCodecFeatureKind(eFeature EStructuralFeature) binaryFeatureKind {
	if eReference, _ := eFeature.(EReference); eReference != nil {
		if eReference.IsContainment() {
			if eReference.IsResolveProxies() {
				if eReference.IsMany() {
					return bfkObjectContainmentListProxy
				} else {
					return bfkObjectContainmentProxy
				}
			} else {
				if eReference.IsMany() {
					return bfkObjectContainmentList
				} else {
					return bfkObjectContainment
				}
			}
		} else if eReference.IsContainer() {
			if eReference.IsResolveProxies() {
				return bfkObjectContainerProxy
			} else {
				return bfkObjectContainer
			}
		} else if eReference.IsResolveProxies() {
			if eReference.IsMany() {
				return bfkObjectListProxy
			} else {
				return bfkObjectProxy
			}
		} else {
			if eReference.IsMany() {
				return bfkObjectList
			} else {
				return bfkObject
			}
		}
	} else if eAttribute, _ := eFeature.(EAttribute); eAttribute != nil {
		if eAttribute.IsMany() {
			return bfkDataList
		} else {
			eDataType := eAttribute.GetEAttributeType()
			if eEnum, _ := eDataType.(EEnum); eEnum != nil {
				return bfkEnum
			}

			switch eDataType.GetInstanceTypeName() {
			case "float64":
				return bfkFloat64
			case "float32":
				return bfkFloat32
			case "int":
				return bfkInt
			case "int64":
				return bfkInt64
			case "int32":
				return bfkInt32
			case "int16":
				return bfkInt16
			case "byte":
				return bfkByte
			case "bool":
				return bfkBool
			case "string":
				return bfkString
			case "[]byte":
				return bfkByteArray
			case "*time.Time":
				return bfkDate
			}

			return bfkData
		}
	}
	return -1
}
