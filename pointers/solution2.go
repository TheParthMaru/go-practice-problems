package pointers

import "fmt"

func Solution2() {
	num := 5
	square(&num)
	fmt.Println(num)
}

func square(ptr *int) {
	*ptr = *ptr * *ptr
}
