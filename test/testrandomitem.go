package test

type TestRandomItem struct {
	ID    string
	Value int
	Rate  int64
}

func (content TestRandomItem) GetRate() int64 {
	return content.Rate
}

func (content TestRandomItem) GetName() string {
	return content.ID
}

func (content TestRandomItem) GetContent() any {
	return content.Value
}
