package strtpl

import (
	"testing"
)

func BenchmarkEval(b *testing.B) {
	var tests = []struct {
		input string
		data  interface{}
	}{
		{"replacement from {{ .Map }}", map[string]string{"Map": "any map"}},
		{"replacement from {{ .Struct }} and {{ .Struct }} really", struct{ Struct string }{Struct: "any struct"}},
		{"replacement from string {{ . }}", "this is given here"},
	}

	for n := 0; n < b.N; n++ {
		for _, tt := range tests {
			Eval(tt.input, tt.data)
		}
	}
}
