package strtpl

import (
	"bytes"
	"html/template"
)

func Eval(templateString string, data interface{}) (output string, err error) {
	var outputBuffer bytes.Buffer
	t, err := template.New("tmpl").Parse(templateString)
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

func MustEval(templateString string, data interface{}) (output string) {
	var err error
	output, err = Eval(templateString, data)
	if err != nil {
		panic(err)
	}
	return output
}
