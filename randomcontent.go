package randomselector

// RandomContent contains target of random picking and rates
type RandomContent struct {
	//Content is object to be randomly selected
	Content any

	//Rate tells how often this content will appear
	//
	//Zero means this content will never appear. Negative values are ERRONEOUS
	Rate int

	//Name is name of content
	Name string
}

func (content RandomContent) GetRate() int {
	return content.Rate
}

func (content RandomContent) GetName() string {
	return content.Name
}

func NewRandomContent(name string, rate int, content any) *RandomContent {
	return &RandomContent{
		Content: content,
		Rate:    rate,
		Name:    name,
	}

}
