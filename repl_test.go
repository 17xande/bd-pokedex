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
		if len(actual) != len(c.expected) {
			t.Errorf("expecten length of %d actual doesn't match input %d", len(actual), len(c.expected))
		}

		for i := range actual {
			word := actual[i]
			expected := c.expected[i]

			if word != expected {
				t.Errorf("word %d:%s doesn't match %s", i, word, expected)
			}
		}
	}
}
