package abstractfactory

import "fmt"

const (
	Huawei = iota
	Xiaomi
	Unsupported
)

// 抽象工厂接口,需要能够生产手机和Ipad
type AbstractFactory interface {
	CreateCellphone() Cellphone
	CreateIpad() Ipad
	CreateSmartSoundBox() SmartSoundBox
}

// 超级工厂接口，创建一个工厂
type HyperFactory interface {
	CreateFactory(typ int) AbstractFactory
}

// 超级工厂实例
type HypeFactoryImpl struct {}

// 根据给定参数创建工厂
func (*HypeFactoryImpl) CreateFactory(typ int) AbstractFactory {
	switch  typ{
	case Huawei:
		return &HuaweiFactory{}
	case Xiaomi:
		return &XiaomiFactory{}
	default:
		return nil
	}
}

// 手机接口
type Cellphone interface {
	Call()
}

// Ipad接口
type Ipad interface {
	Play()
}

// 智能音箱接口
type SmartSoundBox interface {
	Listen()
}

// 华为工厂,实现了抽象工厂的两个接口
type HuaweiFactory struct {}

func (*HuaweiFactory) CreateCellphone() Cellphone {
	return &HuaweiCellphone{}
}

func (*HuaweiFactory) CreateIpad() Ipad {
	return &HuaweiIpad{}
}

// 华为工厂不生产智能音箱
func (*HuaweiFactory) CreateSmartSoundBox() SmartSoundBox {
	fmt.Println("Huawei not produce SmartSoundBox")
	return nil
}

// 华为手机，实现了手机接口
type HuaweiCellphone struct {}

func (*HuaweiCellphone) Call() {
	fmt.Println("I made a call on my HuaweiCellphone")
}

// 华为Ipad
type HuaweiIpad struct {}

func (*HuaweiIpad) Play() {
	fmt.Println("I am playing with HuaweiIpad")
}


// 小米工厂,实现了抽象工厂的两个接口
type XiaomiFactory struct {}

func (*XiaomiFactory) CreateCellphone() Cellphone {
	return &XiaomiCellphone{}
}

func (*XiaomiFactory) CreateIpad() Ipad {
	return &XiaomiIpad{}
}

func (*XiaomiFactory) CreateSmartSoundBox() SmartSoundBox {
	return &XiaomiSmartSoundBox{}
}

// 小米手机，实现了手机接口
type XiaomiCellphone struct {}

func (*XiaomiCellphone) Call() {
	fmt.Println("I made a call on my XiaomiCellphone")
}

// 小米Ipad
type XiaomiIpad struct {}

func (*XiaomiIpad) Play() {
	fmt.Println("I am playing with XiaomiIpad")
}

// 小米智能音箱
type XiaomiSmartSoundBox struct {}

func (*XiaomiSmartSoundBox) Listen() {
	fmt.Println("I am listening with XiaomiSmartSoundBox")
}
