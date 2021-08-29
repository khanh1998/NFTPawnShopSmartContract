package client

type Client struct {
	Pawn *PawnClient
	Bid  *BidClient
}

func NewClient(host string, pawnPath string, bidPath string) *Client {
	return &Client{
		Pawn: NewPawnClient(host, pawnPath),
		Bid:  NewBidClient(host, bidPath),
	}
}
