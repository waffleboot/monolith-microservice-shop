package application

import (
	"log"
	domain "monolith-microservice-shop/pkg/orders/domain/orders"

	"github.com/pkg/errors"
)

type MarkOrderAsPaidCommand struct {
	OrderID domain.OrderID
}

func (s OrdersService) MarkOrderAsPaid(cmd MarkOrderAsPaidCommand) error {
	o, err := s.repo.ByID(cmd.OrderID)
	if err != nil {
		return errors.Wrapf(err, "cannot get order %s", cmd.OrderID)
	}

	o.MarkAsPaid()

	if err := s.repo.Save(o); err != nil {
		return errors.Wrap(err, "cannot save order")
	}

	log.Printf("marked order %s as paid", cmd.OrderID)

	return nil
}
