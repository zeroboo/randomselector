package randomselector

type IRandomitem interface {
	GetRate() int
	GetName() string
	GetContent() any
}
