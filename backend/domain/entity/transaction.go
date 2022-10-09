package entity

type Transaction struct {
	ID              uint64 `json:"id"`
	SubscriptionId  uint64 `json:"subscriptionId"`
	TransactionHash string `json:"transactionHash"`
}
