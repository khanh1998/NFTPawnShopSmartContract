type PawnStatus int

const (
	CREATED PawnStatus = iota
	CANCELLED
	DEAL
	LIQUIDATED
	REPAID
)

type Pawn struct {
	Creator      string     `json:"creator" bson:"creator,omitempty"`
	TokenAddress string     `json:"token_address" bson:"token_address,omitempty"`
	TokenId      int        `json:"token_id" bson:"token_id,omitempty"`
	Status       PawnStatus `json:"status" bson:"status,omitempty"`
}