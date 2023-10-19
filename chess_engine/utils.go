package chess_engine

import "fmt"

type ErrorCode uint8

const (
	NO_ERROR ErrorCode = iota
	INVALID_MOVE
	ENGINE_CHECK_MATE
	USER_CHECK_MATE
	NOT_USER_MOVE
	NOT_ENGINE_MOVE
)

type Result struct {
	Result  bool
	ResCode ErrorCode
	ResStr  string
}

func GetRankFile(pos BoardPosition) (int, int) {
	rank := -1
	file := -1
	switch pos {
	case A1:
		rank = 0
		file = 0
	case A2:
		rank = 1
		file = 0
	case A3:
		rank = 2
		file = 0
	case A4:
		rank = 3
		file = 0
	case A5:
		rank = 4
		file = 0
	case A6:
		rank = 5
		file = 0
	case A7:
		rank = 6
		file = 0
	case A8:
		rank = 7
		file = 0
	case B1:
		rank = 0
		file = 1
	case B2:
		rank = 1
		file = 1
	case B3:
		rank = 2
		file = 1
	case B4:
		rank = 3
		file = 1
	case B5:
		rank = 4
		file = 1
	case B6:
		rank = 5
		file = 1
	case B7:
		rank = 6
		file = 1
	case B8:
		rank = 7
		file = 1
	case C1:
		rank = 0
		file = 2
	case C2:
		rank = 1
		file = 2
	case C3:
		rank = 2
		file = 2
	case C4:
		rank = 3
		file = 2
	case C5:
		rank = 4
		file = 2
	case C6:
		rank = 5
		file = 2
	case C7:
		rank = 6
		file = 2
	case C8:
		rank = 7
		file = 2
	case D1:
		rank = 0
		file = 3
	case D2:
		rank = 1
		file = 3
	case D3:
		rank = 2
		file = 3
	case D4:
		rank = 3
		file = 3
	case D5:
		rank = 4
		file = 3
	case D6:
		rank = 5
		file = 3
	case D7:
		rank = 6
		file = 3
	case D8:
		rank = 7
		file = 3
	case E1:
		rank = 0
		file = 4
	case E2:
		rank = 1
		file = 4
	case E3:
		rank = 2
		file = 4
	case E4:
		rank = 3
		file = 4
	case E5:
		rank = 4
		file = 4
	case E6:
		rank = 5
		file = 4
	case E7:
		rank = 6
		file = 4
	case E8:
		rank = 7
		file = 4
	case F1:
		rank = 0
		file = 5
	case F2:
		rank = 1
		file = 5
	case F3:
		rank = 2
		file = 5
	case F4:
		rank = 3
		file = 5
	case F5:
		rank = 4
		file = 5
	case F6:
		rank = 5
		file = 5
	case F7:
		rank = 6
		file = 5
	case F8:
		rank = 7
		file = 5
	case G1:
		rank = 0
		file = 6
	case G2:
		rank = 1
		file = 6
	case G3:
		rank = 2
		file = 6
	case G4:
		rank = 3
		file = 6
	case G5:
		rank = 4
		file = 6
	case G6:
		rank = 5
		file = 6
	case G7:
		rank = 6
		file = 6
	case G8:
		rank = 7
		file = 6
	case H1:
		rank = 0
		file = 7
	case H2:
		rank = 1
		file = 7
	case H3:
		rank = 2
		file = 7
	case H4:
		rank = 3
		file = 7
	case H5:
		rank = 4
		file = 7
	case H6:
		rank = 5
		file = 7
	case H7:
		rank = 6
		file = 7
	case H8:
		rank = 7
		file = 7
	}

	return rank, file
}

func GetBoardPos(rank int, file int) BoardPosition {
	pos := INVALID_BOARD_POSITION
	if rank < 0 || rank >= 8 || file < 0 || file >= 8 {
		return pos
	}
	switch rank {
	case 0:
		switch file {
		case 0:
			pos = A1
			break
		case 1:
			pos = B1
			break
		case 2:
			pos = C1
			break
		case 3:
			pos = D1
			break
		case 4:
			pos = E1
			break
		case 5:
			pos = F1
			break
		case 6:
			pos = G1
			break
		case 7:
			pos = H1
			break
		}
		break
	case 1:
		switch file {
		case 0:
			pos = A2
			break
		case 1:
			pos = B2
			break
		case 2:
			pos = C2
			break
		case 3:
			pos = D2
			break
		case 4:
			pos = E2
			break
		case 5:
			pos = F2
			break
		case 6:
			pos = G2
			break
		case 7:
			pos = H2
			break
		}
		break
	case 2:
		switch file {
		case 0:
			pos = A3
			break
		case 1:
			pos = B3
			break
		case 2:
			pos = C3
			break
		case 3:
			pos = D3
			break
		case 4:
			pos = E3
			break
		case 5:
			pos = F3
			break
		case 6:
			pos = G3
			break
		case 7:
			pos = H3
			break
		}
		break
	case 3:
		switch file {
		case 0:
			pos = A4
			break
		case 1:
			pos = B4
			break
		case 2:
			pos = C4
			break
		case 3:
			pos = D4
			break
		case 4:
			pos = E4
			break
		case 5:
			pos = F4
			break
		case 6:
			pos = G4
			break
		case 7:
			pos = H4
			break
		}
		break
	case 4:
		switch file {
		case 0:
			pos = A5
			break
		case 1:
			pos = B5
			break
		case 2:
			pos = C5
			break
		case 3:
			pos = D5
			break
		case 4:
			pos = E5
			break
		case 5:
			pos = F5
			break
		case 6:
			pos = G5
			break
		case 7:
			pos = H5
			break
		}
		break
	case 5:
		switch file {
		case 0:
			pos = A6
			break
		case 1:
			pos = B6
			break
		case 2:
			pos = C6
			break
		case 3:
			pos = D6
			break
		case 4:
			pos = E6
			break
		case 5:
			pos = F6
			break
		case 6:
			pos = G6
			break
		case 7:
			pos = H6
			break
		}
		break
	case 6:
		switch file {
		case 0:
			pos = A7
			break
		case 1:
			pos = B7
			break
		case 2:
			pos = C7
			break
		case 3:
			pos = D7
			break
		case 4:
			pos = E7
			break
		case 5:
			pos = F7
			break
		case 6:
			pos = G7
			break
		case 7:
			pos = H7
			break
		}
		break
	case 7:
		switch file {
		case 0:
			pos = A8
			break
		case 1:
			pos = B8
			break
		case 2:
			pos = C8
			break
		case 3:
			pos = D8
			break
		case 4:
			pos = E8
			break
		case 5:
			pos = F8
			break
		case 6:
			pos = G8
			break
		case 7:
			pos = H8
			break
		}
		break
	}

	return pos
}

func MovesContainMove(move Move, moves []*Move) bool {
	retVal := false

	for _, m := range moves {
		if m.StartPos == move.StartPos &&
			m.EndPos == move.EndPos {
			retVal = true
			break
		}
	}

	return retVal
}

func MovesContainEndPos(endPos BoardPosition, moves []*Move) bool {
	retVal := false

	for _, m := range moves {
		if m.EndPos == endPos {
			retVal = true
			break
		}
	}

	return retVal
}

func OppositeColor(color PieceColor) PieceColor {
	if color == White {
		return Black
	}

	return White
}

func GetMoveFromUserString(startStr string, endStr string) *Move {
	var move *Move = nil
	startPos := GetPositionFromString(startStr)
	endPos := GetPositionFromString(endStr)
	if startPos != INVALID_BOARD_POSITION && endPos != INVALID_BOARD_POSITION {
		move = NewMove(startPos, endPos)
	}

	return move
}

func GetPositionFromString(inStr string) BoardPosition {
	var pos BoardPosition = INVALID_BOARD_POSITION
	if inStr == "A1" {
		pos = A1
	} else if inStr == "A2" {
		pos = A2
	} else if inStr == "A3" {
		pos = A3
	} else if inStr == "A4" {
		pos = A4
	} else if inStr == "A5" {
		pos = A5
	} else if inStr == "A6" {
		pos = A6
	} else if inStr == "A7" {
		pos = A7
	} else if inStr == "A8" {
		pos = A8
	} else if inStr == "B1" {
		pos = B1
	} else if inStr == "B2" {
		pos = B2
	} else if inStr == "B3" {
		pos = B3
	} else if inStr == "B4" {
		pos = B4
	} else if inStr == "B5" {
		pos = B5
	} else if inStr == "B6" {
		pos = B6
	} else if inStr == "B7" {
		pos = B7
	} else if inStr == "B8" {
		pos = B8
	} else if inStr == "C1" {
		pos = C1
	} else if inStr == "C2" {
		pos = C2
	} else if inStr == "C3" {
		pos = C3
	} else if inStr == "C4" {
		pos = C4
	} else if inStr == "C5" {
		pos = C5
	} else if inStr == "C6" {
		pos = C6
	} else if inStr == "C7" {
		pos = C7
	} else if inStr == "C8" {
		pos = C8
	} else if inStr == "D1" {
		pos = D1
	} else if inStr == "D2" {
		pos = D2
	} else if inStr == "D3" {
		pos = D3
	} else if inStr == "D4" {
		pos = D4
	} else if inStr == "D5" {
		pos = D5
	} else if inStr == "D6" {
		pos = D6
	} else if inStr == "D7" {
		pos = D7
	} else if inStr == "D8" {
		pos = D8
	} else if inStr == "E1" {
		pos = E1
	} else if inStr == "E2" {
		pos = E2
	} else if inStr == "E3" {
		pos = E3
	} else if inStr == "E4" {
		pos = E4
	} else if inStr == "E5" {
		pos = E5
	} else if inStr == "E6" {
		pos = E6
	} else if inStr == "E7" {
		pos = E7
	} else if inStr == "E8" {
		pos = E8
	} else if inStr == "F1" {
		pos = F1
	} else if inStr == "F2" {
		pos = F2
	} else if inStr == "F3" {
		pos = F3
	} else if inStr == "F4" {
		pos = F4
	} else if inStr == "F5" {
		pos = F5
	} else if inStr == "F6" {
		pos = F6
	} else if inStr == "F7" {
		pos = F7
	} else if inStr == "F8" {
		pos = F8
	} else if inStr == "G1" {
		pos = G1
	} else if inStr == "G2" {
		pos = G2
	} else if inStr == "G3" {
		pos = G3
	} else if inStr == "G4" {
		pos = G4
	} else if inStr == "G5" {
		pos = G5
	} else if inStr == "G6" {
		pos = G6
	} else if inStr == "G7" {
		pos = G7
	} else if inStr == "G8" {
		pos = G8
	} else if inStr == "H1" {
		pos = H1
	} else if inStr == "H2" {
		pos = H2
	} else if inStr == "H3" {
		pos = H3
	} else if inStr == "H4" {
		pos = H4
	} else if inStr == "H5" {
		pos = H5
	} else if inStr == "H6" {
		pos = H6
	} else if inStr == "H7" {
		pos = H7
	} else if inStr == "H8" {
		pos = H8
	}

	return pos
}

func GetPositionString(pos BoardPosition) string {
	var retStr string = "IV"
	if pos == A1 {
		retStr = "A1"
	} else if pos == A2 {
		retStr = "A2"
	} else if pos == A3 {
		retStr = "A3"
	} else if pos == A4 {
		retStr = "A4"
	} else if pos == A5 {
		retStr = "A5"
	} else if pos == A6 {
		retStr = "A6"
	} else if pos == A7 {
		retStr = "A7"
	} else if pos == A8 {
		retStr = "A8"
	} else if pos == B1 {
		retStr = "B1"
	} else if pos == B2 {
		retStr = "B2"
	} else if pos == B3 {
		retStr = "B3"
	} else if pos == B4 {
		retStr = "B4"
	} else if pos == B5 {
		retStr = "B5"
	} else if pos == B6 {
		retStr = "B6"
	} else if pos == B7 {
		retStr = "B7"
	} else if pos == B8 {
		retStr = "B8"
	} else if pos == C1 {
		retStr = "C1"
	} else if pos == C2 {
		retStr = "C2"
	} else if pos == C3 {
		retStr = "C3"
	} else if pos == C4 {
		retStr = "C4"
	} else if pos == C5 {
		retStr = "C5"
	} else if pos == C6 {
		retStr = "C6"
	} else if pos == C7 {
		retStr = "C7"
	} else if pos == C8 {
		retStr = "C8"
	} else if pos == D1 {
		retStr = "D1"
	} else if pos == D2 {
		retStr = "D2"
	} else if pos == D3 {
		retStr = "D3"
	} else if pos == D4 {
		retStr = "D4"
	} else if pos == D5 {
		retStr = "D5"
	} else if pos == D6 {
		retStr = "D6"
	} else if pos == D7 {
		retStr = "D7"
	} else if pos == D8 {
		retStr = "D8"
	} else if pos == E1 {
		retStr = "E1"
	} else if pos == E2 {
		retStr = "E2"
	} else if pos == E3 {
		retStr = "E3"
	} else if pos == E4 {
		retStr = "E4"
	} else if pos == E5 {
		retStr = "E5"
	} else if pos == E6 {
		retStr = "E6"
	} else if pos == E7 {
		retStr = "E7"
	} else if pos == E8 {
		retStr = "E8"
	} else if pos == F1 {
		retStr = "F1"
	} else if pos == F2 {
		retStr = "F2"
	} else if pos == F3 {
		retStr = "F3"
	} else if pos == F4 {
		retStr = "F4"
	} else if pos == F5 {
		retStr = "F5"
	} else if pos == F6 {
		retStr = "F6"
	} else if pos == F7 {
		retStr = "F7"
	} else if pos == F8 {
		retStr = "F8"
	} else if pos == G1 {
		retStr = "G1"
	} else if pos == G2 {
		retStr = "G2"
	} else if pos == G3 {
		retStr = "G3"
	} else if pos == G4 {
		retStr = "G4"
	} else if pos == G5 {
		retStr = "G5"
	} else if pos == G6 {
		retStr = "G6"
	} else if pos == G7 {
		retStr = "G7"
	} else if pos == G8 {
		retStr = "G8"
	} else if pos == H1 {
		retStr = "H1"
	} else if pos == H2 {
		retStr = "H2"
	} else if pos == H3 {
		retStr = "H3"
	} else if pos == H4 {
		retStr = "H4"
	} else if pos == H5 {
		retStr = "H5"
	} else if pos == H6 {
		retStr = "H6"
	} else if pos == H7 {
		retStr = "H7"
	} else if pos == H8 {
		retStr = "H8"
	}

	return retStr
}

func GetMoveString(move Move) string {
	return fmt.Sprintf("%s - %s", GetPositionString(move.StartPos), GetPositionString(move.EndPos))
}
