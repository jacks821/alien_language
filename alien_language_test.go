package alien_language

import "testing"

type testoptions struct {
	input string
	output []string
}

var extendtests = []testoptions {
	{ "(ab)bc", []string{"ab", "bc"} },
	{ "(ab)(bc)c", []string{"ab", "bc", "c"} },
}

func TestExtendTest (t *testing.T) {
	for _, pair := range extendtests {
		v := ExtendTest(pair.input)
		if v != pair.output {
			t.Error(
				"For", pair.input,
				"expected", pair.output,
				"got", v,
			)
		}
	}
}


