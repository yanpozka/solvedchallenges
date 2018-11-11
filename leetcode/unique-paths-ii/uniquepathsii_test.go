package uniquepathsii

import "testing"

func TestUniquePathsii(t *testing.T) {
	{
		input := [][]int{
			{1},
		}
		output := 0

		if sol := uniquePathsWithObstacles(input); sol != output {
			t.Fatalf("Expected output: %d but we got: %d", output, sol)
		}
	}
	{
		input := [][]int{
			{0, 0, 0},
			{0, 1, 0},
			{0, 0, 0},
		}
		output := 2

		if sol := uniquePathsWithObstacles(input); sol != output {
			t.Fatalf("Expected output: %d but we got: %d", output, sol)
		}
	}
}
