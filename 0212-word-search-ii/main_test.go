package main

import "testing"

var tests = []struct {
	board [][]byte
	words []string
	want  []string
}{
	{
		[][]byte{
			{'o', 'a', 'a', 'n'},
			{'e', 't', 'a', 'e'},
			{'i', 'h', 'k', 'r'},
			{'i', 'f', 'l', 'v'},
		},
		[]string{"oath", "pea", "eat", "rain"},
		[]string{"oath", "eat"},
	},
	{
		[][]byte{
			{'a', 'a'},
		},
		[]string{"a"},
		[]string{"a"},
	},
}

func TestFindWords(t *testing.T) {
	for _, test := range tests {
		got := findWords(test.board, test.words)
		if len(got) != len(test.want) {
			t.Errorf("findWords(%v, %v) = %v; want %v\n", test.board, test.words, got, test.want)
			continue
		}
		for i := range test.want {
			if got[i] != test.want[i] {
				t.Errorf("findWords(%v, %v) = %v; want %v\n", test.board, test.words, got, test.want)
				break
			}
		}
	}
}
