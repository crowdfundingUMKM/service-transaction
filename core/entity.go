package core

import (
	"time"

	"github.com/leekchan/accounting"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Transaction represents the MongoDB document structure for the "transaction" collection
type Transaction struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UnixID         string             `bson:"unix_id" json:"unix_id"`
	CampaignID     string             `bson:"campaign_id" json:"campaign_id"`
	UserInvestorID int                `bson:"user_investor_id" json:"user_investor_id"`
	Amount         int                `bson:"amount" json:"amount"`
	Code           string             `bson:"code" json:"code"`
	StatusPayment  string             `bson:"status_payment" json:"status_payment"`
	Fraud          bool               `bson:"fraud" json:"fraud"`
	CreatedAt      time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt      time.Time          `bson:"updated_at" json:"updated_at"`
}

func (t Transaction) AmountFormatIDR() string {
	ac := accounting.Accounting{Symbol: "Rp", Precision: 2, Thousand: ".", Decimal: ","}
	return ac.FormatMoney(t.Amount)
}
