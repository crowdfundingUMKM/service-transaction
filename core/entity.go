package core

import (
	"time"

	"github.com/leekchan/accounting"
)

// Transaction represents the MongoDB document structure for the "transaction" collection
type Transaction struct {
	ID             int       `json:"id"`
	UnixID         string    `json:"unix_id"`
	CampaignID     string    `json:"campaign_id"`
	UserInvestorID int       `json:"user_investor_id"`
	OrderID        string    `json:"order_id"`
	PaymentType    string    `json:"payment_type"`
	Amount         int       `json:"amount"`
	Code           string    `json:"code"`
	StatusPayment  string    `json:"status_payment"`
	ExpiryTime     time.Time `json:"expiry_time"`
	Fraud          string    `json:"fraud"`
	UrlPayment     string    `json:"url_payment"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func (t Transaction) AmountFormatIDR() string {
	ac := accounting.Accounting{Symbol: "Rp", Precision: 2, Thousand: ".", Decimal: ","}
	return ac.FormatMoney(t.Amount)
}
