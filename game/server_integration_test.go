package poker_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	poker "github.com/biggsean/learn-go-with-tests/game"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	database, cleanDatabase := poker.CreateTempFile(t, "[]")
	defer cleanDatabase()
	store, err := poker.NewFileSystemPlayerStore(database)
	poker.AssertNoError(t, err)
	server := poker.NewPlayerServer(store)
	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), poker.NewPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), poker.NewPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), poker.NewPostWinRequest(player))

	t.Run("get score", func(t *testing.T) {
		res := httptest.NewRecorder()
		server.ServeHTTP(res, poker.NewGetScoreRequest(player))
		poker.AssertStatus(t, res.Code, http.StatusOK)

		poker.AssertResponseBody(t, res.Body.String(), "3")
	})
	t.Run("get league", func(t *testing.T) {
		res := httptest.NewRecorder()
		server.ServeHTTP(res, poker.NewGetLeagueRequest())
		poker.AssertStatus(t, res.Code, http.StatusOK)

		got := poker.GetLeagueFromResponse(t, res.Body)
		want := poker.League{
			{"Pepper", 3},
		}
		poker.AssertLeague(t, got, want)
	})
}
