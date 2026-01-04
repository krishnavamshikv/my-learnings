package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "krishna"
	want := regexp.MustCompile(`\b` + name + `\b`)
	message, err := Hello("krishna")
	if !want.MatchString(message) || err != nil {
		t.Errorf(`Hello("krishna") = %q , %v , want match for %#q, nil`, message, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	message, err := Hello("")
	if message != "" || err == nil {
		t.Errorf(`Hello("") = %q, %v, want "",error`, message, err)
	}
}
