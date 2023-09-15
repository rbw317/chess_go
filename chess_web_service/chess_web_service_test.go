package chess_web_service

import (
	"chess_go/chess_engine"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetInstance(t *testing.T) {
	ws := GetWebServiceInstance(":8080")

	if ws == nil {
		t.Errorf("Error! GetWebServiceInstance function returned nil for ws instance!")
	}

	ws2 := GetWebServiceInstance(":8080")

	if ws != ws2 {
		t.Errorf("Error! GetWebServiceInstance function returned a different pointer for second call!")
	}
}

func TestGetGames(t *testing.T) {
	var result GetGamesResult
	ws := GetWebServiceInstance(":8080") // Have to create the web service
	ws.Games = nil
	ws.Games = append(ws.Games, chess_engine.NewGame(chess_engine.White))
	ws.Games = append(ws.Games, chess_engine.NewGame(chess_engine.Black))
	req := httptest.NewRequest(http.MethodGet, "/games", nil)
	w := httptest.NewRecorder()
	GetGames(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	if http.StatusOK != res.StatusCode {
		t.Fatalf("GetGames return status expected a StatusOK, instead got: %d", res.StatusCode)
	}

	json.Unmarshal(data, &result)
	if !result.Result.Result {
		t.Fatalf("GetGames returned failue!")
	}

	if len(result.Games) != 2 {
		t.Fatalf("Expected GetGames to return 2 games, actually returned %d!", len(result.Games))
	}

	if result.Games[0].ID != 0 {
		t.Fatalf("Expected GetGames first id to be 0, actually returned %d!", result.Games[0].ID)
	}

	if result.Games[1].ID != 1 {
		t.Fatalf("Expected GetGames first id to be 1, actually returned %d!", result.Games[1].ID)
	}

}

func TestCreateGame(t *testing.T) {
	ws := GetWebServiceInstance(":8080") // Have to create the web service
	ws.Games = nil
	req := httptest.NewRequest(http.MethodPost, "/games?color=black", nil)
	w := httptest.NewRecorder()
	CreateGame(w, req)
	gameLen := len(ws.Games)
	res := w.Result()
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	if http.StatusOK != res.StatusCode {
		t.Fatalf("CreateGame return status expected a StatusOK, instead got: %d", res.StatusCode)
	}

	var m map[string]string
	json.Unmarshal(data, &m)
	if m["error"] != "" {
		t.Errorf("Received error %s from Game POST call", m["error"])
	}

	if gameLen != 1 {
		t.Fatalf("After CreateGame expected 1 game in the web service list but have %d", len(ws.Games))
	}

	if ws.Games[0].UserColor != chess_engine.Black {
		t.Fatalf("After CreateGame expected user color to be white but is black.")
	}
}

func TestGetGameInfo(t *testing.T) {
	var result GetGameInfoResult
	ws := GetWebServiceInstance(":8080") // Have to create the web service
	ws.Games = nil
	ws.Games = append(ws.Games, chess_engine.NewGame(chess_engine.White))
	ws.Games = append(ws.Games, chess_engine.NewGame(chess_engine.Black))
	req := httptest.NewRequest(http.MethodGet, "/games/1", nil)
	w := httptest.NewRecorder()
	//GetGameInfo(w, req)
	ws.Router.ServeHTTP(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	if http.StatusOK != res.StatusCode {
		t.Fatalf("GetGames return status expected a StatusOK, instead got: %d", res.StatusCode)
	}

	json.Unmarshal(data, &result)
	if !result.Result.Result {
		t.Fatalf("GetGameInfo returned failue!")
	}

	if result.Game == nil {
		t.Fatalf("GetGameInfo did not return game info")
	}

	if result.Game.ID != 1 {
		t.Fatalf("Expected GetGameInfo id to be 1, actually returned %d!", result.Game.ID)
	}

	if result.Game.Status != "Engine Move" {
		t.Fatalf("Expected GetGameInfo status to be Engine Move, actually returned %s!", result.Game.Status)
	}
}
