package simplefactory

import (
	"testing"
)

func TestSimpleFactory(t *testing.T) {
	factory := new(ToyFactory)

	toy1 := factory.New(puzzle)
	toy1.Play()

	toy2 := factory.New(marble)
	toy2.Play()

	toy3 := factory.New(unsupported)
	if toy3 != nil {
		t.Fail()
	}
}
