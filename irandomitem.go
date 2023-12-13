package randomselector

type IRandomItem interface {
	GetRate() int
	GetName() string
	GetContent() any
}
