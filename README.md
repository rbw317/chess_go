# chess_go
Chess engine written in Go

The default for the program is to instantiate a web server on port 80 which servers the pages under the chess_web_page folder and
implements the web service in the chess_web_service folder.  If a "-c" command line argument is passed in then it will 
run the console interface for playing chess games from the command line.  The HTML GUI chess board is from https://chessboardjs.com/.  

The engine is up and running and can be tried out at http://www.chessengine.org.

# Rest Web Service
Provides an API for creating and playing chess games with the Go chess engine.

API Calls
- /api/games GET - Gets the list of games.  Returns a JSON array of GameInfo objects
    i.e. [{'id':'1','status':'UserMove'}]

- /api/games?color={user_color} POST - Creates a new chess game.  Returns a GameInfo object for the new game
    i.e. {'id':'1','status':'UserMove'}
- /api/games/{game_id} GET - Returns the GameInfo object for the specified game id or an Error object.
- /api/games/{game_id}/moves GET - Returns a JSON array of MoveInfo objects which are the moves that have been made 
in the specified game or an Error object.
- /api/games/{game_id}/moves?start={start_position}&end={end_position} POST - Makes a new move in the game.  Passed in start 
and end positions in rowcol format (i.e. a1, h3, etc). If no parameters passed in then it should be an engine move.  
Returns a MoveInfo with an updated GameInfo object or an Error object.
- /api/games/{game_id}/moves/{move_id} GET - Returns a MoveInfo object for the specified move ID or an Error.

Building and Deploying the Docker Container
- Build the container: docker build --tag chess-go:latest .
- Run the container: docker run -p 80:80 chess-go:latest
