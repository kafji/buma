package addsource

import "context"

type AddSourceError string

func (s AddSourceError) Error() string {
	return string(s)
}

const (
	ErrEmptySourceName     = AddSourceError("source name must not be empty")
	ErrEmptySourceURL      = AddSourceError("source url must not be empty")
	ErrSourceAlreadyExists = AddSourceError("source name already exists")
)

type AddFeedSource interface {
	AddFeedSource(ctx context.Context, userID int, sourceName, sourceURL string) (sourceID int, ok bool)
}

func AddSource(
	ctx context.Context,
	afs AddFeedSource,
	userID int,
	sourceName,
	sourceURL string,
) (sourceID int, err error) {
	if sourceName == "" {
		err = ErrEmptySourceName
		return
	}

	if sourceURL == "" {
		err = ErrEmptySourceURL
		return
	}

	sourceID, ok := afs.AddFeedSource(ctx, userID, sourceName, sourceURL)
	if !ok {
		err = ErrSourceAlreadyExists
		return
	}

	return
}
