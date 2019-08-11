package poker_test

import poker "github.com/biggsean/learn-go-with-tests/game"
import "strings"
import "testing"

func TestCLI(t *testing.T) {
	t.Run("record sean win from user input", func(t *testing.T) {
		in := strings.NewReader("Sean wins\n")
		playerStore := &poker.StubPlayerStore{}
		cli := poker.NewCLI(playerStore, in)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, playerStore, "Sean")
	})
	t.Run("record shane win from user input", func(t *testing.T) {
		in := strings.NewReader("Shane wins\n")
		playerStore := &poker.StubPlayerStore{}
		cli := poker.NewCLI(playerStore, in)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, playerStore, "Shane")
	})
}
