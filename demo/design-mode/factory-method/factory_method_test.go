package factorymethod

import "testing"

func TestFactoryMethod(t *testing.T) {
	var factory ToyFactory

	factory = &PuzzleFactory{}
	toy1 := factory.New()
	toy1.Play()

	factory = &MarbleFactory{}
	toy2 := factory.New()
	toy2.Play()
}

/*
超级接口
type HyperFactory interface {
	CreateFactory(typ int) AbstractFactory
}
抽象工厂接口
type AbstractFactory interface {
	CreateCellphone() Cellphone
	CreateIpad() Ipad
	CreateSmartSoundBox() SmartSoundBox
}


1. HuaweiIpad HuaweiCellphone 实体生产 -> Cellphone Ipad (实现接口)
2. HuaweiFactory 实体 -> AbstractFactory(接口)
*/
