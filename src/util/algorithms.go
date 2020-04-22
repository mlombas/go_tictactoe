package util

func ContainsByte(arr []byte, element byte) bool {
   for _, v := range arr {
      if v == element {
         return true
      }
   }

   return false
}
