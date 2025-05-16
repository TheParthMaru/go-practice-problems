package pointers

import "fmt"

func Solution6(n int) {
	ptrSlice := allocateSlice(n)

	for _, value := range *ptrSlice {
		fmt.Println(value)
	}
}

func allocateSlice(n int) *[]int {
	mySlice := make([]int, n)

	ptrSlice := &mySlice

	// println(mySlice)
	// println(&mySlice)
	// println(&mySlice[0])
	// println(&mySlice[1])
	// println(ptrSlice)
	// println(*ptrSlice)

	return ptrSlice
}
