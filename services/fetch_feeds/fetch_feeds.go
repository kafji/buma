package fetchfeeds

import (
	"context"
	"fmt"
	"math"
	"runtime"

	"golang.org/x/exp/slog"
)

type QueryFeedSources interface {
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

func fetchFeeds(ctx context.Context, ff FetchFeed, pf PutFeed, urls []string) {
	for _, url := range urls {
		fetched := ff.FetchFeed(ctx, url)

		items := make([]StorableFeedItem, 0, len(fetched))
		for _, item := range fetched {
			items = append(items, StorableFeedItem{FetchedFeedItem: item, SourceURL: url})
		}

		slog.Info(fmt.Sprintf("fetched %d items", len(items)), "url", url)

		pf.PutFeed(ctx, items)
	}
}

func FetchFeeds(ctx context.Context, qfs QueryFeedSources, ff FetchFeed, pf PutFeed) {
	srcs := qfs.GetFeedSources(ctx)

	slog.Info(fmt.Sprintf("fetching feeds from %d sources", len(srcs)))

	srcsLen := len(srcs)
	maxPar := runtime.GOMAXPROCS(0)

	chunkSize := max(int(math.Ceil(float64(srcsLen)/float64(maxPar))), 1)

	cs := []chan struct{}{}

	for i := 0; i < srcsLen; i += chunkSize {
		c := make(chan struct{})
		cs = append(cs, c)

		i := i
		j := min(i+chunkSize, srcsLen)

		go func() {
			fetchFeeds(ctx, ff, pf, srcs[i:j])
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
