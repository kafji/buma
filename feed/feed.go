package feed

import (
	"context"

	"github.com/mmcdole/gofeed"
	"github.com/rs/zerolog/log"
	"kafji.net/buma/services"
)

func FetchFeeds(ctx context.Context, url string) []services.FetchedFeedItem {
	parser := gofeed.NewParser()
	feed, err := parser.ParseURLWithContext(url, ctx)
	if err != nil {
		log.Error().Str("tag", "feed").Err(err).Str("url", url).Msg("failed to fetch feed items")
		return nil
	}

	items := []services.FetchedFeedItem{}
	for _, item := range feed.Items {
		items = append(items, services.FetchedFeedItem{Title: item.Title, URL: item.Link})
	}

	return items
}
