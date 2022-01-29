package earticle

import (
	"fmt"
)

type Article struct {
	name,
	description string
}

func NewArticle(name, description string) *Article {
	return &Article{name, description}
}

func (article *Article) String() string {
	return fmt.Sprintf("%s (%s)", article.name, article.description)
}
