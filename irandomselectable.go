package randomselector

// IRandomSelectable objects can return a randomly content inside
type IRandomSelectable interface {
	//SelectRandom Randomly returns an object. Nil is a possible result, it means selected nothing
	//
	//Rates is random in [0, GetMaxRate())
	SelectRandom() any

	//GetContents returns all possible option of box
	GetContents() []RandomContent

	//GetMaxRate return max values (exclusive in random rates)
	GetMaxRate() int
}
