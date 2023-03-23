package feed

import (
	"context"

	"github.com/mmcdole/gofeed"
	"golang.org/x/exp/slog"
	"kafji.net/buma/services"
)

func FetchFeeds(ctx context.Context, url string) []services.FetchedFeedItem {
	parser := gofeed.NewParser()
	feed, err := parser.ParseURLWithContext(url, ctx)
	if err != nil {
		slog.Error("failed to fetch feed items", err, url)
		return nil
	}

	items := []services.FetchedFeedItem{}
	for _, item := range feed.Items {
		items = append(items, services.FetchedFeedItem{Title: item.Title, URL: item.Link})
	}

	return items
}
