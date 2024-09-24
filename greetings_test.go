package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {

	name := "Gladys"
	want := regexp.MustCompile(`\b` + name + `\b`)

	message, err := Hello(name)

	if !want.MatchString(message) || err != nil {
		t.Fatalf(`Hello("Gladys") = %q, %v, want %q, nil`, message, err, want)
	}

}

func TestHelloEmpty(t *testing.T) {

	message, err := Hello("")

	if message != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, nil`, message, err)
	}

}
