package test

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zeroboo/randomselector/randombag"
	"github.com/zeroboo/randomselector/sample"
)

func TestMain(m *testing.M) {
	fmt.Println("Test randomselector main function")

	exitCode := m.Run()

	os.Exit(exitCode)
}

func TestSample_NoError(t *testing.T) {
	fmt.Println("Test sample function")
	sample.Sample()

}

// go test -timeout 30s -run ^TestRandomBagCreating_Correct$ github.com/zeroboo/randomselector/test -v
func TestRandomBagCreating_Correct(t *testing.T) {
	var randomBox *randombag.RandomBag = randombag.NewRandomBag(1235, true,
		*randombag.NewRandomItem("1", 100, "string 1"),
		*randombag.NewRandomItem("2", 200, "string 2"),
		*randombag.NewRandomItem("3", 300, "string 3"),
	)

	assert.Equal(t, int64(1235), randomBox.GetMaxRate(), "Correct max rate")
	assert.Equal(t, 3, len(randomBox.GetContents()), "Correct content size")
	assert.Equal(t, "[100 300 600]", fmt.Sprintf("%v", randomBox.GetAccRates()), "Correct content size")
}

// Full rate means box has no chance of missing in selecting
// go test -timeout 30s -run ^TestRandomBag_SelectingFullRate_NoNilResult$ github.com/zeroboo/randomselector/test -v
func TestRandomBag_SelectingFullRate_NoNilResult(t *testing.T) {
	var randomBox *randombag.RandomBag = randombag.NewRandomBag(1000, true,
		*randombag.NewRandomItem("1", 100, "string 1"),
		*randombag.NewRandomItem("2", 200, "string 2"),
		*randombag.NewRandomItem("3", 300, "string 3"),
		*randombag.NewRandomItem("4", 400, "string 5"),
	)

	assert.Equal(t, int64(1000), randomBox.GetMaxRate(), "Correct max rate")
	assert.Equal(t, 4, len(randomBox.GetContents()), "Correct content size")
	assert.Equal(t, "[100 300 600 1000]", fmt.Sprintf("%v", randomBox.GetAccRates()), "Correct content size")
	for i := 0; i < 100; i++ {
		selectedValue, _ := randomBox.SelectRandom()
		t.Logf("Selected values: %v", selectedValue)
		assert.NotEqual(t, nil, selectedValue, "Selected value must not be nil")
	}

}

// TestRandomBag_SelectingFullRateWithStruct_NoNilResult: test box full rate with content is a struct
// go test -timeout 30s -run ^TestRandomBag_SelectingFullRateWithStruct_NoNilResult$ github.com/zeroboo/randomselector/test -v
func TestRandomBag_SelectingFullRateWithStruct_NoNilResult(t *testing.T) {
	var randomBox *randombag.RandomBag = randombag.NewRandomBag(1000, true,
		*randombag.NewRandomItem("1", 100, TestRandomItem{ID: "item1", Value: 1}),
		*randombag.NewRandomItem("2", 200, TestRandomItem{ID: "item2", Value: 2}),
		*randombag.NewRandomItem("3", 300, TestRandomItem{ID: "item3", Value: 3}),
		*randombag.NewRandomItem("4", 400, TestRandomItem{ID: "item4", Value: 4}),
	)

	assert.Equal(t, int64(1000), randomBox.GetMaxRate(), "Correct max rate")
	assert.Equal(t, 4, len(randomBox.GetContents()), "Correct content size")
	assert.Equal(t, "[100 300 600 1000]", fmt.Sprintf("%v", randomBox.GetAccRates()), "Correct content size")
	for i := 0; i < 100; i++ {
		selectedValue, _ := randomBox.SelectRandom()
		t.Logf("Selected values: %v %v", selectedValue, selectedValue.(TestRandomItem).GetContent())
		assert.NotEqual(t, nil, selectedValue, "Selected value must not be nil")
		assert.IsType(t, TestRandomItem{}, selectedValue, "Correct selected type")
	}

}

// TestRandomBag_SelectedNil_Correct: test box full rate with content is a struct
// go test -timeout 30s -run ^TestRandomBag_SelectedNil_Correct$ github.com/zeroboo/randomselector/test -v
func TestRandomBag_SelectedNil_Correct(t *testing.T) {
	var randomBox *randombag.RandomBag = randombag.NewRandomBag(1000,
		true,
		*randombag.NewRandomItem("1", 0, TestRandomItem{ID: "item1", Value: 2, Rate: 0}),
		*randombag.NewRandomItem("1", 0, TestRandomItem{ID: "item2", Value: 2, Rate: 0}),
	)

	assert.Equal(t, int64(1000), randomBox.GetMaxRate(), "Correct max rate")
	assert.Equal(t, 2, len(randomBox.GetContents()), "Correct content size")
	assert.Equal(t, "[0 0]", fmt.Sprintf("%v", randomBox.GetAccRates()), "Correct content size")
	for i := 0; i < 100; i++ {
		selectedValue, _ := randomBox.SelectRandom()
		t.Logf("Selected values: %v", selectedValue)
		assert.Equal(t, nil, selectedValue, "Selected value must be nil")
	}
}

// go test -timeout 30s -run ^TestAddItemToBag_Correct$ github.com/zeroboo/randomselector/test -v
func TestAddItemToBag_Correct(t *testing.T) {
	item1 := TestRandomItem{
		ID:    "item1",
		Value: 1,
		Rate:  1,
	}
	item2 := TestRandomItem{
		ID:    "item2",
		Value: 2,
		Rate:  2,
	}
	bag := randombag.NewRandomBag(randombag.RandomRateNone, true)

	bag.AddItem(item1)
	bag.AddItem(item2)
	bag.AddItem(item2)

	t.Logf("Bag: %v", bag.String())
	assert.Equal(t, int64(5), bag.GetMaxRate(), "Correct max rates")
	assert.Equal(t, 3, len(bag.GetContents()), "Correct item counts")

}
