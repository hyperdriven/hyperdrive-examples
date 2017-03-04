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

func main() {
	api := hyperdrive.NewAPI()
	api.AddEndpoint(Example)
	http.ListenAndServe(":8080", api.Router)
}
