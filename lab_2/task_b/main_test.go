package main_test

import (
	"github.com/hellkirl/ads/lab_2/task_b"
	"testing"
)

func TestCanPlace(t *testing.T) {
	tests := []struct {
		name     string
		a        []int
		k        int
		distance int
		want     bool
	}{
		{name: "1", a: []int{2, 5, 6, 12, 22}, k: 3, distance: 10, want: true},
		{name: "2", a: []int{2, 5, 6, 12, 22}, k: 3, distance: 11, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := main.CanPlace(tt.a, tt.k, tt.distance)
			if got != tt.want {
				t.Errorf("CanPlace() = %v, want %v", got, tt.want)
			}
		})
	}
}
