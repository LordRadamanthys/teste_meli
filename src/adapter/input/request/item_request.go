package request

type ItemRequest struct {
	ID       string `json:"id"`
	Quantity int    `json:"quantity"`
}
