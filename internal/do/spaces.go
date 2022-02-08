package do

import (
	"context"

	"github.com/digitalocean/godo"
)

func GetSpace(ctx context.Context, client *godo.Client, spaceID string) (*godo.Volume, error) {
	space, _, err := client.Storage.GetVolume(ctx, spaceID)
	if err != nil {
		return nil, err
	}

	return space, nil
}
