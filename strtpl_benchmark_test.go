package strtpl

import (
	"testing"
)

func BenchmarkEval(b *testing.B) {
	var err error

	tests := []struct {
		input string
		data  interface{}
	}{
		{"replacement from {{ .Map }}", map[string]string{"Map": "any map"}},
		{"replacement from {{ .Struct }} and {{ .Struct }} really", struct{ Struct string }{Struct: "any struct"}},
		{"replacement from string {{ . }}", "this is given here"},
	}

	for n := 0; n < b.N; n++ {
		for _, tt := range tests {
			_, err = Eval(tt.input, tt.data)
			_ = err
		}
	}
}

func BenchmarkEvalHTML(b *testing.B) {
	var err error

	tests := []struct {
		input string
		data  interface{}
	}{
		{"replacement from {{ .Map }}", map[string]string{"Map": "any map"}},
		{"replacement from {{ .Struct }} and {{ .Struct }} really", struct{ Struct string }{Struct: "any struct"}},
		{"replacement from string {{ . }}", "this <b>is</b> given here"},
	}

	for n := 0; n < b.N; n++ {
		for _, tt := range tests {
			_, err = EvalHTML(tt.input, tt.data)
			_ = err
		}
	}
}
