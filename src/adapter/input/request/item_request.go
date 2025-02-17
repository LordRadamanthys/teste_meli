package request

type ItemRequest struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Quantity int    `json:"quantity"`
}
