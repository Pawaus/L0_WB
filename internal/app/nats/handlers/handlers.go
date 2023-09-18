package handlers

import (
	"encoding/json"
	"fmt"
	"l0_wb/internal/pkg/domain"

	"github.com/go-playground/validator/v10"
	stan "github.com/nats-io/stan.go"
)

type Usecase interface {
	ProcessOrder(uid, data string)
	LoadCache()
}

type Handler struct {
	usecase Usecase
	nats    *stan.Conn
}

func NewHandler(u Usecase, nats *stan.Conn) *Handler {
	handler := &Handler{usecase: u, nats: nats}
	handler.usecase.LoadCache()
	(*handler.nats).Subscribe("input_orders", handler.WorkOrder)
	return handler
}

func (h *Handler) WorkOrder(m *stan.Msg) {
	msg := string(m.Data)
	d := domain.Order{}
	err := json.Unmarshal(m.Data, &d)
	if err != nil {
		fmt.Printf("Error data in stream,failed to parce json: %v", err.Error())
		return
	}
	validate := validator.New(validator.WithRequiredStructEnabled())
	err_valid := validate.Struct(d)
	if err_valid != nil {
		fmt.Printf("Fail in recieve data: fount empty fields or incorrect fields")
		return
	}
	h.usecase.ProcessOrder(d.Order_uid, msg)
}
