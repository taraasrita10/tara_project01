package main

import (
	"testing"
)

func TestAddCheck(t *testing.T) {
	myTrie := newTrie()

	insertWords := []string{"first", "second", "third", "fir"}
	for _, w := range insertWords {
		myTrie.Add(w)
	}

	testsAddCheck := []struct {
		name string
		word string
		want bool
	}{
		{"Exists", "first", true},
		{"Prefix", "thi", false},
		{"EmptyString", "", false},
		{"NonExist", "fourth", false},
	}

	for _, tt := range testsAddCheck {
		t.Run(tt.name, func(t *testing.T) {
			got := myTrie.Check(tt.word)
			if got != tt.want {
				t.Errorf("%s, got : %v, want : %v", tt.name, got, tt.want)
			}
		})
	}

}

func TestRemove(t *testing.T) {
	myTrie := newTrie()

	insertWords := []string{"first", "second", "third", "fir"}
	for _, w := range insertWords {
		myTrie.Add(w)
	}

	removeWords := []string{"second", "first"}

	testsRemove := []struct {
		name string
		word string
		want bool
	}{
		{"Remove", "second", false}, // Post removing the check returns false
		{"RemoveWord", "first", false},
		{"CheckPrefix", "fir", true},
	}

	for _, w := range removeWords {
		myTrie.Remove(w)
	}

	for _, tt := range testsRemove {
		t.Run(tt.name, func(t *testing.T) {
			got := myTrie.Check(tt.word)
			if got != tt.want {
				t.Errorf("%s, got : %v, want : %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestList(t *testing.T) {
	myTrie := newTrie()
	insertWords := []string{"fir", "first", "second", "third"}
	for _, w := range insertWords {
		myTrie.Add(w)
	}

	got := myTrie.List()

	for i, w := range got {
		if w != insertWords[i] {
			t.Errorf("List, got : %v, want : %v", got, insertWords[i])
		}
	}

}
