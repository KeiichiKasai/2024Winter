package model

type Wallet struct {
	OwnerId int     `form:"owner_id,omitempty" `
	Balance float64 `form:"balance,omitempty" `
}
