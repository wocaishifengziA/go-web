package factorymethod

import "fmt"

// 玩具工厂接口
type ToyFactory interface {
	New() Toy
}

// 玩具接口，需要实现一个Play方法
type Toy interface {
	Play()
}

// 拼图玩具，实现了玩具接口
type Puzzle struct {}

func (*Puzzle) Play() {
	fmt.Println("playing puzzle")
}

// 弹珠玩具，实现了玩具接口
type Marble struct {}

func (*Marble) Play() {
	fmt.Println("playing marble")
}

// 拼图厂，实现了玩具厂接口
type PuzzleFactory struct {}

func (p *PuzzleFactory) New() Toy {
	return new(Puzzle)
}

// 拼图厂，实现了玩具厂接口
type MarbleFactory struct {}

func (p *MarbleFactory) New() Toy {
	return new(Marble)
}
