package searcher

import (
	"github.com/stretchr/testify/assert"
	"google.golang.org/api/customsearch/v1"
	"testing"
)

func TestSearchResultToString(t *testing.T) {
	type args struct {
		s *customsearch.Search
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Positive case",
			args: args{s: &customsearch.Search{Items: []*customsearch.Result{
				{
					Link:    "link1",
					Snippet: "snippet1",
				},
				{
					Link:    "link2",
					Snippet: "snippet2",
				},
			}}},
			want: `1. snippet1 (source: link1)
2. snippet2 (source: link2)
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SearchResultToString(tt.args.s)
			assert.Equal(t, tt.want, got)
		})
	}
}
