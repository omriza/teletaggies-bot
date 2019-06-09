package tagging

import (
	"net/url"

	"github.com/pkg/errors"
)

type LinkTagger struct{}

func makeUrl(str string) (*url.URL, bool) {
	u, err := url.Parse(str)
	return u, err == nil && u.Scheme != "" && u.Host != ""
}

func (lt LinkTagger) Tag(tgbl Taggable) ([]Tag, error) {
	candidate := tgbl.(string)

	u, ok := makeUrl(candidate)
	if !ok {
		return nil, errors.New("Bad URL")
	}

	return extractTags(u), nil
}

func extractTags(u *url.URL) []Tag {
	tags := []Tag{Tag(u.Host)}

	return tags
}
