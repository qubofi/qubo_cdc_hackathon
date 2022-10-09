package entity

type Product struct {
	ID           uint64 `json:"id"`
	MerchantId   uint64 `json:"merchantId"`
	ProductName  string `json:"productName"`
	Cost         uint64 `json:"cost"`
	Currency     string `json:"currency"`
	Interval     uint64 `json:"interval"`
	IntervalUnit string `json:"intervalUnit"`
}
