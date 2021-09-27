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
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

// eModelElementExt is the extension of the model object 'EFactory'
type eModelElementExt struct {
	eModelElementImpl
}

func newEModelElementExt() *eModelElementExt {
	eElement := new(eModelElementExt)
	eElement.SetInterfaces(eElement)
	eElement.Initialize()
	return eElement
}

func (eModelElement *eModelElementExt) GetEAnnotation(source string) EAnnotation {
	if eModelElement.eAnnotations != nil {
		for itAnnotation := eModelElement.eAnnotations.Iterator(); itAnnotation.HasNext(); {
			annotation := itAnnotation.Next().(EAnnotation)
			if annotation.GetSource() == source {
				return annotation
			}
		}
	}
	return nil
}

func (eModelElement *eModelElementExt) EObjectForFragmentSegment(uriFragmentSegment string) EObject {
	if len(uriFragmentSegment) > 0 {
		// Is the first character a special character, i.e., something other than '@'?
		firstCharacter := uriFragmentSegment[0]
		if firstCharacter != '@' {
			// Is it the start of a source URI of an annotation?
			if firstCharacter == '%' {
				// Find the closing '%' and make sure it's not just the opening '%'
				index := strings.LastIndex(uriFragmentSegment, "%")
				hasCount := false
				if index != 0 {
					hasCount = uriFragmentSegment[index+1] == '.'
					if index == len(uriFragmentSegment)-1 || hasCount {

						// DecodeResource all encoded characters.
						source := ""
						if encodedSource := uriFragmentSegment[1:index]; encodedSource != "%" {
							source, _ = url.PathUnescape(encodedSource)
						}

						// Check for a count, i.e., a '.' followed by a number.
						count := 0
						if hasCount {
							i, err := strconv.Atoi(uriFragmentSegment[index+2:])
							if err == nil {
								count = i
							}
						}

						// Look for the annotation with the matching source.
						for it := eModelElement.AsEObject().EContents().Iterator(); it.HasNext(); {
							eAnnotation, _ := it.Next().(EAnnotation)
							if eAnnotation != nil {
								otherSource := eAnnotation.GetSource()
								if source == otherSource {
									if count == 0 {
										return eAnnotation
									}
									count--
								}
							}
						}
						return nil
					}
				}
			}

			// Look for trailing count.
			index := strings.LastIndex(uriFragmentSegment, ".")
			name := uriFragmentSegment
			if index != -1 {
				name = uriFragmentSegment[:index]
			}
			count := 0
			if index != -1 {
				i, err := strconv.Atoi(uriFragmentSegment[index+1:])
				if err != nil {
					name = uriFragmentSegment
				} else {
					count = i
				}
			}

			if name == "%" {
				name = ""
			} else {
				name = url.PathEscape(name)
			}

			for it := eModelElement.AsEObject().EContents().Iterator(); it.HasNext(); {
				eNamedElement, _ := it.Next().(ENamedElement)
				if eNamedElement != nil {
					otherName := eNamedElement.GetName()
					if name == otherName {
						if count == 0 {
							return eNamedElement
						}
						count--
					}
				}
			}
			return nil
		}
	}
	return eModelElement.eModelElementImpl.EObjectForFragmentSegment(uriFragmentSegment)
}

func (eModelElement *eModelElementExt) EURIFragmentSegment(feature EStructuralFeature, object EObject) string {
	eNamedElement, _ := object.(ENamedElement)
	if eNamedElement != nil {
		name := eNamedElement.GetName()
		count := 0
		for it := eModelElement.EContents().(EObjectList).GetUnResolvedList().Iterator(); it.HasNext(); {
			otherEObject := it.Next().(EObject)
			if otherEObject == object {
				break
			}
			otherENamedElement, _ := otherEObject.(ENamedElement)
			if otherENamedElement != nil {
				otherName := otherENamedElement.GetName()
				if name == otherName {
					count++
				}
			}
		}
		if len(name) == 0 {
			name = "%"
		} else {
			name = url.PathEscape(name)
		}
		if count > 0 {
			return name + "." + fmt.Sprintf("%d", count)
		}
		return name

	}
	eAnnotation, _ := object.(EAnnotation)
	if eAnnotation != nil {
		source := eAnnotation.GetSource()
		count := 0
		for it := eModelElement.EContents().Iterator(); it.HasNext(); {
			otherEObject := it.Next().(EObject)
			if otherEObject == object {
				break
			}
			otherEAnnotation, _ := otherEObject.(EAnnotation)
			if otherEAnnotation != nil {
				otherSource := otherEAnnotation.GetSource()
				if source == otherSource {
					count++
				}
			}
		}

		result := "%"
		if len(source) == 0 {
			result += "%"
		} else {
			result += url.PathEscape(source)
		}
		result += "%"
		if count > 0 {
			result += "." + fmt.Sprintf("%d", count)
		}
		return result
	}
	return eModelElement.eModelElementImpl.EURIFragmentSegment(feature, object)
}
