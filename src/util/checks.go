package util

func IsInBound(value, min, max int) bool {
   return value >= min &&  value < max
}
