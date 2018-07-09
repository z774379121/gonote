package main

import (
	"fmt"
)

func main() {
	slice := []int{0, 1, 2, 3}
	myMap := make(map[int]*int)

	for index, value := range slice {
		myMap[index] = &value
	}
	fmt.Println("=====new map=====")
	prtMap(myMap)
}

func prtMap(myMap map[int]*int) {
	for key, value := range myMap {
		fmt.Printf("map[%v]=%v\n", key, *value)
	}
}

/*func main() {
	slice := []int{0, 1, 2, 3}
	myMap := make(map[int]*int)

	for index, value := range slice {
		num := value
		myMap[index] = &num
	}
	fmt.Println("=====new map=====")
	prtMap(myMap)
}

func prtMap(myMap map[int]*int) {
	for key, value := range myMap {
		fmt.Printf("map[%v]=%v\n", key, *value)
	}
}
*/

// 并发安全map
/*type SafeMap struct {
	sync.RWMutex
	Map map[int]int
}

func main() {
	safeMap := newSafeMap()

	for i := 0; i < 100000; i++ {
		go safeMap.writeMap(i, i)
		go safeMap.readMap(i)
	}

}

func newSafeMap() *SafeMap {
	sm := new(SafeMap)
	sm.Map = make(map[int]int)
	return sm

}

func (sm *SafeMap) readMap(key int) int {
	sm.RLock()
	value := sm.Map[key]
	sm.RUnlock()
	return value
}

func (sm *SafeMap) writeMap(key int, value int) {
	sm.Lock()
	sm.Map[key] = value
	sm.Unlock()
}*/


// easy read fmt
/*type User struct {
	Name string
	Age  int
	Addr string
}

func main() {
	u := User{"user1", 12, "addr"}
	fmt.Printf("%v\n", u)
	fmt.Printf("%+v\n", u)
	fmt.Printf("%#v\n", u)
}*/