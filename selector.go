package randomselector

import (
	"errors"
	"fmt"
	"math/rand"
)

func init() {

}

// Selects values randomly with equally rate for each value
//
// Return:
//   - selected value if selecting is successful
//   - error if any, nil if selecting is successful
func SelectValues(values ...any) (any, error) {
	if len(values) == 0 {
		return nil, errors.New("no value to select")
	}
	index := rand.Intn(len(values))
	return values[index], nil
}

type WeightValue struct {
	Value  any
	Weight float64
}

func makeAccRates(values []float64) []float64 {
	accRate := float64(0)
	result := make([]float64, len(values))
	for i := 1; i < len(values); i++ {
		accRate += values[i]
		result[i] = accRate
	}
	return result
}
func randomIndex(accRates []float64) int {
	r := rand.Float64()
	for i, v := range accRates {
		if r < v {
			return i
		}
	}
	return -1
}

// Randomly selects one of weighted values
//
// Return:
//   - selected value
//   - error if any, nil if selecting is successful
func SelectWithWeight(values ...WeightValue) (any, error) {
	if len(values) == 0 {
		return nil, errors.New("no value to select")
	}

	var maxWeight float64 = 0
	var weights []float64 = make([]float64, len(values))
	for i, v := range values {
		if v.Weight < 0 {
			return nil, fmt.Errorf("invalid weight %v at index %v", v.Weight, i)
		}
		maxWeight += v.Weight
		weights[i] = v.Weight
	}
	accRates := makeAccRates(weights)
	selectedIndex := randomIndex(accRates)
	if selectedIndex < 0 {
		return nil, fmt.Errorf("selected failed, index=%v", selectedIndex)
	}

	return values[selectedIndex].Value, nil
}
