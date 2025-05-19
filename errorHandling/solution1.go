package errorHandling

import (
	"fmt"
	"log"
	"math"
)

func Solution1() {
	// result, err := sqrt(69.69)
	result, err := sqrt(-98)

	if err != nil {
		log.Println("error occurred:", err)
		// fmt.Println(err)
		// log.Fatal("error occurred:", err)
	} else {
		fmt.Println("Result:", result)
	}
	fmt.Println("Program Exit")
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0.00, fmt.Errorf("input num %f cannot be < than 0", x)
	} else {
		return math.Sqrt(x), nil
	}
}
