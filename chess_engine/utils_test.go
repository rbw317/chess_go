package chess_engine

import "testing"

func TestGetRankAndFileFromPos(t *testing.T) {
	rank, file := GetRankFile(A1)
	if rank != 0 || file != 0 {
		t.Errorf("Error! GetRankFile return invalid values for A1!")
	}

	rank, file = GetRankFile(A2)
	if rank != 1 || file != 0 {
		t.Errorf("Error! GetRankFile return invalid values for A2!")
	}

	rank, file = GetRankFile(A3)
	if rank != 2 || file != 0 {
		t.Errorf("Error! GetRankFile return invalid values for A3!")
	}

	rank, file = GetRankFile(A4)
	if rank != 3 || file != 0 {
		t.Errorf("Error! GetRankFile return invalid values for A4!")
	}

	rank, file = GetRankFile(A5)
	if rank != 4 || file != 0 {
		t.Errorf("Error! GetRankFile return invalid values for A5!")
	}

	rank, file = GetRankFile(A6)
	if rank != 5 || file != 0 {
		t.Errorf("Error! GetRankFile return invalid values for A6!")
	}

	rank, file = GetRankFile(A7)
	if rank != 6 || file != 0 {
		t.Errorf("Error! GetRankFile return invalid values for A7!")
	}

	rank, file = GetRankFile(A8)
	if rank != 7 || file != 0 {
		t.Errorf("Error! GetRankFile return invalid values for A8!")
	}

	rank, file = GetRankFile(B1)
	if rank != 0 || file != 1 {
		t.Errorf("Error! GetRankFile return invalid values for B1!")
	}

	rank, file = GetRankFile(B2)
	if rank != 1 || file != 1 {
		t.Errorf("Error! GetRankFile return invalid values for B2!")
	}

	rank, file = GetRankFile(B3)
	if rank != 2 || file != 1 {
		t.Errorf("Error! GetRankFile return invalid values for B3!")
	}

	rank, file = GetRankFile(B4)
	if rank != 3 || file != 1 {
		t.Errorf("Error! GetRankFile return invalid values for B4!")
	}

	rank, file = GetRankFile(B5)
	if rank != 4 || file != 1 {
		t.Errorf("Error! GetRankFile return invalid values for B5!")
	}

	rank, file = GetRankFile(B6)
	if rank != 5 || file != 1 {
		t.Errorf("Error! GetRankFile return invalid values for B6!")
	}

	rank, file = GetRankFile(B7)
	if rank != 6 || file != 1 {
		t.Errorf("Error! GetRankFile return invalid values for B7!")
	}

	rank, file = GetRankFile(B8)
	if rank != 7 || file != 1 {
		t.Errorf("Error! GetRankFile return invalid values for B8!")
	}

	rank, file = GetRankFile(C1)
	if rank != 0 || file != 2 {
		t.Errorf("Error! GetRankFile return invalid values for C1!")
	}

	rank, file = GetRankFile(C2)
	if rank != 1 || file != 2 {
		t.Errorf("Error! GetRankFile return invalid values for C2!")
	}

	rank, file = GetRankFile(C3)
	if rank != 2 || file != 2 {
		t.Errorf("Error! GetRankFile return invalid values for C3!")
	}

	rank, file = GetRankFile(C4)
	if rank != 3 || file != 2 {
		t.Errorf("Error! GetRankFile return invalid values for C4!")
	}

	rank, file = GetRankFile(C5)
	if rank != 4 || file != 2 {
		t.Errorf("Error! GetRankFile return invalid values for C5!")
	}

	rank, file = GetRankFile(C6)
	if rank != 5 || file != 2 {
		t.Errorf("Error! GetRankFile return invalid values for C6!")
	}

	rank, file = GetRankFile(C7)
	if rank != 6 || file != 2 {
		t.Errorf("Error! GetRankFile return invalid values for C7!")
	}

	rank, file = GetRankFile(C8)
	if rank != 7 || file != 2 {
		t.Errorf("Error! GetRankFile return invalid values for C8!")
	}

	rank, file = GetRankFile(D1)
	if rank != 0 || file != 3 {
		t.Errorf("Error! GetRankFile return invalid values for D1!")
	}

	rank, file = GetRankFile(D2)
	if rank != 1 || file != 3 {
		t.Errorf("Error! GetRankFile return invalid values for D2!")
	}

	rank, file = GetRankFile(D3)
	if rank != 2 || file != 3 {
		t.Errorf("Error! GetRankFile return invalid values for D3!")
	}

	rank, file = GetRankFile(D4)
	if rank != 3 || file != 3 {
		t.Errorf("Error! GetRankFile return invalid values for D4!")
	}

	rank, file = GetRankFile(D5)
	if rank != 4 || file != 3 {
		t.Errorf("Error! GetRankFile return invalid values for D5!")
	}

	rank, file = GetRankFile(D6)
	if rank != 5 || file != 3 {
		t.Errorf("Error! GetRankFile return invalid values for D6!")
	}

	rank, file = GetRankFile(D7)
	if rank != 6 || file != 3 {
		t.Errorf("Error! GetRankFile return invalid values for D7!")
	}

	rank, file = GetRankFile(D8)
	if rank != 7 || file != 3 {
		t.Errorf("Error! GetRankFile return invalid values for D8!")
	}

	rank, file = GetRankFile(E1)
	if rank != 0 || file != 4 {
		t.Errorf("Error! GetRankFile return invalid values for E1!")
	}

	rank, file = GetRankFile(E2)
	if rank != 1 || file != 4 {
		t.Errorf("Error! GetRankFile return invalid values for E2!")
	}

	rank, file = GetRankFile(E3)
	if rank != 2 || file != 4 {
		t.Errorf("Error! GetRankFile return invalid values for E3!")
	}

	rank, file = GetRankFile(E4)
	if rank != 3 || file != 4 {
		t.Errorf("Error! GetRankFile return invalid values for E4!")
	}

	rank, file = GetRankFile(E5)
	if rank != 4 || file != 4 {
		t.Errorf("Error! GetRankFile return invalid values for E5!")
	}

	rank, file = GetRankFile(E6)
	if rank != 5 || file != 4 {
		t.Errorf("Error! GetRankFile return invalid values for E6!")
	}

	rank, file = GetRankFile(E7)
	if rank != 6 || file != 4 {
		t.Errorf("Error! GetRankFile return invalid values for E7!")
	}

	rank, file = GetRankFile(E8)
	if rank != 7 || file != 4 {
		t.Errorf("Error! GetRankFile return invalid values for E8!")
	}

	rank, file = GetRankFile(F1)
	if rank != 0 || file != 5 {
		t.Errorf("Error! GetRankFile return invalid values for F1!")
	}

	rank, file = GetRankFile(F2)
	if rank != 1 || file != 5 {
		t.Errorf("Error! GetRankFile return invalid values for F2!")
	}

	rank, file = GetRankFile(F3)
	if rank != 2 || file != 5 {
		t.Errorf("Error! GetRankFile return invalid values for F3!")
	}

	rank, file = GetRankFile(F4)
	if rank != 3 || file != 5 {
		t.Errorf("Error! GetRankFile return invalid values for F4!")
	}

	rank, file = GetRankFile(F5)
	if rank != 4 || file != 5 {
		t.Errorf("Error! GetRankFile return invalid values for F5!")
	}

	rank, file = GetRankFile(F6)
	if rank != 5 || file != 5 {
		t.Errorf("Error! GetRankFile return invalid values for F6!")
	}

	rank, file = GetRankFile(F7)
	if rank != 6 || file != 5 {
		t.Errorf("Error! GetRankFile return invalid values for F7!")
	}

	rank, file = GetRankFile(F8)
	if rank != 7 || file != 5 {
		t.Errorf("Error! GetRankFile return invalid values for F8!")
	}

	rank, file = GetRankFile(G1)
	if rank != 0 || file != 6 {
		t.Errorf("Error! GetRankFile return invalid values for G1!")
	}

	rank, file = GetRankFile(G2)
	if rank != 1 || file != 6 {
		t.Errorf("Error! GetRankFile return invalid values for G2!")
	}

	rank, file = GetRankFile(G3)
	if rank != 2 || file != 6 {
		t.Errorf("Error! GetRankFile return invalid values for G3!")
	}

	rank, file = GetRankFile(G4)
	if rank != 3 || file != 6 {
		t.Errorf("Error! GetRankFile return invalid values for G4!")
	}

	rank, file = GetRankFile(G5)
	if rank != 4 || file != 6 {
		t.Errorf("Error! GetRankFile return invalid values for G5!")
	}

	rank, file = GetRankFile(G6)
	if rank != 5 || file != 6 {
		t.Errorf("Error! GetRankFile return invalid values for G6!")
	}

	rank, file = GetRankFile(G7)
	if rank != 6 || file != 6 {
		t.Errorf("Error! GetRankFile return invalid values for G7!")
	}

	rank, file = GetRankFile(G8)
	if rank != 7 || file != 6 {
		t.Errorf("Error! GetRankFile return invalid values for G8!")
	}

	rank, file = GetRankFile(H1)
	if rank != 0 || file != 7 {
		t.Errorf("Error! GetRankFile return invalid values for H1!")
	}

	rank, file = GetRankFile(H2)
	if rank != 1 || file != 7 {
		t.Errorf("Error! GetRankFile return invalid values for H2!")
	}

	rank, file = GetRankFile(H3)
	if rank != 2 || file != 7 {
		t.Errorf("Error! GetRankFile return invalid values for H3!")
	}

	rank, file = GetRankFile(H4)
	if rank != 3 || file != 7 {
		t.Errorf("Error! GetRankFile return invalid values for H4!")
	}

	rank, file = GetRankFile(H5)
	if rank != 4 || file != 7 {
		t.Errorf("Error! GetRankFile return invalid values for H5!")
	}

	rank, file = GetRankFile(H6)
	if rank != 5 || file != 7 {
		t.Errorf("Error! GetRankFile return invalid values for H6!")
	}

	rank, file = GetRankFile(H7)
	if rank != 6 || file != 7 {
		t.Errorf("Error! GetRankFile return invalid values for H7!")
	}

	rank, file = GetRankFile(H8)
	if rank != 7 || file != 7 {
		t.Errorf("Error! GetRankFile return invalid values for H8!")
	}

}
