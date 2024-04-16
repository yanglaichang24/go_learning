package main

import(
   "context"
   "go-micro.dev/v4"
)
type Request struct {
	Name string `json:"name"`
}

type Response struct {
	Message string `json:"message"`
}

type Helloworld struct {
}

func (h *Helloworld) Greeting(ctx context.Context, req *Request, rsp *Response) error {
	rsp.Message = "Hello " + req.Name
	return nil
}

func main() {
	// create a new service
	service := micro.NewService(
		micro.Name("helloworld"),
		micro.Handle(new(Helloworld)),
		micro.Address(":8089"),
	)

	// initialise flags
	service.Init()

	// start the service
	service.Run()

}
