package entity

type Article struct {
	name        string
	description string
	attributes  interface{}
}

func getNewarticle(name, description string, attributes interface{}) *Article {
	return &Article{name: name, description: description, attributes: attributes}
}

func (art *Article) Display() {
	//
}
