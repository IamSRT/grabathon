package request_response

import "grabathon/models"

/**
 * Created by Sai Ravi Teja K on 28, Nov 2019
 * © Refugee Inc
 */

type Wallet struct {
	Id      int `json:"id"`
	UserId  string `json:"user_id"`
	Balance float64 `json:"balance"`
}

func GetWalletRequestResponse(wallet models.Wallet) Wallet {
	wlt := Wallet{
		Id:         wallet.Id,
		UserId: wallet.UserId,
		Balance: wallet.Balance,
	}

	return wlt
}

func GetWalletModel(wallet Wallet) models.Wallet {
	wlt := models.Wallet{
		Id:         wallet.Id,
		UserId: wallet.UserId,
		Balance: wallet.Balance,
	}

	return wlt
}