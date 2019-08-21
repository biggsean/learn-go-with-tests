package poker

import "time"

// Game is an interface for games
type Game interface {
	Start(numberOfPlayers int)
	Finish(winner string)
}

// TexasHoldem is the poker game
type TexasHoldem struct {
	alerter BlindAlerter
	store   PlayerStore
}

// Start starts the game
func (p *TexasHoldem) Start(numberOfPlayers int) {
	blindIncrement := time.Duration(5+numberOfPlayers) * time.Minute

	blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
	blindTime := 0 * time.Second
	for _, blind := range blinds {
		p.alerter.ScheduleAlertAt(blindTime, blind)
		blindTime = blindTime + blindIncrement
	}
}

func NewTexasHoldem(alerter BlindAlerter, store PlayerStore) *TexasHoldem {
	return &TexasHoldem{
		alerter: alerter,
		store:   store,
	}
}

// Finish finishes the game
func (p *TexasHoldem) Finish(winner string) {
	p.store.RecordWin(winner)
}
