package menu

import "errors"

var (
	ErrCustomerNameIsRequired    error = errors.New("Customer Name is required")
	ErrFulfilledIsRequired       error = errors.New("Fulfilled is required")
	ErrIDIsRequired              error = errors.New("ID is required")
	ErrItemIDIsRequired          error = errors.New("ItemID is required")
	ErrItemNotFound              error = errors.New("Item Not Found")
	ErrItemRepositoryIsRequired  error = errors.New("ItemRepository is required")
	ErrObservationIsRequired     error = errors.New("Observation is required")
	ErrOrderRepositoryIsRequired error = errors.New("OrderRepository is required")
	ErrQuantityIsRequired        error = errors.New("Quantity is required")
)
