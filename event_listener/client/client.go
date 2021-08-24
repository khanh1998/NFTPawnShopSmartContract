package client

type Client struct {
	Pawn PawnClient
}

func NewClient(host string) *Client {
	return &Client{
		Pawn: *NewPawnClient(host),
	}
}
