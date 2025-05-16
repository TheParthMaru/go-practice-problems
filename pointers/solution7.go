package pointers

import "fmt"

func Solution7() {
	num := 69
	ptrCopyNum := cloneInt(&num)

	fmt.Println("Address of original num:", &num)
	fmt.Println("Address of copyNum:", ptrCopyNum)
	fmt.Println("Value of copyNum:", *ptrCopyNum)
}

func cloneInt(ptr *int) *int {
	copyNum := *ptr

	ptrCopyNum := &copyNum

	return ptrCopyNum
}
