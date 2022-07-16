package abstractfactory

import (
	"testing"
)

func TestAbstractFactory(t *testing.T) {
	// 创建一个超级工厂，用于生产工厂
	var hyperFactory HyperFactory
	var factory AbstractFactory

	hyperFactory = &HypeFactoryImpl{}

	// 创建具体的工厂

	// 创建华为工厂
	factory = hyperFactory.CreateFactory(Huawei)
	factory.CreateCellphone().Call()
	factory.CreateIpad().Play()
	if factory.CreateSmartSoundBox() != nil {
		t.Fail()
	}

	// 创建小米工厂
	factory = hyperFactory.CreateFactory(Xiaomi)
	factory.CreateCellphone().Call()
	factory.CreateIpad().Play()
	factory.CreateSmartSoundBox().Listen()
}
