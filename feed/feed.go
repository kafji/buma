package feed

import (
	"context"

	"github.com/mmcdole/gofeed"
	"golang.org/x/exp/slog"
	fetchfeeds "kafji.net/buma/services/fetch_feeds"
)

func FetchFeeds(ctx context.Context, url string) []fetchfeeds.FetchedFeedItem {
	logger := slog.With("tag", "feed/fetch")
	logger.Info("fetching feed", "url", url)

	parser := gofeed.NewParser()
	feed, err := parser.ParseURLWithContext(url, ctx)
	if err != nil {
		logger.Error("failed to fetch feed items", err, url)
		return nil
	}

	items := make([]fetchfeeds.FetchedFeedItem, 0, len(feed.Items))
	for _, item := range feed.Items {
		items = append(items, fetchfeeds.FetchedFeedItem{Title: item.Title, URL: item.Link})
	}

	return items
}
