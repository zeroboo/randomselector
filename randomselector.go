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
