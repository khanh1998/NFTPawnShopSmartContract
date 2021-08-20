package model

type Bid struct {
	ID                 string `json:"id" bson:"id,omitempty"`
	Creator            string `json:"creator" bson:"creator,omitempty"`
	LoanAmount         string `json:"loan_amount" bson:"loan_amount,omitempty"`
	Interest           string `json:"interest" bson:"interest,omitempty"`
	LoanStartTime      string `json:"loan_start_time" bson:"loan_start_time,omitempty"`
	LoanDuration       string `json:"loan_duration" bson:"loan_duration,omitempty"`
	IsInterestProRated string `json:"pro_rated" bson:"pro_rated,omitempty"`
}

const (
	BidCollectionName = "bids"
)
