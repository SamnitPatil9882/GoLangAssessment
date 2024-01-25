package main

import (
	"fmt"
)

func accessSlice(slc []int, ind int) (int, error) {

	if ind > len(slc)-1 {
		return 0, fmt.Errorf("Error occured")
	}

	return ind, nil
}
func main() {
	arr := [4]int{1, 2, 3, 4}
	slc := arr[:4]
	// res1, err1 := accessSlice(slc, 4)
	res1, err1 := accessSlice(slc, 2)
	if err1 != nil {
		fmt.Print(err1)
	} else {
		fmt.Printf("item %d, value %d\n", res1, slc[res1])
	}
	// res2, err2 := accessSlice(slc, 4)

}
