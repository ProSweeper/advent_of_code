package main

import "testing"

var sampleGrid = []string{
	"..##.......",
	"#...#...#..",
	".#....#..#.",
	"..#.#...#.#",
	".#...##..#.",
	"..#.##.....",
	".#.#.#....#",
	".#........#",
	"#.##...#...",
	"#...##....#",
	".#..#...#.#",
}

type TestSlope struct {
	down     int
	right    int
	treesHit int
}

func TestCountTreeHits(t *testing.T) {
	slopesToTest := map[string]TestSlope{
		"first":  {down: 1, right: 1, treesHit: 2},
		"second": {down: 1, right: 3, treesHit: 7},
		"third":  {down: 1, right: 5, treesHit: 3},
		"fourt":  {down: 1, right: 7, treesHit: 4},
		"fifth":  {down: 2, right: 1, treesHit: 2},
	}
	for name, slope := range slopesToTest {
		t.Run(name, func(t *testing.T) {
			got := CountTreeHits(sampleGrid, slope.down, slope.right)
			want := slope.treesHit
			if got != want {
				t.Errorf("given %+v, got %d, want %d", slope, got, want)
			}
		})
	}
}
