package main

import (
	"reflect"
	"testing"
)

func TestNewGame(t *testing.T) {
	game := NewGame()

	expectedBoard := [][]string{
		{EmptyCell, EmptyCell, EmptyCell, EmptyCell, EmptyCell, EmptyCell, EmptyCell, EmptyCell},
		{EmptyCell, EmptyCell, EmptyCell, EmptyCell, EmptyCell, EmptyCell, EmptyCell, EmptyCell},
		{EmptyCell, EmptyCell, EmptyCell, EmptyCell, EmptyCell, EmptyCell, EmptyCell, EmptyCell},
		{EmptyCell, EmptyCell, EmptyCell, PlayerWhite, PlayerBlack, EmptyCell, EmptyCell, EmptyCell},
		{EmptyCell, EmptyCell, EmptyCell, PlayerBlack, PlayerWhite, EmptyCell, EmptyCell, EmptyCell},
		{EmptyCell, EmptyCell, EmptyCell, EmptyCell, EmptyCell, EmptyCell, EmptyCell, EmptyCell},
		{EmptyCell, EmptyCell, EmptyCell, EmptyCell, EmptyCell, EmptyCell, EmptyCell, EmptyCell},
		{EmptyCell, EmptyCell, EmptyCell, EmptyCell, EmptyCell, EmptyCell, EmptyCell, EmptyCell},
	}

	if !reflect.DeepEqual(game.board, expectedBoard) {
		t.Errorf("NewGame() failed, expected board:\n%v\ngot:\n%v", expectedBoard, game.board)
	}

	if game.currentTurn != PlayerBlack {
		t.Errorf("NewGame() failed, expected current turn: %s, got: %s", PlayerBlack, game.currentTurn)
	}
}

func TestGame_PrintBoard(t *testing.T) {
	game := NewGame()

	expectedOutput :=
		`   A  B  C  D  E  F  G  H
1  .  .  .  .  .  .  .  .
2  .  .  .  .  .  .  .  .
3  .  .  .  .  .  .  .  .
4  .  .  .  ○   ●  .  .  .
5  .  .  .  ●   ○   .  .  .
6  .  .  .  .  .  .  .  .
7  .  .  .  .  .  .  .  .
8  .  .  .  .  .  .  .  .
`

	// Redirect stdout to a buffer
	buf := captureOutput(func() {
		game.PrintBoard()
	})

	if buf.String() != expectedOutput {
		t.Errorf("PrintBoard() failed, expected output:\n%s\ngot:\n%s", expectedOutput, buf.String())
	}
}

func TestGame_PlayMove(t *testing.T) {
	game := NewGame()

	// Valid move
	if !game.PlayMove(2, 3) {
		t.Errorf("PlayMove(2, 3) failed, expected true")
	}

	expectedBoard := [][]string{
		{EmptyCell, EmptyCell, EmptyCell, EmptyCell, EmptyCell, EmptyCell, EmptyCell, EmptyCell},
		{EmptyCell, EmptyCell, EmptyCell, PlayerBlack, EmptyCell, EmptyCell, EmptyCell, EmptyCell},
		{EmptyCell, EmptyCell, EmptyCell, ., EmptyCell, EmptyCell, EmptyCell, EmptyCell},
		{EmptyCell, EmptyCell, EmptyCell, PlayerWhite, PlayerBlack, EmptyCell, EmptyCell, EmptyCell},
		{EmptyCell, EmptyCell, EmptyCell, PlayerBlack, PlayerWhite, EmptyCell, EmptyCell, EmptyCell},
		{EmptyCell, EmptyCell, EmptyCell, EmptyCell, EmptyCell, EmptyCell, EmptyCell, EmptyCell},
		{EmptyCell, EmptyCell, EmptyCell, EmptyCell, EmptyCell, EmptyCell, EmptyCell, EmptyCell},
		{EmptyCell, EmptyCell, EmptyCell, EmptyCell, EmptyCell, EmptyCell, EmptyCell, EmptyCell},
	}

	if !reflect.DeepEqual(game.board, expectedBoard) {
		t.Errorf("PlayMove(2, 3) failed, expected board:\n%v\ngot:\n%v", expectedBoard, game.board)
	}

	// Invalid move
	if game.PlayMove(0, 0) {
		t.Errorf("PlayMove(0, 0) failed, expected false")
	}
}

func TestGame_isValidMove(t *testing.T) {
	game := NewGame()

	// Valid move
	if !game.isValidMove(2, 3) {
		t.Errorf("isValidMove(2, 3) failed, expected true")
	}

	// Invalid move: occupied cell
	if game.isValidMove(3, 3) {
		t.Errorf("isValidMove(3, 3) failed, expected false")
	}

	// Invalid move: out of bounds
	if game.isValidMove(8, 0) {
		t.Errorf("isValidMove(8, 0) failed, expected false")
	}
}

func TestGame_checkDirection(t *testing.T) {
	game := NewGame()

	// Valid direction
	if !game.checkDirection(3, 3, 1, 1) {
		t.Errorf("checkDirection(3, 3, 1, 1) failed, expected true")
	}

	// Invalid direction
	if game.checkDirection(3, 3, -1, -1) {
		t.Errorf("checkDirection(3, 3, -1, -1) failed, expected false")
	}
}

func TestGame_flipPieces(t *testing.T) {
	game := NewGame()

	game.flipPieces(3, 3)

	expectedBoard := [][]string{
		{EmptyCell, EmptyCell, EmptyCell, EmptyCell, EmptyCell, EmptyCell, EmptyCell, EmptyCell},
		{EmptyCell, EmptyCell, EmptyCell, PlayerBlack, EmptyCell, EmptyCell, EmptyCell, EmptyCell},
		{EmptyCell, EmptyCell, EmptyCell, PlayerBlack, EmptyCell, EmptyCell, EmptyCell, EmptyCell},
		{EmptyCell, EmptyCell, EmptyCell, PlayerBlack, PlayerBlack, EmptyCell, EmptyCell, EmptyCell},
		{EmptyCell, EmptyCell, EmptyCell, PlayerBlack, PlayerWhite, EmptyCell, EmptyCell, EmptyCell},
		{EmptyCell, EmptyCell, EmptyCell, EmptyCell, EmptyCell, EmptyCell, EmptyCell, EmptyCell},
		{EmptyCell, EmptyCell, EmptyCell, EmptyCell, EmptyCell, EmptyCell, EmptyCell, EmptyCell},
		{EmptyCell, EmptyCell, EmptyCell, EmptyCell, EmptyCell, EmptyCell, EmptyCell, EmptyCell},
	}

	if !reflect.DeepEqual(game.board, expectedBoard) {
		t.Errorf("flipPieces(3, 3) failed, expected board:\n%v\ngot:\n%v", expectedBoard, game.board)
	}
}

func TestGame_flipDirection(t *testing.T) {
	game := NewGame()

	game.flipDirection(3, 3, 1, 1)

	expectedBoard := [][]string{
		{EmptyCell, EmptyCell, EmptyCell, EmptyCell, EmptyCell, EmptyCell, EmptyCell, EmptyCell},
		{EmptyCell, EmptyCell, EmptyCell, EmptyCell, PlayerBlack, EmptyCell, EmptyCell, EmptyCell},
		{EmptyCell, EmptyCell, EmptyCell, EmptyCell, EmptyCell, PlayerBlack, EmptyCell, EmptyCell},
		{EmptyCell, EmptyCell, EmptyCell, EmptyCell, EmptyCell, EmptyCell, PlayerBlack, EmptyCell},
		{EmptyCell, EmptyCell, EmptyCell, EmptyCell, PlayerBlack, PlayerWhite, EmptyCell, EmptyCell},
		{EmptyCell, EmptyCell, EmptyCell, EmptyCell, EmptyCell, EmptyCell, EmptyCell, EmptyCell},
		{EmptyCell, EmptyCell, EmptyCell, EmptyCell, EmptyCell, EmptyCell, EmptyCell, EmptyCell},
		{EmptyCell, EmptyCell, EmptyCell, EmptyCell, EmptyCell, EmptyCell, EmptyCell, EmptyCell},
	}

	if !reflect.DeepEqual(game.board, expectedBoard) {
		t.Errorf("flipDirection(3, 3, 1, 1) failed, expected board:\n%v\ngot:\n%v", expectedBoard, game.board)
	}
}

func TestGame_switchTurn(t *testing.T) {
	game := NewGame()

	// Initial turn is PlayerBlack
	if game.currentTurn != PlayerBlack {
		t.Errorf("Initial switchTurn() failed, expected current turn: %s, got: %s", PlayerBlack, game.currentTurn)
	}

	game.switchTurn()

	// After switching turn, it should be PlayerWhite
	if game.currentTurn != PlayerWhite {
		t.Errorf("switchTurn() failed, expected current turn: %s, got: %s", PlayerWhite, game.currentTurn)
	}
}

func TestGame_getOtherPlayer(t *testing.T) {
	game := NewGame()

	// Initial turn is PlayerBlack
	if game.getOtherPlayer() != PlayerWhite {
		t.Errorf("Initial getOtherPlayer() failed, expected other player: %s, got: %s", PlayerWhite, game.getOtherPlayer())
	}

	game.switchTurn()

	// After switching turn, other player should be PlayerBlack again
	if game.getOtherPlayer() != PlayerBlack {
		t.Errorf("getOtherPlayer() failed, expected other player: %s, got: %s", PlayerBlack, game.getOtherPlayer())
	}
}

func TestGame_isGameOver(t *testing.T) {
	game := NewGame()

	// Initial game is not over
	if game.isGameOver() {
		t.Errorf("Initial isGameOver() failed, expected false")
	}

	// Make a move that ends the game
	game.PlayMove(2, 3)
	game.PlayMove(2, 4)
	game.PlayMove(3, 2)
	game.PlayMove(3, 5)
	game.PlayMove(4, 3)
	game.PlayMove(4, 4)
	game.PlayMove(5, 4)

	// Game is now over, as no more valid moves can be made
	if !game.isGameOver() {
		t.Errorf("isGameOver() failed, expected true")
	}
}

func TestGame_hasValidMove(t *testing.T) {
	game := NewGame()

	// Initial game has valid moves for both players
	if !game.hasValidMove(PlayerBlack) || !game.hasValidMove(PlayerWhite) {
		t.Errorf("Initial hasValidMove() failed, expected true for both players")
	}

	// Block all valid moves for both players
	game.PlayMove(2, 3)
	game.PlayMove(2, 4)
	game.PlayMove(3, 2)
	game.PlayMove(3, 5)
	game.PlayMove(4, 3)
	game.PlayMove(4, 4)
	game.PlayMove(5, 4)

	// Both players have no valid moves
	if game.hasValidMove(PlayerBlack) || game.hasValidMove(PlayerWhite) {
		t.Errorf("hasValidMove() failed, expected false for both players")
	}
}

func TestGame_getWinner(t *testing.T) {
	game := NewGame()

	// Initial game has no winner yet
	if winner := game.getWinner(); winner != "" {
		t.Errorf("Initial getWinner() failed, expected \"\", got: %s", winner)
	}

	// Make some moves to finish the game
	game.PlayMove(2, 3)
	game.PlayMove(2, 4)
	game.PlayMove(3, 2)
	game.PlayMove(3, 5)
	game.PlayMove(4, 3)
	game.PlayMove(4, 4)
	game.PlayMove(5, 4)

	// PlayerBlack has more pieces than PlayerWhite, so PlayerBlack wins
	if winner := game.getWinner(); winner != PlayerBlack {
		t.Errorf("getWinner() failed, expected winner: %s, got: %s", PlayerBlack, winner)
	}
}
