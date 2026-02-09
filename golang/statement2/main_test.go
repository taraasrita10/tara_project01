package main

import (
	"testing"
)

func TestCheck(t *testing.T) {
	testsCheck := []struct {
		name  string
		word  string
		slice []string
		want  bool
	}{
		{"checkTrue", "first", []string{"first", "second"}, true},
		{"checkFalse", "third", []string{"first", "second"}, false},
		{"checkEmpty", "first", []string{}, false},
	}

	for _, tt := range testsCheck {
		t.Run(tt.name, func(t *testing.T) {
			got := check(tt.word, &tt.slice)
			if got != tt.want {
				t.Errorf("%s, got : %v, want : %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestAdd(t *testing.T) {
	testsAdd := []struct {
		name      string
		word      string
		slice     []string
		wantslice []string
	}{
		{"add", "first", []string{"second"}, []string{"second", "first"}},
		{"addEmpty", "", []string{"second"}, []string{"second"}},
	}

	for _, tt := range testsAdd {
		t.Run(tt.name, func(t *testing.T) {
			add(tt.word, &tt.slice)
			got := tt.slice
			for i, _ := range got {
				if got[i] != tt.wantslice[i] {
					t.Errorf("%s, got : %v, want : %v", tt.name, got, tt.wantslice)
				}
			}
		})
	}
}

func TestRemove(t *testing.T) {
	testsRemove := []struct {
		name  string
		word  string
		slice []string
		want  []string
	}{
		{"removeExists", "first", []string{"first", "second"}, []string{"second"}},
		{"removeNonExists", "first", []string{"second", "third"}, []string{"second", "third"}},
		{"removeEmpty", "", []string{"first", "second"}, []string{"first", "second"}},
	}

	for _, tt := range testsRemove {
		t.Run(tt.name, func(t *testing.T) {
			remove(tt.word, &tt.slice)
			got := tt.slice
			for i, _ := range got {
				if got[i] != tt.want[i] {
					t.Errorf("%s, got : %v, want : %v", tt.name, got, tt.want)
				}
			}
		})
	}
}
