package main

import (
	"fmt"
	"net/http"

	"github.com/hyperdriven/hyperdrive"
)

type ExampleEndpoint struct {
	hyperdrive.Endpoint
}

func (e *ExampleEndpoint) Get(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "hello world")
	return
}

var Example = &ExampleEndpoint{hyperdrive.Endpoint{"Example Endpoint", "Example Endpoint, written with hyperdriven/hyperdrive", "/example"}}

type PanicEndpoint struct {
	hyperdrive.Endpoint
}

func (e *PanicEndpoint) Get(rw http.ResponseWriter, r *http.Request) {
	panic("ERROR! Everyone panic!")
	return
}

var PanicExample = &PanicEndpoint{hyperdrive.Endpoint{"Panic Endpoint", "This example tests panic recovery.", "/panic"}}

func main() {
	api := hyperdrive.NewAPI()
	api.AddEndpoint(Example)
	api.AddEndpoint(PanicExample)
	api.Start()
}
