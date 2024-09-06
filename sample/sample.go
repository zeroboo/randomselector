package sample

import (
	"fmt"

	"github.com/zeroboo/randomselector"
)

func Sample() {

	//Select values randomly with equally rate for each value, 1/3 chance for each value
	value, errSelect := randomselector.SelectValues("1", "2", "3")
	fmt.Println("Select values: ", value, errSelect)

	//Select one of 2 string: "hello", "world" with equal rate (50% chance for each)
	weightValue, errSelect := randomselector.SelectWithWeight(
		randomselector.WeightValue{Value: "1", Weight: 1},
		randomselector.WeightValue{Value: "2", Weight: 1},
	)
	fmt.Println("Select weight value: ", weightValue, errSelect)

}
