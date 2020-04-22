package player

import (
   ".."
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

func (u *UserPlayer) GetMove(b application.Board) (x, y int) {
   for {
      x, y = u.getUserEnteredCoords()
      if b.IsOccupied(x, y) {
         io.ReportUserMistake("tried to override a already played tile")
      } else {
         return
      }
   }
}

func (u *UserPlayer) getUserEnteredCoords() (x, y int) {
   switch u.inType {
   case NUMPAD:
      pos :=  io.AskIntBounded(
         "Press the key in your numeric keyboard corresponding to the tile " +
         "you want to play",
         1,
         10,
      )
      x, y = convertToPosition(pos)

   case PAIR:
      x, y = io.AskIntPairBounded(
         "Enter the position you want to play in a tuple, as in \"x,y\"",
         0,
         3,
         0,
         3,
      )
   }

   return
}

func convertToPosition(p int) (int, int) {
   return (p - 1) % 3, 2 - (p - 1) / 3
}
