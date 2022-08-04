package syncmap

import (
	"fmt"
	"sync"
)

type IntInterfaceMap struct {
	m sync.Map
}
func (m IntInterfaceMap) do(){
	m.m.Store("a", []int{1, 2, 3})
	fmt.Println(m.m)
}
