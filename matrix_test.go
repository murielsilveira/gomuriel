package gomuriel

import (
  "testing"
)

func TestWithSizeThree(t *testing.T) {
  var size = 3
  var input = [][]int{
    {1, 2, 3},
    {4, 5, 6},
    {7, 8, 9},
  }

  var sum = SumDiagonals(input, size)

  if sum != 0 {
    t.Errorf("Expected 0, got %d", sum)
  }
}

func TestWithSizeThreeExampleTwo(t *testing.T) {
  var size = 3
  var input = [][]int{
    {11, 2, 4},
    {4, 5, 6},
    {10, 8, -12},
  }

  var sum = SumDiagonals(input, size)

  if sum != 15 {
    t.Errorf("Expected 15, got %d", sum)
  }
}

func TestWithSizeFour(t *testing.T) {
  var size = 4
  var input = [][]int{
    {1, 2, 3, 4},
    {5, 6, 7, 8},
    {9, 8, 7, 6},
    {5, 4, 3, 2},
  }

  var sum = SumDiagonals(input, size)

  if sum != 8 {
    t.Errorf("Expected 8, got %d", sum)
  }
}

func TestWithSizeFive(t *testing.T) {
  var size = 5
  var input = [][]int{
    {1, 2, 3, 4, 5},
    {5, 6, 7, 8, 9},
    {9, 8, 7, 6, 7},
    {5, 4, 3, 2, 3},
    {5, 4, 3, 2, 3},
  }

  var sum = SumDiagonals(input, size)

  if sum != 10 {
    t.Errorf("Expected 10, got %d", sum)
  }
}
