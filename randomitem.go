package randomselector

type RandomItemInterface interface {
	GetRate() int
	GetName() string
	GetContent() any
}
