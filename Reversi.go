package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Constants
const (
	BoardSize   = 8
	EmptyCell   = " . "
	PlayerBlack = " ● "
	PlayerWhite = " ○ "
)

// Game represents the state of the game.
type Game struct {
	board       [][]string
	currentTurn string
}

// NewGame initializes a new game.
func NewGame() *Game {
	board := make([][]string, BoardSize)
	for i := range board {
		board[i] = make([]string, BoardSize)
		for j := range board[i] {
			board[i][j] = EmptyCell
		}
	}

	// Starting pieces
	board[3][3] = PlayerWhite
	board[3][4] = PlayerBlack
	board[4][3] = PlayerBlack
	board[4][4] = PlayerWhite

	return &Game{
		board:       board,
		currentTurn: PlayerBlack,
	}
}

// PrintBoard prints the current state of the game board.
func (g *Game) PrintBoard() {
	fmt.Println("   A  B  C  D  E  F  G  H")
	for i, row := range g.board {
		fmt.Printf("%d ", i+1)
		for _, cell := range row {
			fmt.Printf("%s", cell)
		}
		fmt.Println()
	}
}

// PlayMove plays a move on the game board.
func (g *Game) PlayMove(row, col int) bool {
	if g.isValidMove(row, col) {
		g.board[row][col] = g.currentTurn
		g.flipPieces(row, col)

		// Check if the game is over after the current move
		if g.isGameOver() {
			return true
		}

		// If the other player doesn't have a valid move, the current player continues the turn
		if !g.hasValidMove(g.getOtherPlayer()) {
			fmt.Printf("%s cannot make a valid move. %s continues the turn.\n", g.getOtherPlayer(), g.currentTurn)
			return true
		}

		// Switch the turn to the other player
		g.switchTurn()

		return true
	}
	return false
}

// isValidMove checks if a move is valid.
func (g *Game) isValidMove(row, col int) bool {
	if row < 0 || row >= BoardSize || col < 0 || col >= BoardSize || g.board[row][col] != EmptyCell {
		return false
	}

	// Check if at least one piece can be flipped in any direction
	for dr := -1; dr <= 1; dr++ {
		for dc := -1; dc <= 1; dc++ {
			if dr == 0 && dc == 0 {
				continue
			}
			if g.checkDirection(row, col, dr, dc) {
				return true
			}
		}
	}

	return false
}

// checkDirection checks if there are pieces to flip in a specific direction.
func (g *Game) checkDirection(row, col, dr, dc int) bool {
	otherPlayer := g.getOtherPlayer()

	r, c := row+dr, col+dc

	// Check if the first piece in the direction is the other player's piece
	if r >= 0 && r < BoardSize && c >= 0 && c < BoardSize && g.board[r][c] == otherPlayer {
		// Keep moving in the direction until an empty cell or the current player's piece is found
		for r >= 0 && r < BoardSize && c >= 0 && c < BoardSize {
			if g.board[r][c] == EmptyCell {
				return false
			} else if g.board[r][c] == g.currentTurn {
				return true
			}
			r += dr
			c += dc
		}
	}

	return false
}

// flipPieces flips the opponent's pieces in all valid directions after playing a move.
func (g *Game) flipPieces(row, col int) {
	for dr := -1; dr <= 1; dr++ {
		for dc := -1; dc <= 1; dc++ {
			if dr == 0 && dc == 0 {
				continue
			}
			if g.checkDirection(row, col, dr, dc) {
				g.flipDirection(row, col, dr, dc)
			}
		}
	}
}

// flipDirection flips the opponent's pieces in a specific direction.
func (g *Game) flipDirection(row, col, dr, dc int) {
	r, c := row+dr, col+dc

	for r >= 0 && r < BoardSize && c >= 0 && c < BoardSize {
		if g.board[r][c] == g.currentTurn {
			return
		}

		g.board[r][c] = g.currentTurn
		r += dr
		c += dc
	}
}

// switchTurn switches the current turn to the other player.
func (g *Game) switchTurn() {
	if g.currentTurn == PlayerBlack {
		g.currentTurn = PlayerWhite
	} else {
		g.currentTurn = PlayerBlack
	}
}

// getOtherPlayer returns the player symbol of the other player.
func (g *Game) getOtherPlayer() string {
	if g.currentTurn == PlayerBlack {
		return PlayerWhite
	}
	return PlayerBlack
}

func (g *Game) isGameOver() bool {
	return !g.hasValidMove(PlayerBlack) && !g.hasValidMove(PlayerWhite)
}

// hasValidMove checks if the current player has a valid move.
func (g *Game) hasValidMove(player string) bool {
	for row := 0; row < BoardSize; row++ {
		for col := 0; col < BoardSize; col++ {
			if g.isValidMove(row, col) && g.board[row][col] == EmptyCell {
				return true
			}
		}
	}
	return false
}

// getWinner returns the winner of the game.
func (g *Game) getWinner() string {
	blackCount, whiteCount := 0, 0
	for row := 0; row < BoardSize; row++ {
		for col := 0; col < BoardSize; col++ {
			if g.board[row][col] == PlayerBlack {
				blackCount++
			} else if g.board[row][col] == PlayerWhite {
				whiteCount++
			}
		}
	}

	if blackCount > whiteCount {
		return PlayerBlack
	} else if whiteCount > blackCount {
		return PlayerWhite
	} else {
		return "Draw"
	}
}

func parseInput(input string) (int, int, error) {
	parts := strings.Fields(input)
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("invalid input format")
	}

	row, err := strconv.Atoi(parts[0])
	if err != nil || row < 1 || row > 8 {
		return 0, 0, fmt.Errorf("invalid row value")
	}

	col := int(parts[1][0] - 'A')
	if col < 0 || col >= BoardSize {
		return 0, 0, fmt.Errorf("invalid column value")
	}

	return row, col, nil
}

// getPieceCounts returns the counts of black and white pieces on the board.
func (g *Game) getPieceCounts() (int, int) {
	blackCount, whiteCount := 0, 0
	for row := 0; row < BoardSize; row++ {
		for col := 0; col < BoardSize; col++ {
			if g.board[row][col] == PlayerBlack {
				blackCount++
			} else if g.board[row][col] == PlayerWhite {
				whiteCount++
			}
		}
	}
	return blackCount, whiteCount
}

func main() {
	game := NewGame()

	scanner := bufio.NewScanner(os.Stdin)
	for !game.isGameOver() {
		game.PrintBoard()
		blackCount, whiteCount := game.getPieceCounts()
		fmt.Printf("Current Turn: %s\n", game.currentTurn)
		fmt.Printf("Black Pieces: %d\n", blackCount)
		fmt.Printf("White Pieces: %d\n", whiteCount)

		fmt.Print("Enter row (1-8) and column (A-H) separated by space: ")
		if scanner.Scan() {
			input := scanner.Text()
			row, col, err := parseInput(input)
			if err != nil {
				fmt.Println("Invalid input. Please try again.")
				continue
			}

			if game.PlayMove(row-1, col) {
				// The player switch is now handled within the PlayMove method.
			} else {
				fmt.Println("Invalid move. Please try again.")
			}
		}
	}

	game.PrintBoard()
	blackCount, whiteCount := game.getPieceCounts()
	fmt.Printf("Game Over! Winner: %s\n", game.getWinner())
	fmt.Printf("Black Pieces: %d\n", blackCount)
	fmt.Printf("White Pieces: %d\n", whiteCount)
}
