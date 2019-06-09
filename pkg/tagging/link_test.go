package tagging

import (
	"reflect"
	"testing"
)

func TestLinkTagger_Tag(t *testing.T) {
	tests := []struct {
		name     string
		taggable Taggable
		want     []Tag
		wantErr  bool
	}{
		{"Scheme only", "https", nil, true},
		{"Scheme with slashes", "http://", nil, true},
		{"No scheme", "www.google.com", nil, true},
		{"Path", "/some/kind/of/magic", nil, true},
		{"Generic string", "try-this", nil, true},
		{"Scheme with service name", "http://some-svc", []Tag{"some-svc"}, false},
		{"Partial domain", "http://www.google", []Tag{"www.google"}, false},
		{"Normal url", "http://www.google.com", []Tag{"www.google.com"}, false},
		{"Normal url with fragment", "http://www.google.com/page.htm#kickass", []Tag{"www.google.com"}, false},
		{"Normal url with port", "http://www.google.com:8080", []Tag{"www.google.com:8080"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lt := LinkTagger{}
			got, err := lt.Tag(tt.taggable)
			if (err != nil) != tt.wantErr {
				t.Errorf("LinkTagger.Tag() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LinkTagger.Tag() = %v, want %v", got, tt.want)
			}
		})
	}
}
