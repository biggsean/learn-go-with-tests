package poker

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"
)

// StubPlayerStore mocks a PlayerStore
type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
	league   League
}

// GetPlayerScore gets the player's score
func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

// RecordWin records a win for a player in the mock
func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

// GetLeague gets the slice of players and their wins
func (s *StubPlayerStore) GetLeague() League {
	return s.league
}

// NewGetLeagueRequest creates a get request at /league
func NewGetLeagueRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/league", nil)
	return req
}

// NewGetScoreRequest creates a request for the wins of a provided user
func NewGetScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return req
}

// NewPostWinRequest POSTS a win for a provided player
func NewPostWinRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
	return req
}

// GetLeagueFromResponse translates the json reponse in to type League
func GetLeagueFromResponse(t *testing.T, body io.Reader) []Player {
	t.Helper()
	league, _ := NewLeague(body)
	return league
}

// CreateTempFile creates a temp file for the testing db
func CreateTempFile(t *testing.T, initialData string) (*os.File, func()) {
	t.Helper()

	tmpfile, err := ioutil.TempFile("", "db")

	if err != nil {
		t.Fatalf("could not create temp file %v", err)
	}

	tmpfile.Write([]byte(initialData))

	removeFile := func() {
		tmpfile.Close()
		os.Remove(tmpfile.Name())
	}

	return tmpfile, removeFile
}

// AssertContentType asserts that the content-type header is correct
func AssertContentType(t *testing.T, res *httptest.ResponseRecorder, want string) {
	t.Helper()
	if res.Result().Header.Get("content-type") != want {
		t.Errorf("resoinse did not have content-type of %s. got %v", want, res.Result().Header)
	}

}

// AssertLeague returns the expected []Player
func AssertLeague(t *testing.T, got, want []Player) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

// AssertResponseBody is the response expected
func AssertResponseBody(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %q, want %q", got, want)
	}
}

// AssertStatus is the expected status code
func AssertStatus(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}

// AssertPlayerWin records the win for the correct player
func AssertPlayerWin(t *testing.T, store *StubPlayerStore, winner string) {
	t.Helper()

	if len(store.winCalls) != 1 {
		t.Fatalf("got %d calls to RecordWin want %d", len(store.winCalls), 1)
	}

	if store.winCalls[0] != winner {
		t.Errorf("didn't record correct winner, got %q, want %q", store.winCalls, winner)
	}
}

// AssertNoError is returned
func AssertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("didn't expect an error but got one, %v", err)
	}
}

// AssertScoreEquals the desired score
func AssertScoreEquals(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
