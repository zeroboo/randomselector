package randomselector

import (
	"fmt"
	"os"
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	log.Println("Test randomselector main function")
	log.SetFormatter(&log.TextFormatter{
		DisableQuote: true,
	})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.TraceLevel)

	exitCode := m.Run()

	os.Exit(exitCode)
}

// go test -timeout 30s -run ^TestRandomPick$ github.com/zeroboo/randomselector -v
func TestRandomPick(t *testing.T) {
	log.Printf("HelloWorld")
	assert.Equal(t, "Hello", "Hello", "")
}

// go test -timeout 30s -run ^TestRandomBagCreating_Correct$ github.com/zeroboo/randomselector -v
func TestRandomBagCreating_Correct(t *testing.T) {
	var randomBox *RandomBag = CreateRandomBox(1235, true, RandomContent{
		Name:    "1",
		Content: "string 1",
		Rate:    100,
	}, RandomContent{
		Name:    "2",
		Content: "string 2",
		Rate:    200,
	}, RandomContent{
		Name:    "3",
		Content: "string 3",
		Rate:    300,
	})

	assert.Equal(t, 1235, randomBox.GetMaxRate(), "Correct max rate")
	assert.Equal(t, 3, len(randomBox.GetContents()), "Correct content size")
	assert.Equal(t, "[100 300 600]", fmt.Sprintf("%v", randomBox.GetAccRates()), "Correct content size")
}

// Full rate means box has no chance of missing in selecting
// go test -timeout 30s -run ^TestRandomBag_SelectingFullRate_NoNilResult$ github.com/zeroboo/randomselector -v
func TestRandomBag_SelectingFullRate_NoNilResult(t *testing.T) {
	var randomBox *RandomBag = CreateRandomBox(1000, true, RandomContent{
		Name:    "1",
		Content: "string 1",
		Rate:    100,
	}, RandomContent{
		Name:    "2",
		Content: "string 2",
		Rate:    200,
	}, RandomContent{
		Name:    "3",
		Content: "string 3",
		Rate:    300,
	}, RandomContent{
		Name:    "4",
		Content: "string 4",
		Rate:    400,
	})

	assert.Equal(t, 1000, randomBox.GetMaxRate(), "Correct max rate")
	assert.Equal(t, 4, len(randomBox.GetContents()), "Correct content size")
	assert.Equal(t, "[100 300 600 1000]", fmt.Sprintf("%v", randomBox.GetAccRates()), "Correct content size")
	for i := 0; i < 100; i++ {
		selectedValue := randomBox.SelectRandom()
		log.Printf("Selected values: %v", selectedValue)
		assert.NotEqual(t, nil, selectedValue, "Selected value must not be nil")
	}

}

type TestItem struct {
	ID    string
	Value int
	Rate  int
}

func (content TestItem) GetRate() int {
	return content.Rate
}

func (content TestItem) GetName() string {
	return content.ID
}

// TestRandomBag_SelectingFullRateWithStruct_NoNilResult: test box full rate with content is a struct
// go test -timeout 30s -run ^TestRandomBag_SelectingFullRateWithStruct_NoNilResult$ github.com/zeroboo/randomselector -v
func TestRandomBag_SelectingFullRateWithStruct_NoNilResult(t *testing.T) {
	var randomBox *RandomBag = CreateRandomBox(1000, true, RandomContent{
		Name: "1",
		Content: TestItem{
			ID:    "item1",
			Value: 1,
		},
		Rate: 100,
	}, RandomContent{
		Name: "2",
		Content: TestItem{
			ID:    "item2",
			Value: 2,
		},
		Rate: 200,
	}, RandomContent{
		Name: "3",
		Content: TestItem{
			ID:    "item3",
			Value: 3,
		},
		Rate: 300,
	}, RandomContent{
		Name: "4",
		Content: TestItem{
			ID:    "item4",
			Value: 4,
		},
		Rate: 400,
	})

	assert.Equal(t, 1000, randomBox.GetMaxRate(), "Correct max rate")
	assert.Equal(t, 4, len(randomBox.GetContents()), "Correct content size")
	assert.Equal(t, "[100 300 600 1000]", fmt.Sprintf("%v", randomBox.GetAccRates()), "Correct content size")
	for i := 0; i < 100; i++ {
		selectedValue := randomBox.SelectRandom()
		log.Printf("Selected values: %v", selectedValue)
		assert.NotEqual(t, nil, selectedValue, "Selected value must not be nil")
		assert.IsType(t, TestItem{}, selectedValue, "Correct selected type")
	}

}

// TestRandomBag_SelectedNil_Correct: test box full rate with content is a struct
// go test -timeout 30s -run ^TestRandomBag_SelectedNil_Correct$ github.com/zeroboo/randomselector -v
func TestRandomBag_SelectedNil_Correct(t *testing.T) {
	var randomBox *RandomBag = CreateRandomBox(1000, true, RandomContent{
		Name: "1",
		Content: TestItem{
			ID:    "item1",
			Value: 1,
		},
		Rate: 0,
	}, RandomContent{
		Name: "2",
		Content: TestItem{
			ID:    "item2",
			Value: 2,
		},
		Rate: 0,
	})

	assert.Equal(t, 1000, randomBox.GetMaxRate(), "Correct max rate")
	assert.Equal(t, 2, len(randomBox.GetContents()), "Correct content size")
	assert.Equal(t, "[0 0]", fmt.Sprintf("%v", randomBox.GetAccRates()), "Correct content size")
	for i := 0; i < 100; i++ {
		selectedValue := randomBox.SelectRandom()
		log.Printf("Selected values: %v", selectedValue)
		assert.Equal(t, nil, selectedValue, "Selected value must be nil")
	}
}

// go test -timeout 30s -run ^TestAddItemToBag_Correct$ github.com/zeroboo/randomselector -v
func TestAddItemToBag_Correct(t *testing.T) {
	item1 := TestItem{
		ID:    "item1",
		Value: 1,
		Rate:  1,
	}
	item2 := TestItem{
		ID:    "item2",
		Value: 2,
		Rate:  2,
	}
	bag := RandomBag{
		configMaxRate:       RandomRateNone,
		returnSelectedItems: true,
	}

	bag.AddItem(item1)
	bag.AddItem(item2)
	bag.AddItem(item2)

	log.Infof("Bag: %v", bag.String())
	assert.Equal(t, 5, bag.GetMaxRate(), "Correct max rates")
	assert.Equal(t, 3, len(bag.contents), "Correct item counts")
}
