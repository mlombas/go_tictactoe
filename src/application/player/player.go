package player

import "../board"

type Player interface {
   GetMove(board.Board) board.Position
}
