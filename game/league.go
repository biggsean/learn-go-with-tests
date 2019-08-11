package main

import (
	"encoding/json"
	"fmt"
	"io"
)

// League is a slice of type Player
type League []Player

// Find iterates through the league and returns a player
func (l League) Find(name string) *Player {
	for i, p := range l {
		if p.Name == name {
			return &l[i]
		}
	}
	return nil
}

// NewLeague parses the json and returns League
func NewLeague(rdr io.Reader) (League, error) {
	var league []Player
	err := json.NewDecoder(rdr).Decode(&league)
	if err != nil {
		err = fmt.Errorf("problem parsing league, %v", err)
	}
	return league, err
}
