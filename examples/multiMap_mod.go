package examples

import (
	"fmt"
	"github.com/researchlab/go-learning/base"
)

func MultiMapMgr() {
	mMap := MultiMapInit()
	MultiMapRawOrder(mMap)
	MultiMapKeysOrder(mMap)
	MultiMapValsOrder(mMap)
}
func MultiMapInit() *base.MultiMap {
	MMap := &base.MultiMap{
		Data: make(map[interface{}]interface{}, 0),
	}
	keys := []string{"1", "1", "5", "7", "9", "0", "2", "4", "6", "8"}
	vals := []string{"0", "5", "2", "9", "8", "1", "3", "7", "4", "6"}
	for i := 0; i < 10; i++ {
		ok, err := MMap.Put(keys[i], vals[i])
		fmt.Println(ok, ":", err)
	}
	return MMap
}

func MultiMapRawOrder(MMap *base.MultiMap) {

	fmt.Println("rawOrder")
	for _, v := range MMap.Keys {
		fmt.Println(v, "=", MMap.Data[v])
	}

}

func MultiMapKeysOrder(MMap *base.MultiMap) {
	fmt.Println("keysOrder")
	MMap.SortedByKeys()
	for _, v := range MMap.Keys {
		fmt.Println(v, "=", MMap.Data[v])
	}

}

func MultiMapValsOrder(MMap *base.MultiMap) {
	fmt.Println("valsOrder")
	MMap.SortedByVals()
	for _, v := range MMap.Keys {
		fmt.Println(v, "=", MMap.Data[v])
	}

}
