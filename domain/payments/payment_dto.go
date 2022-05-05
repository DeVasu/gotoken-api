package payments

type Payment struct {
	Id int64 `json:"paymentId"`
	Name string `json:"name"`
	Type string `json:"type"`
	Logo string `json:"logo"`
	UpdatedAt string `json:"updatedAt,omitempty"`
	CreatedAt string `json:"createdAt,omitempty"`
}


