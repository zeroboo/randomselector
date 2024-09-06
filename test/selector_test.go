package test

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zeroboo/randomselector"
)

type AnyObject struct {
	Value string
}

// go test -timeout 30s -run ^TestSelectEqually_Correct$ github.com/zeroboo/randomselector/test -v
func TestSelector_SelectEqually_Correct(t *testing.T) {
	for i := 0; i < 100; i++ {
		selectedValue, err := randomselector.SelectValues("1", "2", "3")
		t.Logf("Selected values: %v", selectedValue)
		_, isString := selectedValue.(string)
		assert.Contains(t, []string{"1", "2", "3"}, selectedValue, "Selected value must be in list")
		assert.Nil(t, err, "No error")
		assert.True(t, isString, "Correct selected type")
	}
}

// go test -timeout 30s -run ^TestSelector_SelectEquallyStruct_Correct$ github.com/zeroboo/randomselector/test -v
func TestSelector_SelectEquallyStruct_Correct(t *testing.T) {
	for i := 0; i < 100; i++ {
		selectedValue, err := randomselector.SelectValues(AnyObject{Value: "1"}, "2", 3)
		selectedType := reflect.TypeOf(selectedValue)
		t.Logf("Selected values: %v, type %v", selectedValue, selectedType)

		assert.Contains(t, []any{AnyObject{Value: "1"}, "2", 3}, selectedValue, "Selected value must be in list")
		assert.Contains(t, []reflect.Type{reflect.TypeOf(AnyObject{Value: "1"}), reflect.TypeOf("2"), reflect.TypeOf(3)},
			selectedType,
			"Correct selected object type")

		assert.Nil(t, err, "No error")
	}
}

// go test -timeout 30s -run ^TestSelector_SelectWeight_Correct$ github.com/zeroboo/randomselector/test -v
func TestSelector_SelectWeight_Correct(t *testing.T) {
	for i := 0; i < 10; i++ {
		selectedValue, err := randomselector.SelectWithWeight(
			randomselector.WeightValue{
				Value:  "2",
				Weight: 1,
			},
			randomselector.WeightValue{
				Value:  3,
				Weight: 1,
			},
			randomselector.WeightValue{
				Value:  AnyObject{Value: "1"},
				Weight: 1.1,
			},
		)
		assert.Nil(t, err, "No error")

		selectedType := reflect.TypeOf(selectedValue)
		t.Logf("Selected values: %v, type %v", selectedValue, selectedType)

		assert.Contains(t, []any{AnyObject{Value: "1"}, "2", 3}, selectedValue, "Selected value must be in list")
		assert.Contains(t, []reflect.Type{reflect.TypeOf(AnyObject{Value: "1"}), reflect.TypeOf("2"), reflect.TypeOf(3)},
			selectedType,
			"Correct selected object type")

		selectedInt, isInt := selectedValue.(int)
		assert.NotNil(t, selectedInt, "Selected value must be int")
		assert.True(t, isInt, "Selected value must be int")

	}
}

// go test -timeout 30s -run ^TestSelector_ZeroWeight_Correct$ github.com/zeroboo/randomselector/test -v
func TestSelector_ZeroWeight_Correct(t *testing.T) {
	for i := 0; i < 10; i++ {
		selectedValue, err := randomselector.SelectWithWeight(
			randomselector.WeightValue{
				Value:  1,
				Weight: 0,
			},
			randomselector.WeightValue{
				Value:  2,
				Weight: 1,
			},
		)

		selectedType := reflect.TypeOf(selectedValue)
		t.Logf("Selected values: %v, type %v", selectedValue, selectedType)

		assert.Equal(t, 2, selectedValue, "Only value has weight selected")
		assert.Contains(t, []reflect.Type{reflect.TypeOf(1)},
			selectedType,
			"Selected value must be int")

		assert.Nil(t, err, "No error")
	}
}
