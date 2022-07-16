package simplefactory

import "fmt"

const(
	puzzle = iota
	marble
	unsupported
)

type Toy interface {
	Play()
}

type ToyFactory struct{}

func (t *ToyFactory) New(typ int) Toy {
	switch typ{
	case puzzle:
		return &Puzzle{}
	case marble:
		return &Marble{}
	default:
		fmt.Println("unsupported type of toy")
		return nil
	}
}

type Puzzle struct{
	Name string
}

func (*Puzzle) Play() {
	fmt.Println("palying puzzle!")
}

type Marble struct {
	Name string
}

func (*Marble) Play() {
	fmt.Println("playing marble")
}