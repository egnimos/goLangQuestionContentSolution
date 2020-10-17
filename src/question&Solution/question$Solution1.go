/*
QUESTION=>
Initialize a SLICE of int using a composite literal; print out the slice; range over the
slice printing out just the index; range over the slice printing out both the index and the value
solution
*/

package main

import (
	"fmt"
)

func main() {
	sl := []int{1, 23, 56, 12, 67, 78, 98, 100}

	//print the index
	for i, _ := range sl {
		fmt.Println("So the Index:", i)
	}

	for i, v := range sl {
		fmt.Println("so the index is: ", i, "- and the value of that index is: ", v)
	}

}
