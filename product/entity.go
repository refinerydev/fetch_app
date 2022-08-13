package product

type ProductEntity struct {
	ID        int    `json:"id"`
	CreatedAt string `json:"createdAt"`
	Price     string `json:"price"`
	Product   string `json:"product"`
}

type ConvertResult struct {
	Result  float32 `json:"result"`
	Success bool    `json:"success"`
}
