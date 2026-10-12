package main_test

import (
	"testing"

	"github.com/hellkirl/ads/lab_2/task_a"
)

func TestArchiver(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{name: "1", s: "AAAABBBB", want: "AAAABBBB"},
		{name: "2", s: "AAAAABBBBB", want: "5(A)5(B)"},
		{name: "3", s: "ABABABAB", want: "4(AB)"},
		{name: "4", s: "ABABABCABABABCABABABC", want: "3(3(AB)C)"},
		{name: "5", s: "BACBACBACBBBBBBACBACBACBBBBB", want: "2(3(BAC)5(B))"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := main.Archiver(tt.s)
			if got != tt.want {
				t.Errorf("Archiver() = %v, want %v", got, tt.want)
			}
		})
	}
}
