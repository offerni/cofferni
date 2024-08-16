package menu

import "github.com/offerni/cofferni"

type Service struct {
	ItemRepo  cofferni.ItemRepository
	OrderRepo cofferni.OrderRepository
}

type NewServiceOpts struct {
	ItemRepository  cofferni.ItemRepository
	OrderRepository cofferni.OrderRepository
}

func NewService(opts NewServiceOpts) (*Service, error) {
	return &Service{
		ItemRepo:  opts.ItemRepository,
		OrderRepo: opts.OrderRepository,
	}, nil
}

func (opts NewServiceOpts) Validate() error {
	if opts.ItemRepository == nil {
		return ErrItemRepositoryIsRequired
	}

	if opts.OrderRepository == nil {
		return ErrOrderRepositoryIsRequired
	}

	return nil
}
