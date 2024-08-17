package rest

type ListResponse[T any] struct {
	Data []T `json:"data"`
}
