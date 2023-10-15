package generic

import (
	"fmt"
)

func genericExtended() {
	testScores64 := map[string]float64{
		"Root":     87.3,
		"Germey":   105,
		"Nick":     63.5,
		"Virginia": 27,
	}

	testScores32 := map[string]float32{
		"Root":     87.3,
		"Germey":   105,
		"Nick":     63.5,
		"Virginia": 27,
	}

	c64 := clone(testScores64)
	fmt.Println(c64)
	c32 := clone(testScores32)
	fmt.Println(c32)
}

/*
*
Note here the Key type has a lower bound of Comparable,
the key has to implement comparable type, so it cannot be clone[K,V any],
has to be clone[K comparable, V any]
*/
func clone[K comparable, V any](m map[K]V) map[K]V {
	result := make(map[K]V, len(m))
	for k, v := range m {
		result[k] = v
	}

	return result
}
