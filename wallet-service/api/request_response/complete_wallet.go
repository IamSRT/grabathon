package request_response

/**
 * Created by Sai Ravi Teja K on 07, Dec 2019
 * © Refugee Inc
 */

type Balance struct {
	UserId  string `json:"user_id"`
	Balance float64 `json:"balance"`
}

type CompleteWallet struct {
	WalletId int       `json:"wallet_id"`
	UserId   string    `json:"user_id"`
	Amount   float64   `json:"amount"`
	Balances []Balance `json:"balances"`
}