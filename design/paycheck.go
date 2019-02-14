package design

import (
	. "github.com/goadesign/goa/design"        // Use . imports to enable the DSL
	. "github.com/goadesign/goa/design/apidsl" // Use . imports to enable the DSL
)

var _ = Resource("paycheck", func() { // Resources group related API endpoints
	BasePath("/v1/paycheck") // together. They map to REST resources for REST

	Action("send", func() {
		Description("Sends ethereum wallet token amount")
		Routing(POST(""))
		Payload(PaycheckPayload)
		Response(OK)
		Response(NotFound)
		Response(BadRequest, StandardErrorMedia)
		Response(Gone, StandardErrorMedia)
		Response(InternalServerError, StandardErrorMedia)
	})
})

// PaycheckPayload is the payload for sending an Ethereum address a token payment
var PaycheckPayload = Type("PaycheckPayload", func() {
	Description("Paycheck payload")
	Attribute("ethereumAddress", String, "ethereum address")
	Attribute("tokenAmount", Integer, "token amount")
	Required("ethereumAddress", "tokenAmount")
})
