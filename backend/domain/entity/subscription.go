package entity

type Subscription struct {
	ID                    uint64 `json:"id"`
	ProductId             uint64 `json:"productId"`
	CustomerWalletAddress string `json:"customerWalletAddress"`
	StartedDate           string `json:"startedDate"`
	LastPaidDate          string `json:"lastPaidDate"`
	CancelledAt           string `json:"cancelledAt"`
}
