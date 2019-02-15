package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/goadesign/goa"
	log "github.com/sirupsen/logrus"
	"github.com/waymobetta/go-coindrop-ap-server/app"
	"github.com/waymobetta/go-coindrop-ap-server/client"
	ethsvc "github.com/waymobetta/go-coindrop-ap-server/services/ethereum"
)

type Transaction struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

// PaycheckController implements the paycheck resource.
type PaycheckController struct {
	*goa.Controller
}

// NewPaycheckController creates a paycheck controller.
func NewPaycheckController(service *goa.Service) *PaycheckController {
	return &PaycheckController{Controller: service.NewController("PaycheckController")}
}

// Send runs the send action.
func (c *PaycheckController) Send(ctx *app.SendPaycheckContext) error {
	// PaycheckController_Send: start_implement

	// Put your logic here

	paycheck := &client.PaycheckPayload{
		EthereumAddress: ctx.Payload.EthereumAddress,
		TokenAmount:     ctx.Payload.TokenAmount,
	}

	// if token 9 decimals
	// default: 18
	tokenAmountInWei := fmt.Sprintf("%v000000000", paycheck.TokenAmount)

	transaction, err := ethsvc.SendToken(tokenAmountInWei, paycheck.EthereumAddress)
	if err != nil {
		return ctx.BadRequest(&app.StandardError{
			Code:    400,
			Message: "could not send token",
		})
	}

	log.Printf("sent %v token to %s\n", paycheck.TokenAmount, paycheck.EthereumAddress)

	tx := &Transaction{
		Status: "OK",
		ID:     transaction,
	}

	res, err := json.Marshal(tx)
	if err != nil {
		return ctx.BadRequest(&app.StandardError{
			Code:    400,
			Message: "could not marshal transaction",
		})
	}

	return ctx.OK(res)
	// PaycheckController_Send: end_implement
}
