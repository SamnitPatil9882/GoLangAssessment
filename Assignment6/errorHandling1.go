package main

import "fmt"

func accessSlice(slc []int, ind int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("internal error:%v\n", r)
		}
	}()
	fmt.Printf("item %d, value %d\n", ind, slc[ind])
}
func main() {
	arr := [4]int{1, 2, 3, 4}
	slc := arr[:4]
	accessSlice(slc, 2)
	accessSlice(slc, 4)

}
