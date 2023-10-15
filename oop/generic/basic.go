package generic

import (
	"fmt"
)

func Generic() {
	testScores64 := []float64{
		87.3,
		105,
		63.5,
		27,
	}

	testScores32 := []float32{
		87.3,
		105,
		63.5,
		27,
	}
	c64 := clone(testScores64)
	// note values are same, and addresses are obviously different
	fmt.Println(c64, testScores64[0] == c64[0], &testScores64[0] == &c64[0])

	c32 := clone(testScores32)
	fmt.Println(c32, testScores32[0] == c32[0], &testScores32[0] == &c32[0])
}

func clone[V any](s []V) []V {
	result := make([]V, len(s))
	for i, v := range s {
		result[i] = v
	}

	return result
}
