package server

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"kafji.net/buma/database"
	"kafji.net/buma/services"
)

var testUser = struct {
	email    string
	password string
}{
	"test@example.com",
	"hunter2",
}

var feedFetcher = services.FetchFeedFunc(
	func(ctx context.Context, url string) []services.FetchedFeedItem {
		testFeedSources := map[string][]services.FetchedFeedItem{
			"http://example.com": {
				services.FetchedFeedItem{Title: "Test Article", URL: "http://example.com/articles/1"},
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
		r := echo.New()
		SetupRouter(r, NewEnvFactory(db))

		createAccount(t, r)

		token := createAccessToken(t, r)

		addSource(t, r, token)

		services.FetchFeeds(ctx, db, feedFetcher, db)

		getSources(t, r, token)

		getFeed(t, r, token)
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
	req := httptest.NewRequest("POST", "/account", bytes.NewReader(body))
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
	req := httptest.NewRequest("POST", "/account/token", bytes.NewReader(body))
	req.Header.Add("content-type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "createAccessToken: received error response")

	res := struct {
		Token string `json:"token"`
	}{}
	err := json.Unmarshal(w.Body.Bytes(), &res)
	if !assert.Nil(t, err, "createAccessToken: failed to read response body") {
		t.Fail()
	}

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
	req := httptest.NewRequest("POST", "/user/source", bytes.NewReader(body))
	req.Header.Add("content-type", "application/json")
	req.Header.Add("authorization", fmt.Sprintf("Bearer %s", token))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "addSource: received error response")
}

func getSources(t *testing.T, r http.Handler, token string) {
	req := httptest.NewRequest("GET", "/user/sources", nil)
	req.Header.Add("content-type", "application/json")
	req.Header.Add("authorization", fmt.Sprintf("Bearer %s", token))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "getSources: received error response")

	res := struct {
		Sources []struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"sources"`
	}{}
	err := json.Unmarshal(w.Body.Bytes(), &res)
	if !assert.Nil(t, err, "getSources: failed to read response body") {
		t.Fail()
	}

	assert.Equal(t, 1, len(res.Sources))
	assert.Equal(t, "Test Blog", res.Sources[0].Name)
	assert.Equal(t, "http://example.com", res.Sources[0].URL)
}

func getFeed(t *testing.T, r http.Handler, token string) {
	req := httptest.NewRequest("GET", "/user/feed", nil)
	req.Header.Add("content-type", "application/json")
	req.Header.Add("authorization", fmt.Sprintf("Bearer %s", token))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "getFeed: received error response")

	res := struct {
		Feed struct {
			Items []struct {
				Title string `json:"title"`
				URL   string `json:"url"`
			} `json:"items"`
		} `json:"feed"`
	}{}
	err := json.Unmarshal(w.Body.Bytes(), &res)
	if !assert.Nil(t, err, "getFeed: failed to read response body") {
		t.Fail()
	}

	assert.Equal(t, 1, len(res.Feed.Items))
	assert.Equal(t, "Test Article", res.Feed.Items[0].Title)
	assert.Equal(t, "http://example.com/articles/1", res.Feed.Items[0].URL)
}
