package poker

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// PlayerPrompt is the string for the number of players request
const PlayerPrompt = "Please enter the number of players: "

// BadPlayerInputErrMsg is the string returned when bad input is passed
const BadPlayerInputErrMsg = "don't be silly."

// CLI is the cli struct
type CLI struct {
	in   *bufio.Scanner
	out  io.Writer
	game Game
}

// NewCLI is a constructor for a CLI struct
func NewCLI(in io.Reader, out io.Writer, game Game) *CLI {
	return &CLI{
		in:   bufio.NewScanner(in),
		out:  out,
		game: game,
	}
}

// PlayPoker plays the poker game
func (cli *CLI) PlayPoker() {
	fmt.Fprint(cli.out, PlayerPrompt)
	numberOfPlayersInput := cli.readLine()
	numberOfPlayers, err := strconv.Atoi(strings.Trim(numberOfPlayersInput, "\n"))
	if err != nil {
		fmt.Fprint(cli.out, BadPlayerInputErrMsg)
		return
	}

	cli.game.Start(numberOfPlayers)

	winnerInput := cli.readLine()
	winner := (extractWinner(winnerInput))

	cli.game.Finish(winner)
}

func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1)
}

func (cli *CLI) readLine() string {
	cli.in.Scan()
	return cli.in.Text()
}
