package player

import (
   "math/rand"
   "../board"
   "../../io"
)

type _Case struct {
   value float64
   play board.Position
}

type AIPlayer struct {
   phrases []string
}

func NewAIPlayer(phrases ...string) AIPlayer {
   return AIPlayer { phrases: phrases }
}

func (self AIPlayer) GetMove(b board.Board) board.Position {
   possible := getPossibleMoves(b)

   cases := make([]_Case, len(possible))
   for i, v := range possible {
      cases[i] = _Case { 
         value: calcPlayValue(b, v),
         play: v,
      }
   }

   max := cases[0]
   for _, v := range cases {
      if v.value > max.value {
         max = v
      }
   }
   
   self.saySomething()

   return max.play
}

func (self *AIPlayer) saySomething() {
   selected := self.phrases[rand.Intn(len(self.phrases))]
   io.Announce("Evil Ruin-Seeking AI", selected)
}

func calcPlayValue(b board.Board, play board.Position) float64 {
   board := b.Copy()
   turn := board.GetTurn()
   board.PlayTurn(play)

   ch := make(chan float64)
   go calculateWinCount(&board, turn, ch)
   return <-ch
}

func calculateWinCount(b *board.Board, 
whoShouldWin board.Tile, ch chan float64) {
   if b.Ended() {
      winner := b.WhoWon()
      if winner == whoShouldWin {
         ch <- 1.0
      } else if winner == board.NOTHING {
         ch <- 0.0
      } else {
         ch <- -1.0
      }
   } else {
      possible := getPossibleMoves(*b)

      subCh := make(chan float64)
      nRoutines := len(possible)
      for i := 0; i < nRoutines; i++ {
         newBoard := b.Copy()
         newBoard.PlayTurn(possible[i])
         go calculateWinCount(&newBoard, whoShouldWin, subCh)
      }

      sum := 0.0
      for i := 0; i < nRoutines; i++ {
         var returned float64
         returned = <-subCh
         sum += returned
      }

      ch <- (sum / 10)
   }
}

func getPossibleMoves(b board.Board) []board.Position {
   result := make([]board.Position, 0)
   for x := 0; x < b.GetDimension(); x++ {
      for y := 0; y < b.GetDimension(); y++ {
         pos := board.NewPosition(x, y)
         if !b.IsOccupied(pos) {
            result = append(result, pos)
         }
      }
   }

   return result
}

