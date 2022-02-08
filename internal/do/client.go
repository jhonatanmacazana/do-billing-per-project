package do

import "github.com/digitalocean/godo"

func NewClient(token string) *godo.Client {
	client := godo.NewFromToken(token)

	return client
}
