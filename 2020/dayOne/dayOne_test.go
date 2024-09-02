package main

import "testing"

func TestTwoSum(t *testing.T) {
	sampleList := []int{
		1721,
		979,
		366,
		299,
		675,
		1456,
	}

	got := TwoSum(sampleList)
	want := 514579
	if want != got {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
func TestThreeSum(t *testing.T) {
	sampleList := []int{
		1721,
		979,
		366,
		299,
		675,
		1456,
	}

	got := ThreeSum(sampleList)
	want := 241861950
	if want != got {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
