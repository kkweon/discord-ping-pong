package searcher

import (
	"context"
	"fmt"
	"google.golang.org/api/customsearch/v1"
	"google.golang.org/api/option"
	"strings"
	"time"
)

type Searcher struct {
	APIKey string
}

func New(apiKey string) Searcher {
	return Searcher{APIKey: apiKey}
}

func (s *Searcher) Search(term string) (*customsearch.Search, error) {
	ctx, _ := context.WithTimeout(context.Background(), time.Minute)

	svc, err := customsearch.NewService(ctx, option.WithAPIKey(s.APIKey))
	if err != nil {
		return nil, err
	}

	ret, err := svc.Cse.List(term).Num(5).Cx("20ea5fad086b46abc").Do()
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func SearchResultToString(s *customsearch.Search) string {
	b := strings.Builder{}
	for i, item := range s.Items {
		b.WriteString(fmt.Sprintf("%d. %s (source: %s)\n", i+1, item.Snippet, item.Link))
	}
	return b.String()
}
