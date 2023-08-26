package chess_engine

import "testing"

func TestNewGame(t *testing.T) {
	g := NewGame(White)
	if g == nil {
		t.Errorf("Error! NewGame function returned nil for game!")
	}

	if g.Board == nil {
		t.Errorf("Error! NewGame function returned nil for board!")
	}

	if g.UserColor != White {
		t.Errorf("Error! NewGame function returned black for user color!")
	}
}

func TestGameUserMove(t *testing.T) {
	g := NewGame(Black)
	res := g.UserMove(NewMove(A7, A6))
	if res.Result {
		t.Errorf("Error! Game.UserMove function returned success for invalid user move (wrong color moving)!")
	}

	g = NewGame(White)
	res = g.UserMove(NewMove(A2, A3))
	if !res.Result {
		t.Errorf("Error! Game.UserMove function returned failure for valid user move!")
	}
	if len(g.Moves) != 1 {
		t.Errorf("Error! Game.Moves array length not 1 after making first user move!!!")
	}

	g = NewGame(White)
	res = g.UserMove(NewMove(A2, A8))
	if res.Result {
		t.Errorf("Error! Game.UserMove function returned success for invalid user move!")
	}

	g = &Game{}
	g.Board = &Board{}
	g.EngineColor = White
	g.UserColor = Black
	g.Status = GAME_STATUS_USER_MOVE
	g.Board.Squares[A1] = Square{true, NewKing(A1, White)}
	g.Board.Squares[A3] = Square{true, NewQueen(A3, Black)}
	g.Board.Squares[C3] = Square{true, NewQueen(C3, Black)}

	g.UserMove(NewMove(C3, B3))
	if g.Status != GAME_STATUS_ENGINE_CHECKMATE {
		t.Errorf("Error! Game.UserMove function did not set status to Engine Check Mate!")
	}

}

func TestGameUpdateGameStatus(t *testing.T) {
	g := &Game{}
	g.Board = &Board{}
	g.EngineColor = White
	g.UserColor = Black
	g.Board.Squares[A1] = Square{true, NewKing(A1, White)}
	g.Board.Squares[A3] = Square{true, NewQueen(A3, Black)}
	g.Board.Squares[B3] = Square{true, NewQueen(B3, Black)}
	g.UpdateGameStatus()
	if g.Status != GAME_STATUS_ENGINE_CHECKMATE {
		t.Errorf("Error! Game.UpdateGameStatus function did not set status to Engine Check Mate!")
	}

	g.EngineColor = Black
	g.UserColor = White
	g.UpdateGameStatus()
	if g.Status != GAME_STATUS_USER_CHECKMATE {
		t.Errorf("Error! Game.UpdateGameStatus function did not set status to User Check Mate!")
	}

	g = &Game{}
	g.Board = &Board{}
	g.Status = GAME_STATUS_USER_MOVE
	g.EngineColor = White
	g.UserColor = Black
	g.Board.Squares[A1] = Square{true, NewKing(A1, White)}
	g.Board.Squares[A3] = Square{true, NewQueen(A3, Black)}
	g.UpdateGameStatus()
	if g.Status != GAME_STATUS_ENGINE_MOVE {
		t.Errorf("Error! Game.UpdateGameStatus function did not set status to Engine Move!")
	}

	g.UpdateGameStatus()
	if g.Status != GAME_STATUS_USER_MOVE {
		t.Errorf("Error! Game.UpdateGameStatus function did not set status to User Move!")
	}
}

func TestGameEngineMoveOneValidMove(t *testing.T) {
	g := &Game{}
	g.Board = &Board{}
	g.Status = GAME_STATUS_ENGINE_MOVE
	g.EngineColor = White
	g.UserColor = Black
	g.Picker = NewBasicMovePicker(g.EngineColor)
	g.Board.Squares[A1] = Square{true, NewKing(A1, White)}
	g.Board.Squares[A3] = Square{true, NewQueen(A3, Black)}
	res, move := g.EngineMove()

	if !res.Result {
		t.Errorf("Error! Game.EngineMove did not successfully make an Engine Move!")
	} else {
		if move.StartPos != A1 && move.EndPos != B1 {
			t.Errorf("Error! Game.EngineMove did not move to the only valid square! (invalid move positions)")
		}
		if g.Board.Squares[A1].Occupied {
			t.Errorf("Error! Game.EngineMove did not move to the only valid square! (A1 still occupied)")
		}
		if !g.Board.Squares[B1].Occupied || g.Board.Squares[B1].CurrPiece == nil {
			t.Errorf("Error! Game.EngineMove did not move to the only valid square! (B1 not occupied)")
		} else {
			if g.Board.Squares[B1].CurrPiece.GetType() != KING && g.Board.Squares[B1].CurrPiece.GetColor() != Black {
				t.Errorf("Error! Game.EngineMove did not move to the only valid square (invalid type and piece)!")
			}
		}
		if g.Status != GAME_STATUS_USER_MOVE {
			t.Errorf("Error! Game.UpdateGameStatus function did not set status to Engine Move!")
		}
	}
}

func TestGameEngineMoveFirstMoveValid(t *testing.T) {
	game := NewGame(Black)
	validMoves := game.Board.GetMoves(White)
	res, move := game.EngineMove()
	if !res.Result {
		t.Errorf("Error! Game.EngineMove returned failure on new game first move!")
	}
	if move == nil {
		t.Errorf("Error! Game.EngineMove return nil for move on new game first move!")
	} else {
		if !MovesContainMove(*move, validMoves) {
			t.Errorf("Error! Game.EngineMove return invalid move on new game first move!")
		}
	}
	if game.Status != GAME_STATUS_USER_MOVE {
		t.Errorf("Error! Game.UpdateGameStatus function did not set status to Engine Move!")
	}
}

func TestGameEngineMovePicksCheckMate(t *testing.T) {
	g := &Game{}
	g.Board = &Board{}
	g.Status = GAME_STATUS_ENGINE_MOVE
	g.EngineColor = Black
	g.UserColor = White
	g.Picker = NewBasicMovePicker(g.EngineColor)
	g.Board.Squares[A1] = Square{true, NewKing(A1, White)}
	g.Board.Squares[A3] = Square{true, NewQueen(A3, Black)}
	g.Board.Squares[C4] = Square{true, NewKnight(C4, Black)}
	res, move := g.EngineMove()

	if !res.Result {
		t.Errorf("Error! Game.EngineMove did not successfully make an Engine Move!")
	} else {
		if move.StartPos != A3 && move.EndPos != B2 {
			t.Errorf("Error! Game.EngineMove did not move to the only CheckMate square!")
		}
		if g.Status == GAME_STATUS_USER_CHECKMATE {
			t.Errorf("Error! Game.EngineMove did not update game.Status to GAME_STATUS_USER_CHECKMATE!")
		}
	}
}
