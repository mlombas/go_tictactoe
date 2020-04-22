package player

import ".."

type Player interface {
   GetMove(application.Board) (int, int)
}
