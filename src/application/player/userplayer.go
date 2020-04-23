package player

import (
   "../board"
   "../../io"
)

type InputType int
const (
   NUMPAD = 0
   PAIR = 1
)

type UserPlayer struct {
   inType InputType
}

func NewUserPlayer() UserPlayer {
   var player UserPlayer
   inType := io.AskOfOptions(
      "Which input style you prefer, (n)umpad or (p)air",
      'n',
      'p',
   )

   switch inType {
   case 'n': player.inType = NUMPAD 
   case 'p': player.inType = PAIR
   }
   
   return player
}

func (u UserPlayer) GetMove(b board.Board) (pos board.Position) {
   for {
      x, y := u.getUserEnteredCoords().Extract()
      pos = board.NewPosition(x, y) 
      if b.IsOccupied(pos) {
         io.ReportUserMistake("tried to override a already played tile")
      } else {
         return
      }
   }
}

func (u *UserPlayer) getUserEnteredCoords() (pos board.Position) {
   switch u.inType {
   case NUMPAD:
      inputPos :=  io.AskIntBounded(
         "Press the key in your numeric keyboard corresponding to the tile " +
         "you want to play",
         1,
         10,
      )
      pos = convertToPosition(inputPos)

   case PAIR:
      x, y := io.AskIntPairBounded(
         "Enter the position you want to play in a tuple, as in \"x,y\"",
         0,
         3,
         0,
         3,
      )
      pos = board.NewPosition(x, y)
   }

   return
}

func convertToPosition(p int) board.Position {
   return board.NewPosition((p - 1) % 3, 2 - (p - 1) / 3)
}
