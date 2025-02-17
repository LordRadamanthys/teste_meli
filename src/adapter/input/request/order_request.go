package request

type OrderRequest struct {
	Itens []ItemRequest `json:"itens"`
}
