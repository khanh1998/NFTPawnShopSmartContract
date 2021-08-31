package client

type Client struct {
	Pawn    *PawnClient
	Bid     *BidClient
	BidPawn *BidPawnClient
}

func NewClient(host string, pawnPath string, bidPath string, bidPawnPath string) *Client {
	return &Client{
		Pawn:    newPawnClient(host, pawnPath),
		Bid:     newBidClient(host, bidPath),
		BidPawn: newBidPawnClient(host, bidPawnPath),
	}
}
