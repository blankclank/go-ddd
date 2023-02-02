package services

import (
	"log"
	"tavern/services/order"

	"github.com/google/uuid"
)

type TavernConfiguration func(t *Tavern) error

type Tavern struct {
	OrderService   *order.OrderService
	BillingService interface{}
}

func NewTavern(cfgs ...TavernConfiguration) (*Tavern, error) {
	t := &Tavern{}
	for _, cfg := range cfgs {
		err := cfg(t)
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}

func WithOrderService(os *order.OrderService) TavernConfiguration {
	return func(t *Tavern) error {
		t.OrderService = os
		return nil
	}
}

func (t *Tavern) Order(customerID uuid.UUID, productIDs []uuid.UUID) error {
	price, err := t.OrderService.CreateOrder(customerID, productIDs)

	if err != nil {
		return nil
	}

	log.Printf("Bill the customer : %0.0f", price)

	return nil
}
