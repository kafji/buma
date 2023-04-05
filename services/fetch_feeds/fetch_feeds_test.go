package fetchfeeds

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

type fakeGetFeedSources struct {
	urls []string
}

func newFakeGetFeedSources(urls ...string) fakeGetFeedSources {
	return fakeGetFeedSources{urls}
}

func (s fakeGetFeedSources) GetFeedSources(ctx context.Context) []string {
	return s.urls
}

type fakeFetchFeed struct {
	feeds map[string][]FetchedFeedItem
}

func newFakeFetchFeed() fakeFetchFeed {
	return fakeFetchFeed{map[string][]FetchedFeedItem{
		"http://example.com": {
			{Title: "Test Article", URL: "http://example.com/articles/1"},
		},
	}}
}

func (s fakeFetchFeed) FetchFeed(ctx context.Context, url string) []FetchedFeedItem {
	return s.feeds[url]
}

type fakePutFeed struct {
	items []StorableFeedItem
}

func newFakePutFeed(items ...StorableFeedItem) fakePutFeed {
	return fakePutFeed{items}
}

func (s *fakePutFeed) PutFeed(ctx context.Context, items []StorableFeedItem) {
	s.items = append(s.items, items...)
}

func TestFetchFeeds(t *testing.T) {
	gfs := newFakeGetFeedSources("http://example.com")
	ff := newFakeFetchFeed()
	pf := newFakePutFeed()

	FetchFeeds(context.Background(), gfs, ff, &pf)

	assert.Equal(t, 1, len(pf.items))
	assert.Equal(t, "Test Article", pf.items[0].Title)
	assert.Equal(t, "http://example.com/articles/1", pf.items[0].URL)
	assert.Equal(t, "http://example.com", pf.items[0].SourceURL)
}
