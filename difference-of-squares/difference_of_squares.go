package diffsquares

// SumOfSquares : sum of the squares of the first n natural numbers
func SumOfSquares(n int) int {
	return (n * (n + 1) * (2*n + 1)) / 6
}

// SquareOfSum : square of the sum of the first n natural numbers
func SquareOfSum(n int) int {
	return (n * (n + 1) / 2) * (n * (n + 1) / 2)
}

// Difference : difference between square of sums and sum of squares
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
