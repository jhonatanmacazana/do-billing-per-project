package do

import (
	"context"
	"fmt"

	"github.com/digitalocean/godo"
)

func ListProjects(ctx context.Context, client *godo.Client) ([]godo.Project, error) {
	list := []godo.Project{}

	options := &godo.ListOptions{}

	for {

		projects, resp, err := client.Projects.List(ctx, options)
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

func GetProjectByName(projects []godo.Project, name string) (*godo.Project, error) {
	for _, v := range projects {
		if v.Name == name {
			return &v, nil
		}
	}

	return nil, fmt.Errorf("project \"%s\" not found", name)
}
