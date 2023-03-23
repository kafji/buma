package services

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

type fakeAddFeedSource struct {
	sources []struct {
		userID     int
		sourceName string
		sourceURL  string
	}
}

func newFakeAddFeedSource() fakeAddFeedSource {
	return fakeAddFeedSource{}
}

func (s *fakeAddFeedSource) AddFeedSource(
	ctx context.Context,
	userID int,
	sourceName,
	sourceURL string,
) (sourceID int, ok bool) {
	exists := func() bool {
		for _, x := range s.sources {
			if x.userID == userID && x.sourceName == sourceName {
				return true
			}
		}
		return false
	}

	if exists() {
		ok = false
		return
	}

	s.sources = append(s.sources, struct {
		userID     int
		sourceName string
		sourceURL  string
	}{userID, sourceName, sourceURL})

	sourceID = len(s.sources)
	ok = true

	return
}

func TestAddSource(t *testing.T) {
	afs := newFakeAddFeedSource()

	sourceID, err := AddSource(context.Background(), &afs, 1, "Test Source", "http://example.com")

	assert.Nil(t, err)
	assert.Equal(t, 1, sourceID)
}

func TestAddSourceNonUniqueName(t *testing.T) {
	afs := newFakeAddFeedSource()

	_, err := AddSource(context.Background(), &afs, 1, "Test Source", "http://example.com")
	assert.Nil(t, err)

	_, err = AddSource(context.Background(), &afs, 1, "Test Source", "http://example.net")

	assert.Equal(t, ErrSourceAlreadyExists, err)
}

func TestAddSourceEmptyName(t *testing.T) {
	afs := newFakeAddFeedSource()

	_, err := AddSource(context.Background(), &afs, 1, "", "http://example.com")

	assert.Equal(t, ErrEmptySourceName, err)
}

func TestAddSourceEmptyURL(t *testing.T) {
	afs := newFakeAddFeedSource()

	_, err := AddSource(context.Background(), &afs, 1, "Test Source", "")

	assert.Equal(t, ErrEmptySourceURL, err)
}
