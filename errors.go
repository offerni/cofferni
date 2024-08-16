package cofferni

import "errors"

var (
	ErrItemIDIsRequired   error = errors.New("ItemID is required")
	ErrQuantityIsRequired error = errors.New("Quantity is required")
)
