package entities

import "fmt"

type Article struct {
	name,
	description string
}

func (article *Article) String() string {
	return fmt.Sprintf("%s (%s)", article.name, article.description)
}
