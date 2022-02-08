package do

import (
	"context"

	"github.com/digitalocean/godo"
)

func GetDomain(ctx context.Context, client *godo.Client, domainID string) (*godo.Domain, error) {
	domain, _, err := client.Domains.Get(ctx, domainID)
	if err != nil {
		return nil, err
	}

	return domain, nil
}
