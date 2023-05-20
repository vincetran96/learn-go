package greetings

import (
	"regexp"
	"testing"
)

// TestHelloName calls greetings.Hello with a valid name
// and checks for a valid return value
// Use regex because Hello() returns a random welcome msg
func TestHelloName(t *testing.T) {
	name := "Peter"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, error := Hello(name)
	if !want.MatchString(msg) || error != nil {
		t.Fatalf(`Hello("Peter") = %q, %v, want match for %#q, nil`, msg, error, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, error := Hello("")
	if msg != "" || error == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, error)
	}
}
