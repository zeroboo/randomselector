package randombag

type IRandomItem interface {
	GetRate() int64
	GetName() string
	GetContent() any
}
