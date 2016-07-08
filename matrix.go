package gomuriel

func SumDiagonals(matrix [][]int, size int) int {
  var sum = 0
  for i := 0; i < size; i = i + 1 {
    sum = sum + matrix[i][i]
    sum = sum - matrix[i][size - i - 1]
  }
  if sum < 0 {
    sum = sum * -1
  }
  return sum
}
