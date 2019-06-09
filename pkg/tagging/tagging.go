package tagging

type Tag string

type Taggable interface{}

type Tagger interface {
	Tag(src Taggable) []Tag
}
