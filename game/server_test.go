package poker

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETPlayers(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
		nil,
		nil,
	}
	server := NewPlayerServer(&store)

	t.Run("returns Pepper's score", func(t *testing.T) {
		req := NewGetScoreRequest("Pepper")
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		AssertResponseBody(t, res.Body.String(), "20")
		AssertStatus(t, res.Code, http.StatusOK)
	})
	t.Run("return Floyd's score", func(t *testing.T) {
		req := NewGetScoreRequest("Floyd")
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		AssertResponseBody(t, res.Body.String(), "10")
		AssertStatus(t, res.Code, http.StatusOK)
	})
	t.Run("return 404 on missing players", func(t *testing.T) {
		req := NewGetScoreRequest("Apollo")
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		AssertStatus(t, res.Code, http.StatusNotFound)
	})
}

func TestStoreWins(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{},
		nil,
		nil,
	}
	server := NewPlayerServer(&store)

	t.Run("it records wins on POST", func(t *testing.T) {
		player := "Pepper"
		req := NewPostWinRequest(player)
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		AssertStatus(t, res.Code, http.StatusAccepted)
		AssertPlayerWin(t, &store, player)
	})
}

func TestLeague(t *testing.T) {
	t.Run("it returns the league table as JSON", func(t *testing.T) {
		wantedLeague := []Player{
			{"Cleo", 32},
			{"Chris", 20},
			{"Tiest", 14},
		}

		store := StubPlayerStore{nil, nil, wantedLeague}
		server := NewPlayerServer(&store)

		req := NewGetLeagueRequest()
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		got := GetLeagueFromResponse(t, res.Body)
		AssertStatus(t, res.Code, http.StatusOK)
		AssertLeague(t, got, wantedLeague)
		AssertContentType(t, res, jsonContentType)
	})
}
