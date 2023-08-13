package chess_engine

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
