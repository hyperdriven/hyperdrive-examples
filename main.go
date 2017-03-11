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

var Example = &ExampleEndpoint{*hyperdrive.NewEndpoint("Example Endpoint", "Example Endpoint, written with hyperdriven/hyperdrive", "/example", "1")}

type PanicEndpoint struct {
	hyperdrive.Endpoint
}

func (e *PanicEndpoint) Get(rw http.ResponseWriter, r *http.Request) {
	panic("ERROR! Everyone panic!")
	return
}

var PanicExample = &PanicEndpoint{*hyperdrive.NewEndpoint("Panic Endpoint", "This example tests panic recovery.", "/panic", "1")}

func main() {
	api := hyperdrive.NewAPI("Example API", "Example API Description")
	api.AddEndpoint(Example)
	api.AddEndpoint(PanicExample)
	api.Start()
}
