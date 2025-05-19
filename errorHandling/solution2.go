package errorHandling

import (
	"fmt"
	"log"
)

var myMap = map[string]string{
	"name":               "parth",
	"age":                "27",
	"relationshipStatus": "Single AF",
}

func Solution2() {
	// key := "isAdult"
	key := "relationshipStatus"
	value, err := lookup(key)

	if err != nil {
		log.Println("error occurred:", err)
		return
	} else {
		fmt.Printf("Key \"%s\" found with value: %s\n", key, value)
	}
}

func lookup(key string) (string, error) {
	// Traditional approach is to use loop
	// for currentKey, value := range myMap {
	// 	if currentKey == key {
	// 		return value, nil
	// 	}
	// }

	// Better way
	value, ok := myMap[key]

	if ok {
		return value, nil
	}

	return "", fmt.Errorf("%s is not a key in myMap", key)
}
