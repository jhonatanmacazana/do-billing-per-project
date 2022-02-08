package do

import (
	"context"

	"github.com/digitalocean/godo"
)

func GetDroplet(ctx context.Context, client *godo.Client, dropletID int) (*godo.Droplet, error) {
	droplet, _, err := client.Droplets.Get(ctx, dropletID)
	if err != nil {
		return nil, err
	}

	return droplet, nil
}
