package errorHandling

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Solution3() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter two integers: ")
	input, readerError := reader.ReadString('\n')

	if readerError != nil {
		log.Println("error occurred in reading:", readerError)
		return
	}

	// Extracting parts
	parts := strings.Fields(strings.TrimSpace(input))

	if len(parts) != 2 {
		log.Println("please enter exactly two space-separated integers")
		return
	}

	num1, err1 := strconv.Atoi(parts[0])
	num2, err2 := strconv.Atoi(parts[1])

	if err1 != nil || err2 != nil {
		log.Println("invalid input: both values must be integers (no decimals)")
		return
	}

	result, err := safeDivide(num1, num2)
	if err != nil {
		log.Println("error occurred:", err)
		return
	}

	fmt.Println("Result:", result)
}

func safeDivide(num1, num2 int) (float64, error) {
	if num2 == 0 {
		return 0, fmt.Errorf("denominator cannot be zero")
	} else {
		return (float64(num1)) / (float64(num2)), nil
	}
}
