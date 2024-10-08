package randombag

import (
	"bytes"
	"fmt"
	"math/rand"

	log "github.com/sirupsen/logrus"
)

// RandomBag is a collection of items. It can return object randomly with replacement (each selecting is independent)
type RandomBag struct {
	contents []IRandomItem

	//configMaxRate is the preset max rate of this bag. Maybe not the effective rate used when random, see [RandomBag.maxRate]
	configMaxRate int64

	//maxRate is the effective maximum value of random (exclusive)
	maxRate int64

	//totalItemRates is the maximum rates that selecting has a valid result. It equals the sum of all items
	totalItemRates int64

	//accRates stores rates of [contents] in continuous list. [accRates] is correspondence to [contents],
	//It means accRates[i] stands for contents[i]
	//It used for randomly pick the index of content: accRates[i+1]-accRates[i] is the rate of content i
	//
	//Eg: a box has 3 contents, each has rate of 10, the accRates will be [10 20 30]
	//  - with random rate 1->9, selected item is content[0]
	//  - with random rate 10->19, selected item is content[1]
	//  - with random rate is 20->29, selected item is content[2]
	//  - with random rate is >=30, selected item is nil
	accRates []int64

	//returnSelectedItems has replacement or not:
	//  - if true, items after selecting will not be changed
	//  - if false, item selected will be removed from list
	returnSelectedItems bool
}

// Select returns an object randomly with replacement.
// Nil result means nothing selected
func (bag *RandomBag) Select() (any, error) {
	if bag.accRates == nil {
		bag.initRates()
	}

	if bag.maxRate <= 0 {
		return nil, fmt.Errorf("invalid max rate %v", bag.maxRate)
	}
	rate := rand.Int63n(bag.maxRate)
	for i := 0; i < len(bag.accRates); i++ {
		if rate < bag.accRates[i] {
			return bag.contents[i].GetContent(), nil
		}
	}
	return nil, nil
}

// GetContents returns all possible option of box
func (bag *RandomBag) GetContents() []IRandomItem {
	return bag.contents
}

// GetMaxRate return max values (exclusive in random rates)
func (bag *RandomBag) GetMaxRate() int64 {
	return bag.maxRate
}

// GetConfigMaxRate return max values configured
func (bag *RandomBag) GetConfigMaxRate() int64 {
	return bag.configMaxRate
}

// GetTotalItemRates return sum of all item rates
func (bag *RandomBag) GetTotalItemRates() int64 {
	return bag.totalItemRates
}

// initRates prepares cached values for picking
func (bag *RandomBag) initRates() int64 {
	bag.accRates = make([]int64, len(bag.contents))
	totalAccRate := int64(0)
	for i := 0; i < len(bag.contents); i++ {
		totalAccRate += bag.contents[i].GetRate()
		bag.accRates[i] = totalAccRate
	}
	bag.updateMaxRates()
	return totalAccRate
}

func (bag *RandomBag) updateMaxRates() {
	if bag.configMaxRate > RandomRateNone {
		bag.maxRate = bag.configMaxRate
	} else {
		bag.maxRate = bag.totalItemRates
	}
}

// GetAccRates return accRates
func (bag *RandomBag) GetAccRates() []int64 {
	return bag.accRates
}

func (bag *RandomBag) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("RandomBag")
	buffer.WriteString("|ReturnSelectedItems:")
	buffer.WriteString(fmt.Sprintf("%v", bag.returnSelectedItems))
	buffer.WriteString("|PresetMaxRate:")
	buffer.WriteString(fmt.Sprintf("%v", bag.configMaxRate))
	buffer.WriteString("|CurrentMaxRate:")
	buffer.WriteString(fmt.Sprintf("%v", bag.maxRate))
	buffer.WriteString("|Contents:")
	buffer.WriteString(fmt.Sprintf("%v", bag.contents))
	return buffer.String()
}

func (bag *RandomBag) AddItem(item IRandomItem) {
	newContent := RandomItem{
		content: item,
		rate:    item.GetRate(),
		name:    item.GetName(),
	}
	bag.contents = append(bag.contents, newContent)
	bag.totalItemRates = bag.initRates()

	bag.updateMaxRates()

	if log.IsLevelEnabled(log.TraceLevel) {
		log.Tracef("RandomBox.AddItem: newItem=%v, rate=%v, maxRate=%v, items=%v, totalItemRate=%v", item.GetName(), item.GetRate(), bag.GetMaxRate(), len(bag.contents), bag.totalItemRates)
	}
}

const RandomRateNone int64 = -1

// NewRandomBag returns a random bag with config:
//
// - maxRate: maximum value of random rate. If maxRate <= RandomRateNone, maxRate will be calculated as the sum of rate of all items
//
// - replacement: picked items have chance to appear in next random or not
//
//   - true means picked items will have chance to appear in next random
//
//   - false means picked items will be removed from next random
//
// - contents: are items to be randomized in random bag
func NewRandomBag(maxRate int64, returnSelectedItems bool, contents ...IRandomItem) *RandomBag {
	var randomBag *RandomBag = &RandomBag{}
	randomBag.contents = contents
	randomBag.totalItemRates = randomBag.initRates()
	randomBag.configMaxRate = maxRate

	if randomBag.configMaxRate > RandomRateNone {
		randomBag.maxRate = randomBag.configMaxRate
	} else {
		randomBag.maxRate = randomBag.totalItemRates
	}
	randomBag.returnSelectedItems = returnSelectedItems

	return randomBag
}

func NewRandomBagNoFailure(hasReplacement bool, items ...IRandomItem) *RandomBag {
	contents := make([]IRandomItem, len(items))
	for i := 0; i < len(items); i++ {
		contents[i] = RandomItem{
			content: items[i],
			rate:    items[i].GetRate(),
		}
	}
	return NewRandomBag(RandomRateNone, hasReplacement, contents...)
}

func NewEmptyRandomBag(maxRate int64, hasReplacement bool) *RandomBag {
	return NewRandomBag(maxRate, hasReplacement)
}
