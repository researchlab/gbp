package examples

import (
	"fmt"
	"github.com/researchlab/golearning/base"
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
	keys := []string{"11", "120", "250", "27", "90", "0", "2", "4", "6", "8"}
	vals := []string{"0", "15", "12", "9", "80", "210", "31", "17", "400", "6"}
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
