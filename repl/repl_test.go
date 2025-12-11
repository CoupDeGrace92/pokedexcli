package repl

import (
	"testing"
	"strings"

)

func TestCleanInput(t *testing.T) {
	cases := []struct{
		input    string
		expected []string
	}{
		{
			input:     "   hello   world",
			expected: []string{"hello","world"},
		},
		{
			input: "what are you doing booboo?",
			expected: []string{"what", "are","you","doing","booboo?"},
		},
	}

	for _, c := range cases {
		actual := CleanInput(c.input)
		if len(actual) != len(c.expected){
			t.Errorf("expected has %v members and actual has %v members", len(c.expected),len(actual))
			t.Logf("Actual slice: %v", actual)
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord{
				t.Errorf("expected %s from clean input, got %s", expectedWord, word)
			}
		}
	}
}

func TestBufioIO(t *testing.T){
	cases:=[]struct{
		input    string
		expected string
	}{
		{
			input: "cool command bro",
			expected: "cool",
		},
		{
			input: "Mom, I am going to Disneyland!!",
			expected: "mom,",
		},
		{
			input: "    testing whitespaces",
			expected: "testing",
		},
	}

	for _, c:= range cases{
		inputString := c.input
		reader := strings.NewReader(inputString)
		actual := bufReaderOneLoop(reader)

		if actual != c.expected {
			t.Errorf("expected: %s actual: %v", c.expected, actual)
		}
	}
}
