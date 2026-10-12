package main_test

import (
	"testing"

	"github.com/hellkirl/ads/lab_2/task_c"
)

func TestCountFractions(t *testing.T) {
	tests := []struct {
		name    string
		mid     float64
		a       []int
		wantCnt int64
		wantP   int
		wantQ   int
	}{
		{"1", 0.5, []int{1, 2, 3}, 2, 1, 2},
		{"2", 0.7, []int{1, 2, 3}, 3, 2, 3},
		{"3", 0.1, []int{1, 2, 3}, 0, 0, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got2, got3 := main.CountFractions(tt.mid, tt.a)
			if got != tt.wantCnt || got2 != tt.wantP || got3 != tt.wantQ {
				t.Errorf("CountFractions() = %v, %v, %v, want %v, %v, %v", got, got2, got3, tt.wantCnt, tt.wantP, tt.wantQ)
			}
		})
	}
}

func TestFindSmallestFraction(t *testing.T) {
	tests := []struct {
		name string
		n    int
		k    int
		a    []int
		wantP int
		wantQ int
	}{
		{"1", 5, 2, []int{1, 7, 11, 13, 17}, 1, 13},
		{"2", 5, 6, []int{1, 7, 11, 13, 17}, 7, 13},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotP, gotQ := main.FindSmallestFraction(tt.n, tt.k, tt.a)
			if gotP != tt.wantP || gotQ != tt.wantQ {
				t.Errorf("FindSmallestFraction() = %v, %v, want %v, %v", gotP, gotQ, tt.wantP, tt.wantQ)
			}
		})
	}
}
