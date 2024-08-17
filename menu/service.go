package menu

import "github.com/offerni/cofferni"

type Service struct {
	itemRepo  cofferni.ItemRepository
	orderRepo cofferni.OrderRepository
}

type NewServiceOpts struct {
	ItemRepository  cofferni.ItemRepository
	OrderRepository cofferni.OrderRepository
}

func NewService(opts NewServiceOpts) (*Service, error) {
	if err := opts.Validate(); err != nil {
		return nil, err
	}

	return &Service{
		itemRepo:  opts.ItemRepository,
		orderRepo: opts.OrderRepository,
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
