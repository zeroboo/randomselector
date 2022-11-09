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

// CreateRandomBox return random bags with config:
//   - maxRate: maximum value of random rate. If maxRate <= RandomRateNone, maxRate will be calculated as the sum of rate of all items
func CreateRandomBox(maxRate int, returnSelectedItems bool, contents ...RandomContent) *RandomBag {
	var randomBag *RandomBag = &RandomBag{}
	randomBag.contents = contents
	randomBag.totalItemRates = randomBag.initRates()
	if maxRate > RandomRateNone {
		randomBag.maxRate = maxRate
	} else {
		randomBag.maxRate = randomBag.totalItemRates
	}
	randomBag.returnSelectedItems = returnSelectedItems
	return randomBag
}

func CreateRandomBoxNoFailure(hasReplacement bool, contents ...RandomContent) *RandomBag {
	return CreateRandomBox(RandomRateNone, hasReplacement, contents...)
}

func CreateRandomBoxNoFailureFromItems(hasReplacement bool, items ...RandomItemInterface) *RandomBag {
	contents := make([]RandomContent, len(items))
	for i := 0; i < len(items); i++ {
		contents[i] = RandomContent{
			Content: items[i],
			Rate:    items[i].GetRate(),
		}
	}
	return CreateRandomBox(RandomRateNone, hasReplacement, contents...)
}
