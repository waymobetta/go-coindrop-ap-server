package design // The convention consists of naming the design
// package "design"
import (
	. "github.com/goadesign/goa/design/apidsl" // Use . imports to enable the DSL
)

var _ = API("coindrop-ap", func() { // API defines the microservice endpoint and
	Title("Coindrop Accounts Payable Server")               // other global properties. There should be one
	Description("An accounts payable service for coindrop") // and exactly one API definition appearing in
	Scheme("http")                                          // the design.
	Host("localhost:8888")
})
