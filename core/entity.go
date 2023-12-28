package core

import (
	"time"

	"github.com/leekchan/accounting"
)

type Transaction struct {
	ID             int    `json:"id"`
	UnixID         string `json:"unix_id"`
	CampaignID     string `json:"campaign_id"`
	UserInvestorID int    `json:"user_investor_id"`
	Amount         int    `json:"amount"`
	Code           string `json:"code"`
	StatusPayment  string `json:"status_payment"`
	PaymentURL     string `json:"payment_url"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func (t Transaction) AmountFormatIDR() string {
	ac := accounting.Accounting{Symbol: "Rp", Precision: 2, Thousand: ".", Decimal: ","}
	return ac.FormatMoney(t.Amount)
}
