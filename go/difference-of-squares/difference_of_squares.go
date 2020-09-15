package diffsquares

// SquareOfSum get the sum of 1+..+n and return the sum squared
// example:
//   if n=5
//   x = 15 (1+2+3+4+5)
//   return 225 (15*15)
func SquareOfSum(n int) int {
	return ((n * (n + 1)) / 2) * ((n * (n + 1)) / 2)
}

// SumOfSquares returns the sum of each number squared
// example:
//   if n=5
//   x = 55 (1*1+2*2+3*3+4*4+5*5)
//   return 55
func SumOfSquares(n int) int {
	return ((n * (n + 1)) * (2*n + 1)) / 6
}

// Difference returns the difference of SquareOfSum and SumOfSquares
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
