package randomselector

import (
	"math/rand"
	"time"
)

func init() {
	var seed int64 = time.Now().UnixNano()
	rand.Seed(seed)
}
func CreateRandomBoxNoReplacement(maxRate int, contents ...RandomContent) *RandomBag {
	var randomBag *RandomBag = &RandomBag{}
	randomBag.contents = contents
	randomBag.maxRate = maxRate
	randomBag.initRates()

	return randomBag
}
