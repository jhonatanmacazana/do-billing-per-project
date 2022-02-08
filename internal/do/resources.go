package do

import (
	"context"
	"strconv"
	"strings"

	"github.com/digitalocean/godo"

	"do-billing-per-project/internal/utils"
)

func ListResourcesFromProject(ctx context.Context, client *godo.Client, projectID string) ([]godo.ProjectResource, error) {
	list := []godo.ProjectResource{}

	options := &godo.ListOptions{}

	for {

		projects, resp, err := client.Projects.ListResources(ctx, projectID, options)
		if err != nil {
			return nil, err
		}

		list = append(list, projects...)

		if resp.Links == nil || resp.Links.IsLastPage() {
			break
		}

		currentPage, err := resp.Links.CurrentPage()
		if err != nil {
			return nil, err
		}

		if currentPage == 0 {
			options.Page = currentPage + 2
		} else {
			options.Page = currentPage + 1
		}

	}

	return list, nil
}

func GetResourcesWithInfo(ctx context.Context, client *godo.Client, resources []godo.ProjectResource) ([]ResourceWithInfo, error) {
	list := []ResourceWithInfo{}

	for _, r := range resources {
		if strings.HasPrefix(r.URN, string(Droplet)) {
			dropletIDstr := strings.Split(r.URN, string(Droplet))[1]
			dropletID, err := strconv.Atoi(dropletIDstr)
			if err != nil {
				return nil, err
			}

			droplet, err := GetDroplet(ctx, client, dropletID)
			if err != nil {
				return nil, err
			}

			createdAt, err := utils.ParseISOString(droplet.Created)
			if err != nil {
				return nil, err
			}

			months := utils.MonthsCountSince(createdAt)

			resourceWithInfo := ResourceWithInfo{
				Name:         droplet.Name,
				Type:         "droplet",
				PriceMonthly: droplet.Size.PriceMonthly,
				PriceHourly:  droplet.Size.PriceHourly,
				Months:       months,
				Total:        float64(months) * droplet.Size.PriceMonthly,
				Meta:         &r,
			}

			list = append(list, resourceWithInfo)
		}

		if strings.HasPrefix(r.URN, string(Domain)) {
			domainName := strings.Split(r.URN, string(Domain))[1]

			resourceWithInfo := ResourceWithInfo{
				Name:         domainName,
				Type:         "domain",
				PriceMonthly: 0,
				PriceHourly:  0,
				Total:        0,
				Meta:         &r,
			}

			list = append(list, resourceWithInfo)
		}

		if strings.HasPrefix(r.URN, string(Space)) {
			spaceName := strings.Split(r.URN, string(Space))[1]

			createdAt, err := utils.ParseISOString(r.AssignedAt)
			if err != nil {
				return nil, err
			}

			months := utils.MonthsCountSince(createdAt)

			resourceWithInfo := ResourceWithInfo{
				Name:         spaceName,
				Type:         "space",
				PriceMonthly: 5,
				Months:       months,
				Total:        float64(months) * 5,
				Meta:         &r,
			}

			list = append(list, resourceWithInfo)
		}
	}

	return list, nil
}
