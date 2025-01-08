package service

type HelloWorldService interface {
	SayHello() string
}

var HelloWorld HelloWorldService = &helloWorldService{}

type helloWorldService struct{}

func NewHelloWorldService() HelloWorldService {
	return &helloWorldService{}
}

// SayHello implements HelloWorldService.
func (h *helloWorldService) SayHello() string {
	return "Hello, World!"
}
