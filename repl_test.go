package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Yo  soNN thisisanother ",
			expected: []string{"yo", "sonn", "thisisanother"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(c.input) != len(c.expected) {
			t.Errorf("expecten length of actual doesn't match input")
		}

		for i := range actual {
			word := actual[i]
			expected := c.expected[i]

			if word != expected {
				t.Errorf("word %d doesn't match", i)
			}
		}
	}
}
