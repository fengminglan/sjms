package chouxiang

import "fmt"

// ServiceFactory 抽象工厂接口
type ServiceFactory interface {
	Push() Push
}

func GetSportsFactory(service string) (ServiceFactory, error) {
	//具体工厂 xhs fb
	if service == "xhs" {
		return &Xhs{}, nil
	}
	if service == "fb" {
		return &Fb{}, nil
	}
	return nil, fmt.Errorf("Wrong type passed")
}
