package propublica_api

type Client struct {
	token string
}

func New(token string) *Client {
	s := &Client{}
	s.token = token

	return s
}