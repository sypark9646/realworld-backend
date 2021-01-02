package domain

import (
	"time"
)

type article struct {
	slug        string
	title       string
	description string
	body        string
	tagList     []string
	createdAt   time.Time
	updatedAt   time.Time
	favoritedBy []user
	author      user
	comments    []comment
}

type comment struct {
	id        int
	createdAt time.Time
	updatedAt time.Time
	body      string
	author    user
}

type articleUpdatableField int

const (
	Title articleUpdatableField = iota
	Description
	Body
)

func UpdateArticle(initial *article, opts ...func(fields *article)) {
	for _, v := range opts {
		v(initial)
	}
}

func SetArticleTitle(input *string) func(fields *article) {
	return func(initial *article) {
		if input != nil {
			initial.title = *input
		}
	}
}

func SetArticleDescription(input *string) func(fields *article) {
	return func(initial *article) {
		if input != nil {
			initial.description = *input
		}
	}
}

func SetArticleBody(input *string) func(fields *article) {
	return func(initial *article) {
		if input != nil {
			initial.body = *input
		}
	}
}

type ArticleFilter func(article) bool

func ArticleHasTag(tag string) ArticleFilter {
	return func(article article) bool {
		for _, articleTag := range article.tagList {
			if articleTag == tag {
				return true
			}
		}
		return false
	}
}

func ArticleHasAuthor(authorName string) ArticleFilter {
	return func(article article) bool {
		return article.author.name == authorName
	}
}

func ArticleIsFavoritedBy(username string) ArticleFilter {
	return func(article article) bool {
		if username == "" {
			return false
		}
		for _, user := range article.favoritedBy {
			if user.name == username {
				return true
			}
		}
		return false
	}
}

type ArticleCollection []article

func (articles ArticleCollection) ApplyLimitAndOffset(limit, offset int) ArticleCollection {
	if limit <= 0 {
		return []article{}
	}

	articlesSize := len(articles)
	min := offset
	if min < 0 {
		min = 0
	}

	if min > articlesSize {
		return []article{}
	}

	max := min + limit
	if max > articlesSize {
		max = articlesSize
	}

	return articles[min:max]
}

func (article *article) UpdateComments(comment comment, add bool) {
	if add {
		article.comments = append(article.comments, comment)
		return
	}

	for i := 0; i < len(article.comments); i++ {
		if article.comments[i].id == comment.id {
			article.comments = append(article.comments[:i], article.comments[i+1:]...) // memory leak ? https://github.com/golang/go/wiki/SliceTricks
		}
	}
}

func (article *article) UpdateFavoritedBy(user user, add bool) {
	if add {
		article.favoritedBy = append(article.favoritedBy, user)
		return
	}

	for i := 0; i < len(article.favoritedBy); i++ {
		if article.favoritedBy[i].name == user.name {
			article.favoritedBy = append(article.favoritedBy[:i], article.favoritedBy[i+1:]...) // memory leak ? https://github.com/golang/go/wiki/SliceTricks
		}
	}
}
