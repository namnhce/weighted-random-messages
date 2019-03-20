package random

import (
	"math/rand"
	"time"
)

// Define Max Value
const (
	MaxWeight         = 1000000000
	MaxVersionsLength = 100
)

// F that function will give a random version from gave weighted version list.
func F(in []int64) int64 {

	if err := validate(in); err != nil {
		return -1
	}

	rand.Seed(time.Now().UTC().UnixNano())
	accumulatedArray := accumulate(in)

	maxVal := accumulatedArray[len(accumulatedArray)-1]
	// Random value from 0 to maxValue
	out := rand.Int63n(maxVal)

	foundIndex := findIndex(accumulatedArray, out, 0, len(in)-1)

	return in[foundIndex]
}

func accumulate(in []int64) []int64 {
	out := []int64{in[0]}

	for i := 1; i < len(in); i++ {
		out = append(out, out[i-1]+in[i])
	}

	return out
}

func validate(in []int64) error {
	sizeIn := len(in)

	if sizeIn < 1 || sizeIn > MaxVersionsLength {
		return errorVersionLengthInvalid
	}

	for i := 0; i < len(in); i++ {
		if in[i] < 0 || in[i] > MaxWeight {
			return errorWeightInvalid
		}
	}

	return nil
}

func findIndex(ac []int64, in int64, l, h int) int {
	var m int

	for l < h {
		m = (l + h) / 2
		if in > ac[m] {
			l = m + 1
			continue
		}

		h = m
	}

	if ac[l] >= in {
		return l
	}

	return -1
}
