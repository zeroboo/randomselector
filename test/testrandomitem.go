package test

type TestRandomItem struct {
	ID    string
	Value int
	Rate  int
}

func (content TestRandomItem) GetRate() int {
	return content.Rate
}

func (content TestRandomItem) GetName() string {
	return content.ID
}

func (content TestRandomItem) GetContent() any {
	return content.ID
}
