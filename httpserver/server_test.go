package httpserver

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"kafji.net/buma/database"
	fetchfeeds "kafji.net/buma/services/fetch_feeds"
)

var testUser = struct {
	email    string
	password string
}{
	"test@example.com",
	"hunter2",
}

var feedFetcher = fetchfeeds.FetchFeedFunc(
	func(ctx context.Context, url string) []fetchfeeds.FetchedFeedItem {
		testFeedSources := map[string][]fetchfeeds.FetchedFeedItem{
			"http://example.com": {
				fetchfeeds.FetchedFeedItem{Title: "Test Article", URL: "http://example.com/articles/1"},
			},
		}
		if xs, ok := testFeedSources[url]; ok {
			return xs
		}
		return nil
	})

func TestServer(t *testing.T) {
	ctx := context.Background()

	database.WithDatabase(ctx, t, func(db database.Database) {
		r := chi.NewRouter()
		setupRouter(r, db)

		createAccount(t, r)

		token := createAccessToken(t, r)

		addSource(t, r, token)

		getUserSources(t, r, token)

		fetchfeeds.FetchFeeds(ctx, db, feedFetcher, db)

		getUserFeed(t, r, token)
	})
}

func createAccount(t *testing.T, r http.Handler) {
	body, _ := json.Marshal(&struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{
		testUser.email,
		testUser.password,
	})
	req := httptest.NewRequest("POST", "/signup", bytes.NewReader(body))
	req.Header.Add("content-type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "createAccount: received error response")
}

func createAccessToken(t *testing.T, r http.Handler) string {
	body, _ := json.Marshal(&struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{
		testUser.email,
		testUser.password,
	})
	req := httptest.NewRequest("POST", "/login", bytes.NewReader(body))
	req.Header.Add("content-type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "createAccessToken: received error response")

	res := struct {
		Token string `json:"token"`
	}{}
	err := json.Unmarshal(w.Body.Bytes(), &res)
	assert.Nil(t, err, "createAccessToken: failed to read response body")
	return res.Token
}

func addSource(t *testing.T, r http.Handler, token string) {
	body, _ := json.Marshal(&struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}{
		"Test Blog",
		"http://example.com",
	})
	req := httptest.NewRequest("POST", "/me/source", bytes.NewReader(body))
	req.Header.Add("content-type", "application/json")
	req.Header.Add("authorization", fmt.Sprintf("Bearer %s", token))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "addSource: received error response")
}

func getUserSources(t *testing.T, r http.Handler, token string) {
	req := httptest.NewRequest("GET", "/me/sources", nil)
	req.Header.Add("content-type", "application/json")
	req.Header.Add("authorization", fmt.Sprintf("Bearer %s", token))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "getUserSources: received error response")

	res := struct {
		Sources []struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"sources"`
	}{}
	err := json.Unmarshal(w.Body.Bytes(), &res)
	assert.Nil(t, err, "getUserSources: failed to read response body")
	assert.Equal(t, 1, len(res.Sources))
	assert.Equal(t, "Test Blog", res.Sources[0].Name)
	assert.Equal(t, "http://example.com", res.Sources[0].URL)
}

func getUserFeed(t *testing.T, r http.Handler, token string) {
	req := httptest.NewRequest("GET", "/me/feed", nil)
	req.Header.Add("content-type", "application/json")
	req.Header.Add("authorization", fmt.Sprintf("Bearer %s", token))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "getUserFeed: received error response")

	res := struct {
		Items []struct {
			Title string `json:"title"`
			URL   string `json:"url"`
		} `json:"items"`
	}{}
	err := json.Unmarshal(w.Body.Bytes(), &res)
	assert.Nil(t, err, "getUserFeed: failed to read response body")
	assert.Equal(t, 1, len(res.Items), "getUserFeed: expecting one item")
	assert.Equal(t, "Test Article", res.Items[0].Title)
	assert.Equal(t, "http://example.com/articles/1", res.Items[0].URL)
}
