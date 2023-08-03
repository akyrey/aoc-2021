package day11

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestInitialStepIncrease(t *testing.T) {
	got := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	initialStepIncrease(got)

	expected := [][]int{{2, 3, 4}, {5, 6, 7}, {8, 9, 10}}

	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected matrix %v, got %v", expected, got)
	}
}

func TestIncreaseAdjacents(t *testing.T) {
	got := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	increaseAdjacents(got, 0, 0)

	expected := [][]int{{1, 3, 3}, {5, 6, 6}, {7, 8, 9}}

	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("expected matrix %v, got %v", expected, got)
	}
}

func TestDay11(t *testing.T) {
	for i := 1; i <= 4; i++ {
		t.Run(fmt.Sprintf("progress after %d step", i), func(t *testing.T) {
			f, err := os.Open("test11.txt")
			if err != nil {
				t.Fatalf("Unable to open test11.txt: %v", err)
			}

			got := buildInitialMatrix(f)
			flashes := 0

			for step := 0; step < i; step++ {
				_ = performStep(flashes, got)
			}

			fileName := fmt.Sprintf("test11_%d_step.txt", i)
			f, err = os.Open(fileName)
			if err != nil {
				fmt.Printf("Unable to open %s, skipping step", fileName)
			} else {
				expected := buildInitialMatrix(f)
				if !reflect.DeepEqual(got, expected) {
					t.Fatalf("expected matrix %v, got %v", expected, got)
				}
			}
		})
	}
}
