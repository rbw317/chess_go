package chess_engine

type Game struct {
	board *Board
	moves []*Move
}

func NewGame() *Game {
	g := &Game{}
	g.board = NewBoard()
	return g
}

func (game *Game) MakeMove(move *Move) (bool, string) {
	return true, "Valid Move"
}

func (game *Game) CheckMate() bool {
	return false
}

// Make move.  Needs to fail for invalid moves
// Invalid moves
// Moving from a square with no piece
// Moving wrong color piece (i.e. move white piece on black's turn)
// Moving a piece to a square it can't get to based on how that piece can move
// Moving to a square that is already occupied by a piece of the same color
// Moving a piece other than the king while king is in check
// Moving any piece after the king is in checkmate (i.e. game is over)
// Need to validate state of the board after the move.
// Move without an attack
// Move with an attack that removes one of the opposing pieces

// Check for check and checkmate.

// Generate engine move
// Search for best move (should be configurable how many moves to look ahead)
// Return new move
