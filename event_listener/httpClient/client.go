package httpClient

type Client struct {
	Pawn    *PawnClient
	Bid     *BidClient
	BidPawn *BidPawnClient
	Notify  *NotifyClient
}

func NewClient(
	host string, pawnPath string, bidPath string, bidPawnPath string, notifyHost string, notifyPath string,
) *Client {
	return &Client{
		Pawn:    newPawnClient(host, pawnPath),
		Bid:     newBidClient(host, bidPath),
		BidPawn: newBidPawnClient(host, bidPawnPath),
		Notify:  newNotifyClient(notifyHost, notifyPath),
	}
}
