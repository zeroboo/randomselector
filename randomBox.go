package randomselector

import (
	"math/rand"
)

// RandomContent contains target of random picking and rates
type RandomContent struct {
	//Name is name of content
	Name string

	//Content is object to be randomly selected
	Content any

	//Rate tells how often this content will appear
	//
	//Zero means this content will never appear. Negative values are ERRONEOUS
	Rate int
}

// RandomSelectable can return a randomly object inside
type RandomSelectable interface {
	//SelectRandom Randomly returns an object. Nil is a possible result, it means selected nothing
	//
	//Rates is random in [0, GetMaxRate())
	SelectRandom() any

	//GetContents returns all possible option of box
	GetContents() []RandomContent

	//GetMaxRate return max values (exclusive in random rates)
	GetMaxRate() int
}

// RandomBag return object randomly with replacement (each selecting is independent)
type RandomBag struct {
	contents []RandomContent
	//maxRate is the maximum value of random (exclusive)
	maxRate int

	//maxRateHasValidItem is the maximum rates that selecting has a valid result
	maxRateHasValidItem int

	//accRates stores rates of [contents] in continuous list. [accRates] is correspondence to [contents],
	//It means accRates[i] stands for contents[i]
	//It used for randomly pick the index of content: accRates[i+1]-accRates[i] is the rate of content i
	//
	//Eg: a box has 3 contents, each has rate of 10, the accRates will be [10 20 30]
	//  - with random rate 1->9, selected item is content[0]
	//  - with random rate 10->19, selected item is content[1]
	//  - with random rate is 20->29, selected item is content[2]
	//  - with random rate is >=30, selected item is nil
	accRates []int
}

// Select returns an object randomly with replacement.
// Nil result means nothing selected
func (bag *RandomBag) SelectRandom() any {
	if bag.accRates == nil {
		bag.initRates()
	}

	rate := rand.Intn(bag.maxRate)
	for i := 0; i < len(bag.accRates); i++ {
		if rate < bag.accRates[i] {
			return bag.contents[i].Content
		}
	}
	return nil
}

// GetContents returns all possible option of box
func (bag *RandomBag) GetContents() []RandomContent {
	return bag.contents
}

// GetMaxRate return max values (exclusive in random rates)
func (bag *RandomBag) GetMaxRate() int {
	return bag.maxRate
}

// initRates prepares cached values for picking
func (bag *RandomBag) initRates() {
	bag.accRates = make([]int, len(bag.contents))
	accRate := 0
	for i := 0; i < len(bag.contents); i++ {
		accRate += bag.contents[i].Rate
		bag.accRates[i] = accRate
	}
	bag.maxRateHasValidItem = accRate
}

// GetAccRates return accRates
func (bag *RandomBag) GetAccRates() []int {
	return bag.accRates
}
