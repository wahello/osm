package strings

import (
	"testing"

	tassert "github.com/stretchr/testify/assert"
)

func TestWhichNotEqual(t *testing.T) {
	assertion := tassert.New(t)

	listOfAlpha := Which{"a", "a"}
	listOfStuff := Which{"a", "b"}

	assertion.Nil(listOfAlpha.NotEqual("a"))
	assertion.Equal(listOfAlpha.NotEqual("b"), []string{"a", "a"})
	assertion.Equal(listOfStuff.NotEqual("a"), []string{"b"})
}
