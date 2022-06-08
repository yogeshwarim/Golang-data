package main

import (
	"context"

	"go-micro.dev/v4"
)

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, name *string, msg *string) error {
	*msg = "Hello " + *name
	return nil
}

type smaller struct{}

func (s *smaller) world(ct context.Context, name1 *string, msg1 *string) error {
	*msg1 = "world" + *name1
	return nil
}

func main() {
	service1 := micro.NewService(
		micro.Name("smaller"),
	)
	// create new service
	service := micro.NewService(
		micro.Name("greeter"),
	)

	service.Init()
	// initialise command line
	service.Init()
	micro.RegisterHandler(service1.Server(), new(smaller))
	// set the handler
	micro.RegisterHandler(service.Server(), new(Greeter))

	// run service
	service1.Run()
	service.Run()
}
