package application

import (
   "../io"
   "./player"
   "./board"
)

type Game struct {
   players map[board.Tile]player.Player
   gameBoard board.Board
}

func NewGame(p1 player.Player, p2 player.Player, b board.Board) Game {
   players := make(map[board.Tile]player.Player)
   players[board.CROSS] = p1
   players[board.NOTCH] = p2
   return Game {
      players: players,
      gameBoard: b,
   }
}

func (g *Game) PlayUntilEnd() board.Board {
   for !g.gameBoard.Ended() {
      io.PrintStringable(&g.gameBoard)
      g.PlayTurn()
   }

   return g.gameBoard
}

func (g *Game) PlayTurn() {
   currPlayer := g.players[g.gameBoard.GetTurn()]
   move := currPlayer.GetMove(g.gameBoard)
   g.gameBoard.PlayTurn(move)
}
