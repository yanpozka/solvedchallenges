package sortcharactersbyfrequency

import "testing"

func TestFrequencySort(t *testing.T) {
	if frequencySort("tree") != "eert" {
		t.Fatal("tree")
	}
	if out := frequencySort("cccaaa"); out != "aaaccc" {
		t.Fatal("cccaaa", out)
	}
	if frequencySort("Aabb") != "bbAa" {
		t.Fatal("Aabb")
	}
}
