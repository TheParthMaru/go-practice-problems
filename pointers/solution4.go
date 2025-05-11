package pointersSolutions

import "fmt"

func Solution4() {
	mySlice := []int{1, 2, 3, 4}
	resetSlice(&mySlice)
	fmt.Println(mySlice)
}

func resetSlice(ptr *[]int) {
	*ptr = []int{}
}
