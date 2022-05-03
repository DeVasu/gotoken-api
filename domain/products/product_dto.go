package products

type Product struct {
	Id int64 `json:"productId,omitempty"`
	CategoryId int64 `json:"categoryId,omitempty"`
	Name string `json:"name,omitempty"`
	Image string `json:"image,omitempty"`
	Price int64 `json:"price,omitempty"`
	Stock int64 `json:"stock,omitempty"`
	Discount Discount `json:"discount,omitempty"`
	UpdatedAt string `json:"updatedAt,omitempty"`
	CreatedAt string `json:"createdAt,omitempty"`
}

type Discount struct {
	Qty int64 `json:"qty,omitempty"`
	Type int64 `json:"type,omitempty"`
	Result int64 `json:"result,omitempty"`
	ExpiredAt int64 `json:"expiredAt,omitempty"` 
}


