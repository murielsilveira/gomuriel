package gomuriel

func SumDiagonals(matrix [][]int, size int) int {
  var sum = 0
  for i := 0; i < size; i = i + 1 {
    for j := 0; j < size; j = j + 1 {
      if i == j {
        sum = sum + matrix[i][j]
      }
      if i == (size - j - 1) {
        sum = sum - matrix[i][j]
      }
    }
  }
  if sum < 0 {
    sum = sum * -1
  }
  return sum
}
