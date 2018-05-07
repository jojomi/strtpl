package strtpl

import (
	"bytes"
	htmlTemplate "html/template"
	textTemplate "text/template"
)

// Eval applies Golang's text templating functions on a string with given data and returns the resulting string.
func Eval(templateString string, data interface{}) (output string, err error) {
	var outputBuffer bytes.Buffer
	t, err := textTemplate.New("tmpl").Parse(templateString)
	if err != nil {
		return
	}
	err = t.Execute(&outputBuffer, data)
	if err != nil {
		return
	}
	output = outputBuffer.String()
	return
}

// MustEval applies Golang's text templating functions on a string with given data and returns the resulting string.
// In case of errors on the way, this function panics.
func MustEval(templateString string, data interface{}) (output string) {
	var err error
	output, err = Eval(templateString, data)
	if err != nil {
		panic(err)
	}
	return output
}

// EvalHTML applies Golang's html templating functions on a string with given data and returns the resulting string.
func EvalHTML(templateString string, data interface{}) (output string, err error) {
	var outputBuffer bytes.Buffer
	t, err := htmlTemplate.New("tmpl-html").Parse(templateString)
	if err != nil {
		return
	}
	err = t.Execute(&outputBuffer, data)
	if err != nil {
		return
	}
	output = outputBuffer.String()
	return
}

// MustEvalHTML applies Golang's html templating functions on a string with given data and returns the resulting string.
// In case of errors on the way, this function panics.
func MustEvalHTML(templateString string, data interface{}) (output string) {
	var err error
	output, err = EvalHTML(templateString, data)
	if err != nil {
		panic(err)
	}
	return output
}
