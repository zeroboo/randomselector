package randomselector

import (
	"math/rand"
	"time"
)

func init() {
	var seed int64 = time.Now().UnixNano()
	rand.Seed(seed)
}

const RandomRateNone int = -1

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
func NewRandomBag(maxRate int, returnSelectedItems bool, contents ...RandomContent) *RandomBag {
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

// NewRandomBagNoFailure return a random box which:
//
//   - All inside items have equal rates
//
//   - Every selecting returns an item (no chance of failure) if the bag has item
func NewRandomBagNoFailure(hasReplacement bool, contents ...RandomContent) *RandomBag {
	return NewRandomBag(RandomRateNone, hasReplacement, contents...)
}

func NewRandomBoxNoFailureFromItems(hasReplacement bool, items ...RandomItemInterface) *RandomBag {
	contents := make([]RandomContent, len(items))
	for i := 0; i < len(items); i++ {
		contents[i] = RandomContent{
			content: items[i],
			rate:    items[i].GetRate(),
		}
	}
	return NewRandomBag(RandomRateNone, hasReplacement, contents...)
}

func NewEmptyRandomBag(maxRate int, hasReplacement bool) *RandomBag {
	return NewRandomBag(maxRate, hasReplacement)
}
