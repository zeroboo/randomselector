package randombag

import "fmt"

// RandomItem contains target of random picking and rates
type RandomItem struct {
	//content refers to the object to be randomly selected
	content any

	//rate tells how often this content will appear
	//
	//Zero means this content will never appear. Negative values are ERRONEOUS
	rate int64

	//name is name of content
	name string
}

func (content RandomItem) GetRate() int64 {
	return content.rate
}

func (content RandomItem) GetName() string {
	return content.name
}

func (content RandomItem) GetContent() any {
	return content.content
}

func NewRandomItem(name string, rate int64, content any) *RandomItem {
	return &RandomItem{
		content: content,
		rate:    rate,
		name:    name,
	}
}

func (content RandomItem) String() string {
	return fmt.Sprintf("RandomContent: %v,%v,%v", content.name, content.rate, content.content)
}
