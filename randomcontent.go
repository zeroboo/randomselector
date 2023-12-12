package randomselector

// RandomContent contains target of random picking and rates
type RandomContent struct {
	//content refers to the object to be randomly selected
	content any

	//rate tells how often this content will appear
	//
	//Zero means this content will never appear. Negative values are ERRONEOUS
	rate int

	//name is name of content
	name string
}

func (content RandomContent) GetRate() int {
	return content.rate
}

func (content RandomContent) GetName() string {
	return content.name
}

func (content RandomContent) GetContent() any {
	return content.content
}

func NewRandomContent(name string, rate int, content any) *RandomContent {
	return &RandomContent{
		content: content,
		rate:    rate,
		name:    name,
	}

}
