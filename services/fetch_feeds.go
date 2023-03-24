package services

import (
	"context"
	"math"
	"runtime"
)

type GetFeedSources interface {
	GetFeedSources(ctx context.Context) []string
}

type FetchFeed interface {
	FetchFeed(ctx context.Context, url string) []FetchedFeedItem
}

type FetchedFeedItem struct {
	Title string
	URL   string
}

type PutFeed interface {
	PutFeed(ctx context.Context, items []StorableFeedItem)
}

type StorableFeedItem struct {
	FetchedFeedItem
	SourceURL string
}

func fetchFeeds(ctx context.Context, ff FetchFeed, pfi PutFeed, urls []string) {
	for _, url := range urls {
		fetched := ff.FetchFeed(ctx, url)

		items := []StorableFeedItem{}
		for _, item := range fetched {
			items = append(items, StorableFeedItem{FetchedFeedItem: item, SourceURL: url})
		}
		pfi.PutFeed(ctx, items)
	}
}

func FetchFeeds(ctx context.Context, gs GetFeedSources, ff FetchFeed, pf PutFeed) {
	ss := gs.GetFeedSources(ctx)
	sslen := len(ss)
	maxpar := runtime.GOMAXPROCS(0)

	chunkSize := max(int(math.Ceil(float64(sslen)/float64(maxpar))), 1)

	cs := []chan struct{}{}

	for i := 0; i < sslen; i += chunkSize {
		c := make(chan struct{})
		cs = append(cs, c)

		i := i
		j := min(i+chunkSize, sslen)

		go func() {
			fetchFeeds(ctx, ff, pf, ss[i:j])
			c <- struct{}{}
		}()
	}

	for _, c := range cs {
		<-c
	}
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

type FetchFeedFunc func(context.Context, string) []FetchedFeedItem

func (s FetchFeedFunc) FetchFeed(ctx context.Context, url string) []FetchedFeedItem {
	return s(ctx, url)
}
