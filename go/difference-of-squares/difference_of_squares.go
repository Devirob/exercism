package diffsquares

func SquareOfSum(n int) int {
	var sumOfNumbers = 0
	for i := 1; i <= n; i++ {
		sumOfNumbers += i
	}
	return sumOfNumbers * sumOfNumbers
}

func SumOfSquares(n int) int {

	var sumOfNumbers = 0
	for i := 1; i <= n; i++ {
		sumOfNumbers += i * i
	}
	return sumOfNumbers

}

func Difference(n int) int {

	return SquareOfSum(n) - SumOfSquares(n)

}
