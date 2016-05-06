package base

import (
	"errors"
	"sort"
)

type MultiMap struct {
	Keys         []interface{}
	Vals         []interface{}
	sortedTarget []interface{}
	Data         map[interface{}]interface{}
}

func (mMap *MultiMap) Put(k, v interface{}) (bool, error) {
	if _, ok := mMap.Data[k]; !ok {
		mMap.Data[k] = v
	} else {
		return false, errors.New("duplicate keys")
	}
	mMap.Keys = append(mMap.Keys, k)
	mMap.Vals = append(mMap.Vals, v)
	return true, nil
}

func (mMap *MultiMap) SortedByKeys() {
	mMap.sortedTarget = mMap.Keys[:]
	sort.Sort(mMap)
}

func (mMap *MultiMap) SortedByVals() {
	mMap.sortedTarget = mMap.Vals[:]
	sort.Sort(mMap)
}

func (mMap *MultiMap) Len() int { return len(mMap.Vals) }

func (mMap *MultiMap) Swap(i, j int) {
	mMap.Vals[i], mMap.Vals[j] = mMap.Vals[j], mMap.Vals[i]
	mMap.Keys[i], mMap.Keys[j] = mMap.Keys[j], mMap.Keys[i]
}

func (mMap *MultiMap) Less(i, j int) bool {

	switch mMap.sortedTarget[0].(type) {
	case string:
		return mMap.sortedTarget[i].(string) < mMap.sortedTarget[j].(string)
	}
	return false
}
