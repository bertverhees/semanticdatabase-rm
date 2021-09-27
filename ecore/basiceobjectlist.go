// *****************************************************************************
// Copyright(c) 2021 MASA Group
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// *****************************************************************************

package ecore

import "strconv"

type basicEObjectList struct {
	BasicENotifyingList
	owner            EObjectInternal
	featureID        int
	inverseFeatureID int
	containment      bool
	inverse          bool
	opposite         bool
	proxies          bool
	unset            bool
}

func NewBasicEObjectList(owner EObjectInternal, featureID int, inverseFeatureID int, containment, inverse, opposite, proxies, unset bool) *basicEObjectList {
	l := new(basicEObjectList)
	l.interfaces = l
	l.data = []interface{}{}
	l.isUnique = true
	l.owner = owner
	l.featureID = featureID
	l.inverseFeatureID = inverseFeatureID
	l.containment = containment
	l.inverse = inverse
	l.opposite = opposite
	l.proxies = proxies
	l.unset = unset
	return l
}

// GetNotifier ...
func (list *basicEObjectList) GetNotifier() ENotifier {
	return list.owner
}

// GetFeature ...
func (list *basicEObjectList) GetFeature() EStructuralFeature {
	if list.owner != nil {
		return list.owner.EClass().GetEStructuralFeature(list.featureID)
	}
	return nil
}

// GetFeatureID ...
func (list *basicEObjectList) GetFeatureID() int {
	return list.featureID
}

// GetUnResolvedList ...
func (list *basicEObjectList) GetUnResolvedList() EList {
	if list.proxies {
		u := new(unResolvedBasicEObjectList)
		u.delegate = list
		return u
	}
	return list
}

func (list *basicEObjectList) IndexOf(elem interface{}) int {
	if list.proxies {
		for i, value := range list.data {
			if value == elem || list.resolve(i, value) == elem {
				return i
			}
		}
		return -1
	}
	return list.basicEList.IndexOf(elem)
}

func (list *basicEObjectList) doGet(index int) interface{} {
	return list.resolve(index, list.basicEList.doGet(index))
}

func (list *basicEObjectList) resolve(index int, object interface{}) interface{} {
	resolved := list.resolveProxy(object.(EObject))
	if resolved != object {
		list.basicEList.doSet(index, resolved)
		var notifications ENotificationChain
		if list.containment {
			notifications = list.interfaces.(eNotifyingListInternal).inverseRemove(object, notifications)
			if resolvedInternal, _ := resolved.(EObjectInternal); resolvedInternal != nil && resolvedInternal.EInternalContainer() == nil {
				notifications = list.interfaces.(eNotifyingListInternal).inverseAdd(resolved, notifications)
			}
		}
		list.createAndDispatchNotification(notifications, RESOLVE, object, resolved, index)
	}
	return resolved
}

func (list *basicEObjectList) resolveProxy(eObject EObject) EObject {
	if list.proxies && eObject.EIsProxy() {
		return list.owner.(EObjectInternal).EResolveProxy(eObject)
	}
	return eObject
}

func (list *basicEObjectList) inverseAdd(object interface{}, notifications ENotificationChain) ENotificationChain {
	internal, _ := object.(EObjectInternal)
	if internal != nil && list.inverse {
		if list.opposite {
			return internal.EInverseAdd(list.owner, list.inverseFeatureID, notifications)
		} else {
			return internal.EInverseAdd(list.owner, EOPPOSITE_FEATURE_BASE-list.featureID, notifications)
		}
	}
	return notifications
}

func (list *basicEObjectList) inverseRemove(object interface{}, notifications ENotificationChain) ENotificationChain {
	internal, _ := object.(EObjectInternal)
	if internal != nil && list.inverse {
		if list.opposite {
			return internal.EInverseRemove(list.owner, list.inverseFeatureID, notifications)
		} else {
			return internal.EInverseRemove(list.owner, EOPPOSITE_FEATURE_BASE-list.featureID, notifications)
		}
	}
	return notifications
}

type unResolvedBasicEObjectList struct {
	delegate *basicEObjectList
}

func (l *unResolvedBasicEObjectList) Add(elem interface{}) bool {
	if l.delegate.isUnique && l.Contains(elem) {
		return false
	}
	l.delegate.interfaces.(abstractEList).doAdd(elem)
	return true
}

// AddAll elements of an list in the current one
func (l *unResolvedBasicEObjectList) AddAll(list EList) bool {
	if l.delegate.isUnique {
		list = getNonDuplicates(list, l)
		if list.Size() == 0 {
			return false
		}
	}
	l.delegate.interfaces.(abstractEList).doAddAll(list)
	return true
}

// Insert an element in the list
func (l *unResolvedBasicEObjectList) Insert(index int, elem interface{}) bool {
	if index < 0 || index > l.Size() {
		panic("Index out of bounds: index=" + strconv.Itoa(index) + " size=" + strconv.Itoa(l.Size()))
	}
	if l.delegate.isUnique && l.Contains(elem) {
		return false
	}
	l.delegate.interfaces.(abstractEList).doInsert(index, elem)
	return true
}

// InsertAll element of an list at a given position
func (l *unResolvedBasicEObjectList) InsertAll(index int, list EList) bool {
	if index < 0 || index > l.Size() {
		panic("Index out of bounds: index=" + strconv.Itoa(index) + " size=" + strconv.Itoa(l.Size()))
	}
	if l.delegate.isUnique {
		list = getNonDuplicates(list, l)
		if list.Size() == 0 {
			return false
		}
	}
	l.delegate.interfaces.(abstractEList).doInsertAll(index, list)
	return true
}

// Move an element to the given index
func (l *unResolvedBasicEObjectList) MoveObject(newIndex int, elem interface{}) {
	oldIndex := l.IndexOf(elem)
	if oldIndex == -1 {
		panic("Object not found")
	}
	l.delegate.interfaces.(abstractEList).doMove(oldIndex, newIndex)
}

// Swap move an element from oldIndex to newIndex
func (l *unResolvedBasicEObjectList) Move(oldIndex, newIndex int) interface{} {
	return l.delegate.interfaces.(abstractEList).doMove(oldIndex, newIndex)
}

// RemoveAt remove an element at a given position
func (l *unResolvedBasicEObjectList) RemoveAt(index int) interface{} {
	return l.delegate.interfaces.(abstractEList).doRemove(index)
}

// Remove an element in an list
func (l *unResolvedBasicEObjectList) Remove(elem interface{}) bool {
	index := l.IndexOf(elem)
	if index == -1 {
		return false
	}
	l.delegate.interfaces.(abstractEList).doRemove(index)
	return true
}

func (l *unResolvedBasicEObjectList) RemoveAll(collection EList) bool {
	modified := false
	for i := l.Size() - 1; i >= 0; {
		if collection.Contains(l.Get(i)) {
			l.RemoveAt(i)
			modified = true
		}
		i--
	}
	return modified
}

// Get an element of the list
func (l *unResolvedBasicEObjectList) Get(index int) interface{} {
	if index < 0 || index >= l.Size() {
		panic("Index out of bounds: index=" + strconv.Itoa(index) + " size=" + strconv.Itoa(l.Size()))
	}
	return l.delegate.data[index]
}

// Set an element of the list
func (l *unResolvedBasicEObjectList) Set(index int, elem interface{}) interface{} {
	if index < 0 || index >= l.Size() {
		panic("Index out of bounds: index=" + strconv.Itoa(index) + " size=" + strconv.Itoa(l.Size()))
	}
	if l.delegate.isUnique {
		currIndex := l.IndexOf(elem)
		if currIndex >= 0 && currIndex != index {
			panic("element already in list")
		}
	}
	return l.delegate.interfaces.(abstractEList).doSet(index, elem)
}

// Size count the number of element in the list
func (l *unResolvedBasicEObjectList) Size() int {
	return l.delegate.Size()
}

// Clear remove all elements of the list
func (l *unResolvedBasicEObjectList) Clear() {
	l.delegate.Clear()
}

// Empty return true if the list contains 0 element
func (l *unResolvedBasicEObjectList) Empty() bool {
	return l.delegate.Empty()
}

// Contains return if an list contains or not an element
func (l *unResolvedBasicEObjectList) Contains(elem interface{}) bool {
	return l.IndexOf(elem) != -1
}

// IndexOf return the index on an element in an list, else return -1
func (l *unResolvedBasicEObjectList) IndexOf(elem interface{}) int {
	return l.delegate.basicEList.IndexOf(elem)
}

// Iterator through the list
func (l *unResolvedBasicEObjectList) Iterator() EIterator {
	return &listIterator{list: l}
}

func (l *unResolvedBasicEObjectList) ToArray() []interface{} {
	return l.delegate.ToArray()
}

func (l *unResolvedBasicEObjectList) GetNotifier() ENotifier {
	return l.delegate.GetNotifier()
}

func (l *unResolvedBasicEObjectList) GetFeature() EStructuralFeature {
	return l.delegate.GetFeature()
}

func (l *unResolvedBasicEObjectList) GetFeatureID() int {
	return l.delegate.GetFeatureID()
}

func (l *unResolvedBasicEObjectList) AddWithNotification(object interface{}, notifications ENotificationChain) ENotificationChain {
	return l.delegate.AddWithNotification(object, notifications)
}

func (l *unResolvedBasicEObjectList) RemoveWithNotification(object interface{}, notifications ENotificationChain) ENotificationChain {
	index := l.delegate.basicEList.IndexOf(object)
	if index != -1 {
		oldObject := l.delegate.basicEList.doRemove(index)
		return l.delegate.createAndAddNotification(notifications, REMOVE, oldObject, nil, index)
	}
	return notifications
}

func (l *unResolvedBasicEObjectList) SetWithNotification(index int, object interface{}, notifications ENotificationChain) ENotificationChain {
	return l.delegate.SetWithNotification(index, object, notifications)
}

func (l *unResolvedBasicEObjectList) GetUnResolvedList() EList {
	return l
}
