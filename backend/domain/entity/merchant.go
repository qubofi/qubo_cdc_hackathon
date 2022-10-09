package entity

type Merchant struct {
	ID            uint64 `json:"id"`
	MerchantName  string `json:"merchantName"`
	Email         string `json:"email"`
	WalletAddress string `json:"walletAddress"`
}
